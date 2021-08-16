package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestValidateToolchain(t *testing.T) {
	setEnvironment()

	t.Run("upload build record with env, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload build record: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 1: error", out)
			}
		}
	})

	t.Run("upload build record with env, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload build record: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 1: error", out)
			}
		}
	})

	t.Run("upload test record with env, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/jpetstore_ut.xml", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload xml: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload xml: directory should have failed ", out, err)
		}
	})

	t.Run("upload test record with env, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/jpetstore_ut.xml", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload xml: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload xml: directory should have failed ", out, err)
		}
	})

	t.Run("upload deploy record with env, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--env", "beta", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload 1: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload 1: error", out)
			}
		}
	})

	t.Run("upload deploy record with env, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--env", "beta", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload 1: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload 1: error", out)
			}
		}
	})

	t.Run("evaluategate with env", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("evaluategate success: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("evaluategate success: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate with env", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("gate-evaluate success: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("gate-evaluate success: should have failed, got ", out, err)
		}
	})

	t.Run("updatepolicies dryrun upload policies with env", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "updatepolicies", "--file", "./files/criteria.json", "--dryrun").CombinedOutput()
		if err == nil {
			out := string(cout)
			if !(strings.Contains(out, "doi: List of Policy modifications:") && strings.Contains(out, "doi: List of Custom Dataset modifications:") && !strings.Contains(out, "doi: Sending data to service:")) {
				t.Error("updatepolicies dryrun should only list the changes without sending data to the service, got ", out, err)
			}
		} else {
			t.Error("updatepolicies dryrun should only list the changes without sending data to the service: command failed")
		}
	})

	t.Run("policies-update dryrun upload policies with env", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "policies-update", "--file", "./files/criteria.json", "--dryrun").CombinedOutput()
		if err == nil {
			out := string(cout)
			if !(strings.Contains(out, "doi: List of Policy modifications:") && strings.Contains(out, "doi: List of Custom Dataset modifications:") && !strings.Contains(out, "doi: Sending data to service:")) {
				t.Error("policies-update dryrun should only list the changes without sending data to the service, got ", out, err)
			}
		} else {
			t.Error("policies-update dryrun should only list the changes without sending data to the service: command failed")
		}
	})

	os.Unsetenv("TOOLCHAIN_ID")
	toolchainID := "aa05b2f4-d99a-43b2-a6dd-0d3328079858"
	if os.Getenv("IBM_CLOUD_DEVOPS_ENV") == "dev" {
		toolchainID = "4dc3fb9b-bbf8-4da4-8d6b-67aefe8826da" // GateServiceFVT toolchain, neal.patel org
	}

	t.Run("upload build record with flag, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--toolchainid", toolchainID, "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload build record: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 1: error", out)
			}
		}
	})

	t.Run("upload build record with flag, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--toolchainid", toolchainID, "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload build record: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 1: error", out)
			}
		}
	})

	t.Run("upload test record with flag, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--toolchainid", toolchainID, "--filelocation", "./files/jpetstore_ut.xml", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload xml: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload xml: directory should have failed ", out, err)
		}
	})

	t.Run("upload test record with flag, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--toolchainid", toolchainID, "--filelocation", "./files/jpetstore_ut.xml", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload xml: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload xml: directory should have failed ", out, err)
		}
	})

	t.Run("upload deploy record with flag, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--toolchainid", toolchainID, "--env", "beta", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload 1: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload 1: error", out)
			}
		}
	})

	t.Run("upload deploy record with flag, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--toolchainid", toolchainID, "--env", "beta", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successfully upload 1: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successfully upload 1: error", out)
			}
		}
	})

	t.Run("evaluategate with flag", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--toolchainid", toolchainID, "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("evaluategate success: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("evaluategate success: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate with flag", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--toolchainid", toolchainID, "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("gate-evaluate success: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("gate-evaluate success: should have failed, got ", out, err)
		}
	})

	t.Run("updatepolicies dryrun upload policies with env", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "updatepolicies", "--toolchainid", toolchainID, "--file", "./files/criteria.json", "--dryrun").CombinedOutput()
		if err == nil {
			out := string(cout)
			if !(strings.Contains(out, "doi: List of Policy modifications:") && strings.Contains(out, "doi: List of Custom Dataset modifications:") && !strings.Contains(out, "doi: Sending data to service:")) {
				t.Error("updatepolicies dryrun should only list the changes without sending data to the service, got ", out, err)
			}
		} else {
			t.Error("updatepolicies dryrun should only list the changes without sending data to the service: command failed")
		}
	})

	t.Run("policies-update dryrun upload policies with env", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "policies-update", "--toolchainid", toolchainID, "--file", "./files/criteria.json", "--dryrun").CombinedOutput()
		if err == nil {
			out := string(cout)
			if !(strings.Contains(out, "doi: List of Policy modifications:") && strings.Contains(out, "doi: List of Custom Dataset modifications:") && !strings.Contains(out, "doi: Sending data to service:")) {
				t.Error("policies-update dryrun should only list the changes without sending data to the service, got ", out, err)
			}
		} else {
			t.Error("policies-update dryrun should only list the changes without sending data to the service: command failed")
		}
	})
}