package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"testing"
	"time"
)

func getGateService() string {
	if os.Getenv("IBM_CLOUD_DEVOPS_ENV") == "dev" {
		return "https://gateservice.us-south.devopsinsights.dev.cloud.ibm.com"
	}
	return "https://gateservice.us-south.devopsinsights.cloud.ibm.com"
}

func getDlmsService() string {
	if os.Getenv("IBM_CLOUD_DEVOPS_ENV") == "dev" {
		return "https://dlms.us-south.devopsinsights.dev.cloud.ibm.com"
	}
	return "https://dlms.us-south.devopsinsights.cloud.ibm.com"
}

func getFVTDebug() string {
	return os.Getenv("FVT_DEBUG")
}

//GetToolchainID - get toolchain id, GateServiceFVT in org DevOpsLifecycleAnalytics
func GetToolchainID() string {
	if os.Getenv("PIPELINE_ENV") == "" {
		return os.Getenv("TOOLCHAIN_ID")
	}
	if os.Getenv("DIFF_TOOLCHAIN") == "true" {
		return os.Getenv("TOOLCHAIN_ID")
	}
	return os.Getenv("PIPELINE_TOOLCHAIN_ID")
}

func setEnvironment() {
	toolchainID, _ := os.LookupEnv("TEST_TOOLCHAIN_ID")
	idsURL := "https://devops.ng.bluemix.net/pipelines"
	if os.Getenv("IBM_CLOUD_DEVOPS_ENV") == "dev" {
		idsURL = "https://devops.stage1.ng.bluemix.net/pipelines"
	}
	if os.Getenv("PIPELINE_ENV") == "" { // standalone
		// remove the pipeline environment variables
		os.Unsetenv("BUILD_PREFIX")
		os.Unsetenv("PIPELINE_TOOLCHAIN_ID")
		os.Unsetenv("PIPELINE_STAGE_INPUT_REV")
		os.Unsetenv("IDS_URL")
		os.Unsetenv("PIPELINE_ID")
		os.Unsetenv("PIPELINE_STAGE_ID")
		os.Unsetenv("IDS_JOB_ID")

		// set standalone environment variables
		os.Setenv("TOOLCHAIN_ID", toolchainID)
		os.Setenv("BUILD_NUMBER", "master:1") // Even though set, this should be ignored
		fmt.Println("Standalone")
	} else {
		// remove the standalone environment variables
		os.Setenv("BUILD_NUMBER", "3")
		os.Setenv("IDS_URL", idsURL)

		if os.Getenv("DIFF_TOOLCHAIN") == "true" { // pipeline env sending data to diff toolchain
			os.Setenv("TOOLCHAIN_ID", toolchainID)
			os.Setenv("PIPELINE_STAGE_INPUT_REV", "9")
			os.Setenv("PIPELINE_TOOLCHAIN_ID", "")
			fmt.Println("CD Pipeline Diff toolchain")
		} else if os.Getenv("PIPELINE_BUILD_STAGE") == "true" { // pipeline build stage
			os.Setenv("TOOLCHAIN_ID", "")
			os.Setenv("PIPELINE_TOOLCHAIN_ID", toolchainID)
			os.Setenv("PIPELINE_STAGE_INPUT_REV", "9AB32CD")
			fmt.Println("CD Pipeline")
		} else {
			os.Setenv("TOOLCHAIN_ID", "")
			os.Setenv("PIPELINE_TOOLCHAIN_ID", toolchainID)
			os.Setenv("PIPELINE_STAGE_INPUT_REV", "9")
			fmt.Println("CD Pipeline")
		}

		// Set CD pipeline specific environment variables
		os.Setenv("PIPELINE_ID", "pipeline1")
		os.Setenv("PIPELINE_STAGE_ID", "stage1")
		os.Setenv("IDS_JOB_ID", "job1")
		os.Setenv("PIPELINE_STAGE_EXECUTION_ID", "stageExecution1")
		os.Setenv("TASK_ID", "task1")

		//os.Setenv("CF_APP", "DemoDRA")
		//os.Setenv("GIT_BRANCH", "master")
		//os.Setenv("GIT_URL", "https://github.com/uparulekar/abc.git")
		//os.Setenv("GIT_COMMIT", "12343453456567")

	}
	// common environment variables
	os.Setenv("BLUEMIX_TRACE", "false")
}

// Login - Need to set the api key
func Login() error {
	setEnvironment()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("Environment variable API_KEY missing")
		return errors.New("API_KEY missing")
	}
	_, ierr := exec.Command("ibmcloud", "iam", "oauth-tokens").Output()
	if ierr != nil {
		apiURL := "cloud.ibm.com"
		if os.Getenv("IBM_CLOUD_DEVOPS_ENV") == "dev" {
			apiURL = "test.cloud.ibm.com"
		}
		out, err := exec.Command("ibmcloud", "login", "-a", apiURL, "--apikey", apiKey, "--no-region").Output()
		fmt.Println(string(out))
		if err == nil {
			fmt.Println("LOGIN successful")
		}
		return err
	}
	return nil
}

// GetOauthTokens - get the IAM token if the user is logged in
func GetOauthTokens() (string, error) {
	cout, ierr := exec.Command("ibmcloud", "iam", "oauth-tokens").Output()
	if ierr == nil {
		out := string(cout)
		return out[len("IAM token:  ") : len(out)-1], nil
	}
	return "", ierr
}

// Logout - logout function
func Logout() error {
	_, err := exec.Command("ibmcloud", "logout").Output()
	//fmt.Println(err)
	if err == nil {
		fmt.Println("LOGOUT successful")
	}
	return err
}

// DeleteToolchainData - delete the toolchain data
func DeleteToolchainData(toolchainID string) error {
	url := getDlmsService() + "/v3/toolchainids/" + toolchainID

	statusCode, body, reqerr := httpRequest("DELETE", url, nil)
	if reqerr != nil {
		return errors.New("Failed to delete toolchain Error: " + reqerr.Error())
	}
	if statusCode != 200 {
		return errors.New("Failed to delete toolchain " + fmt.Sprintf("statusCode %v", statusCode) + " Error: " + fmt.Sprintf("body: %v", body))
	}
	return nil
}

// CreateMochaPolicy - utility function to create a policy
func CreateMochaPolicy(toolchainID string, policyName string) error {
	url := getGateService() + "/api/v5/toolchainids/" + toolchainID + "/policies"

	// prepare the payload
	rule := map[string]interface{}{
		"name":        "testsuccesspercentage",
		"format":      "mocha",
		"stage":       "unittest",
		"percentPass": 100,
	}

	rules := [1]map[string]interface{}{rule}

	payload := map[string]interface{}{
		"name":  policyName,
		"rules": rules,
	}

	//fmt.Printf("%+v", payload)
	bytesRepresentation, marshalerr := json.Marshal(payload)
	if marshalerr != nil {
		return errors.New("Failed to json.Marshal the payload Error: " + marshalerr.Error())
	}

	statusCode, body, reqerr := httpRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if reqerr != nil {
		return errors.New("Failed to create mocha policy Error: " + reqerr.Error())
	}
	if statusCode != 201 {
		return errors.New("Failed to create mocha policy " + fmt.Sprintf("statusCode %v", statusCode) + " Error: " + fmt.Sprintf("body: %v", body))
	}
	return nil
}

// CreateIstanbulPolicy - utility function to create a policy
func CreateIstanbulPolicy(toolchainID string, policyName string) error {
	url := getGateService() + "/api/v5/toolchainids/" + toolchainID + "/policies"

	// prepare the payload
	rule := map[string]interface{}{
		"name":         "testcoveragepercentage",
		"format":       "istanbul",
		"stage":        "code",
		"codeCoverage": 100,
	}

	rules := [1]map[string]interface{}{rule}

	payload := map[string]interface{}{
		"name":  policyName,
		"rules": rules,
	}

	//fmt.Printf("%+v", payload)
	bytesRepresentation, marshalerr := json.Marshal(payload)
	if marshalerr != nil {
		return errors.New("Failed to json.Marshal the payload Error: " + marshalerr.Error())
	}

	statusCode, body, reqerr := httpRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if reqerr != nil {
		return errors.New("Failed to create istanbul policy Error: " + reqerr.Error())
	}
	if statusCode != 201 {
		return errors.New("Failed to create istanbul policy " + fmt.Sprintf("statusCode %v", statusCode) + " Error: " + fmt.Sprintf("body: %v", body))
	}
	return nil
}

// DeletePolicy - utility function to delete a policy
func DeletePolicy(toolchainID string, policyName string) error {
	url := getGateService() + "/api/v5/toolchainids/" + toolchainID + "/policies/" + policyName

	statusCode, body, reqerr := httpRequest("DELETE", url, nil)
	if reqerr != nil {
		return errors.New("Failed to delete policy Error: " + reqerr.Error())
	}
	if statusCode != 200 {
		return errors.New("Failed to delete policy " + fmt.Sprintf("statusCode %v", statusCode) + " Error: " + fmt.Sprintf("body: %v", body))
	}
	return nil
}

// GetTestResultSummary - gets the test results summary
func GetTestResultSummary(toolchainID string, appName string, buildNumber string) (interface{}, error) {
	url := getDlmsService() + "/v3/toolchainids/" + toolchainID + "/buildartifacts/" + url.PathEscape(appName) + "/builds/" + url.PathEscape(buildNumber) + "/summaries"

	statusCode, body, reqerr := httpRequest("GET", url, nil)
	if reqerr != nil {
		return nil, errors.New("Failed to get Test Result Summary Error: " + reqerr.Error())
	}
	if statusCode != 200 {
		return nil, errors.New("Failed to get Test Result Summary " + fmt.Sprintf("statusCode %v", statusCode) + " Error: " + fmt.Sprintf("body: %v", body))
	}
	return body, nil
}

// GetBuildRecords - gets the build records for given build_artifact
func GetBuildRecords(toolchainID string, appName string) (interface{}, error) {
	url := getDlmsService() + "/v3/toolchainids/" + toolchainID + "/buildartifacts/" + url.PathEscape(appName) + "/builds"
	//url := getDlmsService() + "/v3/toolchainids/" + toolchainID + "/builds"

	statusCode, body, reqerr := httpRequest("GET", url, nil)
	if reqerr != nil {
		return nil, errors.New("Failed to get build records Error: " + reqerr.Error())
	}
	if statusCode != 200 {
		return nil, errors.New("Failed to get build records " + fmt.Sprintf("statusCode %v", statusCode) + " Error: " + fmt.Sprintf("body: %v", body))
	}
	return body, nil
}

// GetDeploymentRecords - gets the build records for given build_artifact
func GetDeploymentRecords(toolchainID string, appName string, buildNumber string) (interface{}, error) {
	url := getDlmsService() + "/v3/toolchainids/" + toolchainID + "/buildartifacts/" + url.PathEscape(appName) + "/builds/" + url.PathEscape(buildNumber) + "/deployments"
	//url := getDlmsService() + "/v3/toolchainids/" + toolchainID + "/builds"

	statusCode, body, reqerr := httpRequest("GET", url, nil)
	if reqerr != nil {
		return nil, errors.New("Failed to get deployment records Error: " + reqerr.Error())
	}
	if statusCode != 200 {
		return nil, errors.New("Failed to get deployment records " + fmt.Sprintf("statusCode %v", statusCode) + " Error: " + fmt.Sprintf("body: %v", body))
	}
	return body, nil
}

//func httpRequest(method string, url string, body io.Reader) (int, map[string]interface{}, error) {
func httpRequest(method string, url string, body io.Reader) (int, interface{}, error) {

	token, terr := GetOauthTokens()
	if terr != nil {
		fmt.Println(0, nil, "Failed to get oauth token: "+terr.Error())
	}
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("authorization", token)
	req.Header.Set("content-type", "application/json")

	if getFVTDebug() == "true" {
		fmt.Println("\n****************** REQUEST **********************")
		fmt.Println("Method: ", method)
		fmt.Println("Url: ", url)
		fmt.Println("Body: ", body)
		fmt.Println("Headers: ", req.Header)
		fmt.Println("*************************************************")
	}

	client := &http.Client{
		Timeout: time.Second * 60, // 1 minute timeout
	}
	resp, err := client.Do(req)
	if err == nil {
		var result interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		if getFVTDebug() == "true" {
			fmt.Println("****************** RESPONSE **********************")
			fmt.Println("StatusCode: ", resp.StatusCode)
			fmt.Println("Body: ", result)
			fmt.Println("*************************************************")
		}
		defer resp.Body.Close()

		//return resp.StatusCode, body, nil
		return resp.StatusCode, result, nil
	}
	if getFVTDebug() == "true" {
		fmt.Println("****************** ERROR RESPONSE **********************")
		fmt.Println("Error: " + err.Error())
		fmt.Println("********************************************************")
	}
	return 0, nil, err
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
