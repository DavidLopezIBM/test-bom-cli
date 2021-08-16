package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/trace"
)

type Client struct{}

func (c *Client) HttpRequest(craContext CraContext, method string, url string, body io.Reader) (*http.Request, error) {
	trace.Logger = trace.NewLogger(craContext.ContextTrace)
	var rbuf []byte
	var rbody (io.Reader) = nil

	if body != nil {
		rbuf, err := ioutil.ReadAll(body)

		if err != nil {
			trace.Logger.Printf("Failed to read the body:", err)
		}

		rbody = bytes.NewReader(rbuf)
	}

	req, err := http.NewRequest(method, url, rbody)

	if err != nil {
		trace.Logger.Println("****************** ERROR RESPONSE **********************")
		trace.Logger.Println("Error: " + err.Error())
		trace.Logger.Println("********************************************************")

		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", craContext.IamToken)

	trace.Logger.Println("\n****************** REQUEST **********************")
	trace.Logger.Println("Method: ", method)
	trace.Logger.Println("Url: ", url)
	trace.Logger.Println("Body: ", string(rbuf))
	trace.Logger.Println("Header.Authorization: *****")
	trace.Logger.Println("Header.Content-Type: ", req.Header.Get("Content-Type"))
	trace.Logger.Println("*************************************************\n")

	return req, nil
}

func (c *Client) doRequest(req *http.Request, out interface{}) (int, error) {
	client := &http.Client{
		Timeout: time.Second * 60, // 1 minute timeout
	}

	resp, err := client.Do(req)

	if err != nil {
		trace.Logger.Println("****************** ERROR RESPONSE **********************")
		trace.Logger.Println("Error: " + err.Error())
		trace.Logger.Println("********************************************************")

		return 500, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(out)

	if err != nil {
		trace.Logger.Println("****************** ERROR RESPONSE **********************")
		trace.Logger.Println("Error: " + err.Error())
		trace.Logger.Println("********************************************************")

		return 500, err
	}

	s, _ := json.MarshalIndent(out, "", "\t")
	trace.Logger.Println("****************** RESPONSE **********************")
	trace.Logger.Println("StatusCode: ", resp.StatusCode)
	trace.Logger.Println("Body: ", string(s))
	trace.Logger.Println("*************************************************\n")

	return resp.StatusCode, nil
}

func (c *Client) Do(req *http.Request, out interface{}) (int, error) {
	delay := 2

	for i := 1; i < 5; i++ {
		statusCode, err := c.doRequest(req, out)

		if err != nil {
			return 500, err
		}

		if statusCode != 503 {
			return statusCode, nil
		}

		// TODO: 401/403 refresh the token and re-try once
		fmt.Println("Received", statusCode, "while calling", req.URL, "attempting again...")
		d, _ := time.ParseDuration(strconv.Itoa(delay) + "s")
		time.Sleep(d)
		delay = delay * 2
	}

	return 500, errors.New("Error connecting to the url. Try again.")
}
