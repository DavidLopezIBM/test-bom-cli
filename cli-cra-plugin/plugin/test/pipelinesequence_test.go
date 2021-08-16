package test

import (
	//"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestPipelineSequence(t *testing.T) {
	setEnvironment()

	buildNumber := "test:123"
	logicalAppName := "doi-doi"

	t.Run("Upload build record, publishbuildrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishbuildrecord", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "pass", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Upload build record, publishbuildrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Upload build record, publishbuildrecord: error", out)
			}
		}
	})

	t.Run("Upload build record, buildrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "buildrecord-publish", "--branch", "master", "--repositoryurl", "https://github.com/uparulekar/abc.git", "--commitid", "123123123", "--status", "pass", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Upload build record, buildrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Upload build record, buildrecord-publish: error", out)
			}
		}
	})

	t.Run("Upload test record, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/mochatest.json", "--type", "unittest", "--drilldownurl", "http://www.diod", "--env", "dev", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Upload test record, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Upload test record, publishtestrecord: should have succeeded, got ", out, err)
		}
	})

	t.Run("Upload test record, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/mochatest.json", "--type", "unittest", "--drilldownurl", "http://www.diod", "--env", "dev", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Upload test record, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Upload test record, testrecord-publish: should have succeeded, got ", out, err)
		}
	})

	t.Run("evaluategate", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "gopolicy", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("evaluategate: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("evaluategate: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("gate-evaluate: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("gate-evaluate: should have failed, got ", out, err)
		}
	})

	t.Run("Upload deploy record, publishdeployrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishdeployrecord", "--env", "beta", "--status", "pass", "--joburl", "http://joburl.com", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Upload deploy record, publishdeployrecord: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Upload deploy record, publishdeployrecord: error", out)
			}
		}
	})

	t.Run("Upload deploy record, deployrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "deployrecord-publish", "--env", "beta", "--status", "pass", "--joburl", "http://joburl.com", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Upload deploy record, deployrecord-publish: command failed", err.Error())
		} else {
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Upload deploy record, deployrecord-publish: error", out)
			}
		}
	})
}