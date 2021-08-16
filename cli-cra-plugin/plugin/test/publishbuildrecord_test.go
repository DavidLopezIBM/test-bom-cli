package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestPublishBuildRecord(t *testing.T) {
	setEnvironment()

	t.Run("Missing branch, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: branch flag not specified") == -1 {
				t.Error("Missing branch, publishbuildrecord: error", out)
			}
		} else {
			t.Error("Missing branch, publishbuildrecord: Should have failed")
		}
	})

	t.Run("Missing branch, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: branch flag not specified") == -1 {
				t.Error("Missing branch, buildrecord-publish: error", out)
			}
		} else {
			t.Error("Missing branch, buildrecord-publish: Should have failed")
		}
	})

	t.Run("Missing repository, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--branch", "master").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: repositoryurl flag not specified") == -1 {
				t.Error("Missing repository, publishbuildrecord: error", out)
			}
		} else {
			t.Error("Missing repository, publishbuildrecord: Should have failed")
		}
	})

	t.Run("Missing repository, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--branch", "master").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: repositoryurl flag not specified") == -1 {
				t.Error("Missing repository, buildrecord-publish: error", out)
			}
		} else {
			t.Error("Missing repository, buildrecord-publish: Should have failed")
		}
	})

	t.Run("Missing status, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: status flag not specified") == -1 {
				t.Error("Missing status, publishbuildrecord: error", out)
			}
		} else {
			t.Error("Missing status, publishbuildrecord: Should have failed")
		}
	})

	t.Run("Missing status, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: status flag not specified") == -1 {
				t.Error("Missing status, buildrecord-publish: error", out)
			}
		} else {
			t.Error("Missing status, buildrecord-publish: Should have failed")
		}
	})

	t.Run("Invalid status, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "any", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid status, publishbuildrecord: failed ", err.Error())
		}
		out := string(cout)
		if strings.Index(out, "doi: buildrecord-publish failed, Invalid status value provided") == -1 {
			t.Error("Invalid status, publishbuildrecord: error", out)
		}
	})

	t.Run("Invalid status, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "any", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid status, buildrecord-publish: failed ", err.Error())
		}
		out := string(cout)
		if strings.Index(out, "doi: buildrecord-publish failed, Invalid status value provided") == -1 {
			t.Error("Invalid status, buildrecord-publish: error", out)
		}
	})

	t.Run("Successful upload 1, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload 1, publishbuildrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 1, publishbuildrecord: error", out)
			}
		}
	})

	t.Run("Successful upload 1, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "pass", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload 1, buildrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 1, buildrecord-publish: error", out)
			}
		}
	})

	t.Run("Successful upload 2, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "fail", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload 2, publishbuildrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 2, publishbuildrecord: error", out)
			}
		}
	})

	t.Run("Successful upload 2, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "fail", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload 2, buildrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 2, buildrecord-publish: error", out)
			}
		}
	})

	t.Run("Successful upload 3, publishbuildrecord", func(t *testing.T) {
		target, targetted := os.LookupEnv("INSIGHTS_CLUSTER_URL")
		if targetted == true && target != "" {
			t.Skip("Successful upload 3, publishbuildrecord: Skipped")
		}
		buildNumber := "tester/master:9"
		logicalAppName := "oneibm cloud/DevOps?12"
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "fail", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Successful upload 3, publishbuildrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 3, publishbuildrecord: error", out)
			}
			body, err1 := GetBuildRecords(GetToolchainID(), logicalAppName)
			if err1 == nil {
				assertEqual(t, len(body.([]interface{})), 1)
				firstR := body.([]interface{})[0].(map[string]interface{})
				assertEqual(t, firstR["build_id"], buildNumber)
				assertEqual(t, firstR["build_artifact"], logicalAppName)
			}
		}
	})

	t.Run("Successful upload 3, buildrecord-publish", func(t *testing.T) {
		target, targetted := os.LookupEnv("INSIGHTS_CLUSTER_URL")
		if targetted == true && target != "" {
			t.Skip("Successful upload 3, buildrecord-publish: Skipped")
		}
		buildNumber := "tester/master:19"
		logicalAppName := "oneibm cloud/DevOps?112"
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "fail", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Successful upload 3, buildrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload 3, buildrecord-publish: error", out)
			}
			body, err1 := GetBuildRecords(GetToolchainID(), logicalAppName)
			if err1 == nil {
				assertEqual(t, len(body.([]interface{})), 1)
				firstR := body.([]interface{})[0].(map[string]interface{})
				assertEqual(t, firstR["build_id"], buildNumber)
				assertEqual(t, firstR["build_artifact"], logicalAppName)
			}
		}
	})
}