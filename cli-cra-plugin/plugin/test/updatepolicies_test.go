package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestUpdatePolicies(t *testing.T) {
	Login() // setup

	t.Run("alias should work with help flag", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "up", "--help").Output()
		if err == nil {
			out := string(cout)
			if !strings.Contains(out, "policies-update, updatepolicies, up - Update custom data sets and policies for a toolchain") {
				t.Error("alias should work with help flag: should have printed the helper text, got ", out, err)
			}
		} else {
			t.Error("alias should work with help flag: command failed")
		}
	})

	t.Run("updatepolicies dryrun should only list the changes without sending data to the service", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "updatepolicies", "--file", "./files/criteria.json", "--dryrun").CombinedOutput()
		if err == nil {
			out := string(cout)
			//fmt.Println(out)
			if !(strings.Contains(out, "doi: List of Policy modifications:") && strings.Contains(out, "doi: List of Custom Dataset modifications:") && !strings.Contains(out, "doi: Sending data to service:")) {
				t.Error("dryrun should only list the changes without sending data to the service, got ", out, err)
			}
		} else {
			//fmt.Println(err.Error())
			t.Error("dryrun should only list the changes without sending data to the service: command failed")
		}
	})

	t.Run("policies-update dryrun should only list the changes without sending data to the service", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "policies-update", "--file", "./files/criteria.json", "--dryrun").CombinedOutput()
		if err == nil {
			out := string(cout)
			//fmt.Println(out)
			if !(strings.Contains(out, "doi: List of Policy modifications:") && strings.Contains(out, "doi: List of Custom Dataset modifications:") && !strings.Contains(out, "doi: Sending data to service:")) {
				t.Error("dryrun should only list the changes without sending data to the service, got ", out, err)
			}
		} else {
			//fmt.Println(err.Error())
			t.Error("dryrun should only list the changes without sending data to the service: command failed")
		}
	})

	t.Run("updatepolicies should list the changes and send it to the service", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "updatepolicies", "--file", "./files/criteria.json").Output()
		if err == nil {
			out := string(cout)
			if !(strings.Contains(out, "doi: List of Policy modifications:") && strings.Contains(out, "doi: List of Custom Dataset modifications:") && strings.Contains(out, "doi: Sending data to service:")) {
				t.Error("should list the changes and send it to the service", out, err)
			}
		} else {
			t.Error("should list the changes and send it to the service: command failed")
		}
	})

	t.Run("policies-update should list the changes and send it to the service", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "policies-update", "--file", "./files/criteria.json").Output()
		if err == nil {
			out := string(cout)
			if !(strings.Contains(out, "doi: List of Policy modifications:") && strings.Contains(out, "doi: List of Custom Dataset modifications:") && strings.Contains(out, "doi: Sending data to service:")) {
				t.Error("should list the changes and send it to the service", out, err)
			}
		} else {
			t.Error("should list the changes and send it to the service: command failed")
		}
	})
}