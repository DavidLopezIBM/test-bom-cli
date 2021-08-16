package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestPublishDeployRecord(t *testing.T) {
	setEnvironment()

	t.Run("Missing env field, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: env flag not specified") == -1 {
				t.Error("Missing env field, publishdeployrecord: error", out)
			}
		} else {
			t.Error("Missing env field, publishdeployrecord: Should have failed")
		}
	})

	t.Run("Missing env field, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: env flag not specified") == -1 {
				t.Error("Missing env field, deployrecord-publish: error", out)
			}
		} else {
			t.Error("Missing env field, deployrecord-publish: Should have failed")
		}
	})

	t.Run("Missing status, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--env", "beta").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: status flag not specified") == -1 {
				t.Error("Missing status, publishdeployrecord: error", out)
			}
		} else {
			t.Error("Missing status, publishdeployrecord: Should have failed")
		}
	})

	t.Run("Missing status, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--env", "beta").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: status flag not specified") == -1 {
				t.Error("Missing status, deployrecord-publish: error", out)
			}
		} else {
			t.Error("Missing status, deployrecord-publish: Should have failed")
		}
	})

	t.Run("Invalid status, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--env", "beta", "--status", "some", "--joburl", "http://joburl.com", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid status, publishdeployrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "doi: deployrecord-publish failed, Invalid status value provided") == -1 {
				t.Error("Invalid status, publishdeployrecord: error", out)
			}
		}
	})

	t.Run("Invalid status, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--env", "beta", "--status", "some", "--joburl", "http://joburl.com", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid status, deployrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "doi: deployrecord-publish failed, Invalid status value provided") == -1 {
				t.Error("Invalid status, deployrecord-publish: error", out)
			}
		}
	})

	t.Run("Successfully upload 1, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--env", "beta", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload 1, publishdeployrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload 1, publishdeployrecord: error", out)
			}
		}
	})

	t.Run("Successfully upload 1, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--env", "beta", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload 1, deployrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload 1, deployrecord-publish: error", out)
			}
		}
	})

	t.Run("Successfully upload 2, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--env", "beta", "--status", "fail", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload 2, publishdeployrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload 2, publishdeployrecord: error", out)
			}
		}
	})

	t.Run("Successfully upload 2, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--env", "beta", "--status", "fail", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload 2, deployrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload 2, deployrecord-publish: error", out)
			}
		}
	})

	t.Run("Successfully upload all args, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--env", "beta", "--status", "pass", "--appurl", "http://appurl.com", "--joburl", "http://joburl.com", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload all args, publishdeployrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload all args, publishdeployrecord: error", out)
			}
		}
	})

	t.Run("Successfully upload all args, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--env", "beta", "--status", "pass", "--appurl", "http://appurl.com", "--joburl", "http://joburl.com", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload all args, deployrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload all args, deployrecord-publish: error", out)
			}
		}
	})

	t.Run("Successful upload 3, publishdeployrecord", func(t *testing.T) {
		target, targetted := os.LookupEnv("INSIGHTS_CLUSTER_URL")
		if targetted == true && target != "" {
			t.Skip("Successful upload 3, publishdeployrecord: Skipped")
		}
		buildNumber := "tester/master:9"
		logicalAppName := "oneibm cloud/DevOps!99"
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--env", "beta", "--status", "fail", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Successful upload 3, publishdeployrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 3, publishdeployrecord: error", out)
			}
			body, err1 := GetDeploymentRecords(GetToolchainID(), logicalAppName, buildNumber)
			if err1 == nil {
				assertEqual(t, len(body.([]interface{})), 1)
				firstR := body.([]interface{})[0].(map[string]interface{})
				assertEqual(t, firstR["build_id"], buildNumber)
				assertEqual(t, firstR["build_artifact"], logicalAppName)
			}
		}
	})

	t.Run("Successful upload 3, deployrecord-publish", func(t *testing.T) {
		target, targetted := os.LookupEnv("INSIGHTS_CLUSTER_URL")
		if targetted == true && target != "" {
			t.Skip("Successful upload 3, deployrecord-publish: Skipped")
		}
		buildNumber := "tester/master:19"
		logicalAppName := "oneibm cloud/DevOps!199"
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--env", "beta", "--status", "fail", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Successful upload 3, deployrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 3, deployrecord-publish: error", out)
			}
			body, err1 := GetDeploymentRecords(GetToolchainID(), logicalAppName, buildNumber)
			if err1 == nil {
				assertEqual(t, len(body.([]interface{})), 1)
				firstR := body.([]interface{})[0].(map[string]interface{})
				assertEqual(t, firstR["build_id"], buildNumber)
				assertEqual(t, firstR["build_artifact"], logicalAppName)
			}
		}
	})
}