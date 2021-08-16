package test

import (
	//"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestSonarQubeResults(t *testing.T) {
	setEnvironment()

	t.Run("Missing sqtoken, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/report-task.txt", "--type", "sonarqube", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Missing sqtoken, publishtestrecord: Failed", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: testrecord-publish failed, sqtoken flag required for type=sonarqube") == -1 {
			t.Error("Missing sqtoken, publishtestrecord: error", out)
		}
	})

	t.Run("Missing sqtoken, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/report-task.txt", "--type", "sonarqube", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Missing sqtoken, testrecord-publish: Failed", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: testrecord-publish failed, sqtoken flag required for type=sonarqube") == -1 {
			t.Error("Missing sqtoken, testrecord-publish: error", out)
		}
	})

	t.Run("Invalid sqtoken, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/report-task.txt", "--type", "sonarqube", "--sqtoken", "dummy", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid sqtoken, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: testrecord-publish failed, error: Failed to query Sonarqube server https") == -1 {
			t.Error("Invalid sqtoken, publishtestrecord: error", out)
		}
	})

	t.Run("Invalid sqtoken, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/report-task.txt", "--type", "sonarqube", "--sqtoken", "dummy", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid sqtoken, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: testrecord-publish failed, error: Failed to query Sonarqube server https") == -1 {
			t.Error("Invalid sqtoken, testrecord-publish: error", out)
		}
	})

	t.Run("Invalid file format, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/lcov.info", "--type", "sonarqube", "--sqtoken", "dummy", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid file format, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: testrecord-publish failed, error: Failed to read the Sonarqube result file ./files/lcov.info invalid format") == -1 {
			t.Error("Invalid file format, publishtestrecord: error", out)
		}
	})

	t.Run("Invalid file format, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/lcov.info", "--type", "sonarqube", "--sqtoken", "dummy", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid file format, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: testrecord-publish failed, error: Failed to read the Sonarqube result file ./files/lcov.info invalid format") == -1 {
			t.Error("Invalid file format, testrecord-publish: error", out)
		}
	})

	t.Run("Invalid file, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/invalid-report-task.txt", "--type", "sonarqube", "--sqtoken", "dummy", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid file, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: testrecord-publish failed, error: Failed to query Sonarqube server") == -1 {
			t.Error("Invalid file, publishtestrecord: error", out)
		}
	})

	t.Run("Invalid file, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/invalid-report-task.txt", "--type", "sonarqube", "--sqtoken", "dummy", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid file, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: testrecord-publish failed, error: Failed to query Sonarqube server") == -1 {
			t.Error("Invalid file, testrecord-publish: error", out)
		}
	})

	/*
		t.Run("Successful upload", func(t *testing.T) {
			cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/report-task.txt", "--type", "sonarqube", "--sqtoken", "dummy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
			if err != nil {
				t.Error("Successful upload: Failed to invoke the command ", err)
			}
			out := string(cout)
			if strings.Index(out, "OK") == -1 {
				t.Error("Successful upload: error", out)
			}
		})
	*/
}
