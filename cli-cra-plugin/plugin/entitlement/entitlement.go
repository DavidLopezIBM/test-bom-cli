package entitlement

import (
	"errors"
	"fmt"
	"strings"

	"github.ibm.com/oneibmcloud/cli-cra-plugin/plugin/util"
)

const entitlementsPath = "/cdentitlements/cra/toolchainid/"

var otcBrokerUrls = map[string]string{
	"dev":      "https://otcbroker.us-south.devopsinsights.dev.cloud.ibm.com",
	"us-south": "https://otcbroker.us-south.devopsinsights.cloud.ibm.com",
	"eu-gb":    "https://otcbroker.eu-gb.devopsinsights.cloud.ibm.com",
	"eu-de":    "https://otcbroker.eu-de.devopsinsights.cloud.ibm.com",
	"bnpp":     "https://otcbroker.eu-fr2.devopsinsights.cloud.ibm.com",
	"au-syd":   "https://otcbroker.au-syd.devopsinsights.cloud.ibm.com",
	"jp-tok":   "https://otcbroker.jp-tok.devopsinsights.cloud.ibm.com",
	"us-east":  "https://otcbroker.us-east.devopsinsights.cloud.ibm.com",
	"jp-osa":   "https://otcbroker.jp-osa.devopsinsights.cloud.ibm.com",
	"mon":      "https://otcbroker.mon01.devopsinsights.cloud.ibm.com",
	"ca-tor":   "https://otcbroker.ca-tor.devopsinsights.cloud.ibm.com",
}

type EntitlementResp struct {
	ToolchainId        string            `json:"toolchainId"`
	ToolchainCrn       string            `json:"toolchainCrn"`
	IsEntitled         bool              `json:"isEntitled"`
	IsRgBased          bool              `json:"isRGBased"`
	RgCdInstanceExists bool              `json:"rgCDInstanceExists"`
	ServiceUrls        map[string]string `json:"serviceUrls"`
	Message            string            `json:"message"`
	Details            string            `json:"details"`
	ErrorMessage       string            `json:"msg"`
}

func getEntitlementsUrl(toolchainId string, apiEndpoint string, region string) string {
	// TODO: support dev variable or check for test in toolchain?
	if strings.Contains(apiEndpoint, "test") {
		return otcBrokerUrls["dev"] + entitlementsPath + toolchainId
	}

	return otcBrokerUrls[region] + entitlementsPath + toolchainId
}

// Logic:  https://github.ibm.com/oneibmcloud/devops-insights-ui/blob/1edc5a757512621007c94161ce97faef89e48bb2/src/pages/_app.js#L127-L130
func isEntitled(entitlementsResp EntitlementResp) bool {
	if entitlementsResp.IsEntitled && !entitlementsResp.RgCdInstanceExists {
		return false
	}
	if !entitlementsResp.IsEntitled && !entitlementsResp.IsRgBased {
		return false
	}

	return true
}

func CheckEntitlement(craContext *util.CraContext, apiEndpoint string, region string) error {
	c := new(util.Client)
	entitlementsUrl := getEntitlementsUrl(craContext.ToolchainId, apiEndpoint, region)
	req, err := c.HttpRequest(*craContext, "GET", entitlementsUrl, nil)

	if err != nil {
		return err
	}

	var entitlementsResp EntitlementResp
	statusCode, err := c.Do(req, &entitlementsResp)

	if err != nil {
		return err
	}

	if statusCode == 401 || statusCode == 403 {
		return errors.New("An error occurred when authenticating your request with IBM Cloud. Log in to IBM Cloud, and try your request again.")
	}

	if statusCode != 200 {
		return fmt.Errorf("An error occurred: %s", entitlementsResp.ErrorMessage)
	}

	if !isEntitled(entitlementsResp) {
		fmt.Println(entitlementsResp.Message)
		return errors.New(entitlementsResp.Message)
	}

	// Store response in context
	craContext.ToolchainCrn = entitlementsResp.ToolchainCrn
	craContext.ToolchainId = entitlementsResp.ToolchainId
	craContext.ServiceUrls = entitlementsResp.ServiceUrls

	return nil
}
