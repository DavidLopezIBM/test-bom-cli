package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestPublishTestRecord(t *testing.T) {
	setEnvironment()

	t.Run("Missing File location, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: filelocation flag not specified") == -1 {
				t.Error("Missing File location, publishtestrecord: should have failed, got ", out, err)
			}
		} else {
			t.Error("Missing File location, publishtestrecord: Should have failed ")
		}
	})

	t.Run("Missing File location, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: filelocation flag not specified") == -1 {
				t.Error("Missing File location, testrecord-publish: should have failed, got ", out, err)
			}
		} else {
			t.Error("Missing File location, testrecord-publish: Should have failed ")
		}
	})

	t.Run("Missing Type, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./test/mochatest.json").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: type flag not specified") == -1 {
				t.Error("Missing Type, publishtestrecord: should have failed, got ", out, err)
			}
		} else {
			t.Error("Missing Type, publishtestrecord: Should have failed")
		}
	})

	t.Run("Missing Type, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./test/mochatest.json").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: type flag not specified") == -1 {
				t.Error("Missing Type, testrecord-publish: should have failed, got ", out, err)
			}
		} else {
			t.Error("Missing Type, testrecord-publish: Should have failed")
		}
	})

	t.Run("Invalid Type, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/mochatest.json", "--type", "utest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid Type, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "lifecycle_stage utest not found for toolchain") == -1 {
			t.Error("Invalid Type, publishtestrecord: should have failed, got ", out, err)
		}
	})

	t.Run("Invalid Type, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/mochatest.json", "--type", "utest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid Type, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "lifecycle_stage utest not found for toolchain") == -1 {
			t.Error("Invalid Type, testrecord-publish: should have failed, got ", out, err)
		}
	})

	t.Run("Invalid filelocation Flag, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocatio", "./files/mochatest.json", "--type", "unittest").Output()
		if err != nil {
			if strings.Index(err.Error(), "exit status 1") == -1 {
				t.Error("Invalid filelocation Flag, publishtestrecord: should have failed with exit code 1 ", err.Error())
			}
		}
		out := string(cout)
		if strings.Index(out, "Incorrect Usage: flag provided but not defined: -filelocatio") == -1 {
			t.Error("Invalid filelocation Flag, publishtestrecord: should have failed, got ", out, err)
		}
	})

	t.Run("Invalid filelocation Flag, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocatio", "./files/mochatest.json", "--type", "unittest").Output()
		if err != nil {
			if strings.Index(err.Error(), "exit status 1") == -1 {
				t.Error("Invalid filelocation Flag, testrecord-publish: should have failed with exit code 1 ", err.Error())
			}
		}
		out := string(cout)
		if strings.Index(out, "Incorrect Usage: flag provided but not defined: -filelocatio") == -1 {
			t.Error("Invalid filelocation Flag, testrecord-publish: should have failed, got ", out, err)
		}
	})

	t.Run("Successful upload json, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/mochatest.json", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload json, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload json, publishtestrecord: should have succeeded, got ", out, err)
		}
	})

	t.Run("Successful upload json, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/mochatest.json", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload json, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload json, testrecord-publish: should have succeeded, got ", out, err)
		}
	})

	t.Run("Successful upload all params, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/mochatest.json", "--type", "unittest", "--drilldownurl", "http://www.diod", "--env", "dev", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload all params, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload all params, publishtestrecord: should have succeeded, got ", out, err)
		}
	})

	t.Run("Successful upload all params, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/mochatest.json", "--type", "unittest", "--drilldownurl", "http://www.diod", "--env", "dev", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload all params, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload all params, testrecord-publish: should have succeeded, got ", out, err)
		}
	})

	t.Run("Successful upload lcov, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/lcov.info", "--type", "code", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload lcov, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload lcov, publishtestrecord: should have succeeded, got ", out, err)
		}
	})

	t.Run("Successful upload lcov, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/lcov.info", "--type", "code", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload lcov, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload lcov, testrecord-publish: should have succeeded, got ", out, err)
		}
	})

	// Successfully update buildnumber with slash
	t.Run("Successful upload lcov 1, publishtestrecord", func(t *testing.T) {
		target, targetted := os.LookupEnv("INSIGHTS_CLUSTER_URL")
		if targetted == true && target != "" {
			t.Skip("Successful upload lcov 1, publishtestrecord: Skipped")
		}
		buildNumber := "master/master:9"
		logicalAppName := "oneibm cloud/DevOps#12"
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/lcov.info", "--type", "code", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Successful upload lcov 1, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload lcov 1, publishtestrecord: should have succeeded, got ", out, err)
		}
		body, err1 := GetTestResultSummary(GetToolchainID(), logicalAppName, buildNumber)
		if err1 == nil {
			summaries := body.(map[string]interface{})["summaries"]
			assertEqual(t, len(summaries.([]interface{})), 1)
			firstR := summaries.([]interface{})[0].(map[string]interface{})
			assertEqual(t, firstR["build_id"], buildNumber)
			assertEqual(t, firstR["build_artifact"], logicalAppName)
		}
	})

	t.Run("Successful upload lcov 1, testrecord-publish", func(t *testing.T) {
		target, targetted := os.LookupEnv("INSIGHTS_CLUSTER_URL")
		if targetted == true && target != "" {
			t.Skip("Successful upload lcov 1, testrecord-publish: Skipped")
		}
		buildNumber := "master/master:19"
		logicalAppName := "oneibm cloud/DevOps#112"
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/lcov.info", "--type", "code", "--buildnumber", buildNumber, "--logicalappname", logicalAppName).Output()
		if err != nil {
			t.Error("Successful upload lcov 1, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload lcov 1, testrecord-publish: should have succeeded, got ", out, err)
		}
		body, err1 := GetTestResultSummary(GetToolchainID(), logicalAppName, buildNumber)
		if err1 == nil {
			summaries := body.(map[string]interface{})["summaries"]
			assertEqual(t, len(summaries.([]interface{})), 1)
			firstR := summaries.([]interface{})[0].(map[string]interface{})
			assertEqual(t, firstR["build_id"], buildNumber)
			assertEqual(t, firstR["build_artifact"], logicalAppName)
		}
	})

	// using wild cards
	t.Run("Successful upload with wild card, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/*.json", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Successful upload with wild card, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload with wild card, publishtestrecord: should have succeeded, got ", out, err)
		}
	})

	t.Run("Successful upload with wild card, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/*.json", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Successful upload with wild card, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload with wild card, testrecord-publish: should have succeeded, got ", out, err)
		}
	})

	// specify wildcard in a directory
	t.Run("Successful upload wildcard dir, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/*", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Successful upload 5: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Failed to upload files/lcov.info to DLMS") == -1 {
			t.Error("Successful upload wildcard dir, publishtestrecord: lcov.info should have failed ", out, err)
		}
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload wildcard dir, publishtestrecord: should have succeeded, got ", out, err)
		}
	})

	t.Run("Successful upload wildcard dir, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/*", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Successful upload 5: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Failed to upload files/lcov.info to DLMS") == -1 {
			t.Error("Successful upload wildcard dir, testrecord-publish: lcov.info should have failed ", out, err)
		}
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload wildcard dir, testrecord-publish: should have succeeded, got ", out, err)
		}
	})

	// specify directory - should fail
	t.Run("Upload Dir fail, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Upload Dir fail, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Failed to read the file ./files") == -1 {
			t.Error("Upload Dir fail, publishtestrecord: directory should have failed ", out, err)
		}
	})

	t.Run("Upload Dir fail, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Upload Dir fail, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Failed to read the file ./files") == -1 {
			t.Error("Upload Dir fail, testrecord-publish: directory should have failed ", out, err)
		}
	})

	t.Run("Successful upload xml, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/jpetstore_ut.xml", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload xml, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload xml, publishtestrecord: directory should have failed ", out, err)
		}
	})

	t.Run("Successful upload xml, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/jpetstore_ut.xml", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload xml, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload xml, testrecord-publish: directory should have failed ", out, err)
		}
	})

	t.Run("Invalid Json file, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/invalidjsonfile.json", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid Json file, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: The specified file: ./files/invalidjsonfile.json is not a valid JSON file") == -1 {
			t.Error("Invalid Json file, publishtestrecord: directory should have failed ", out, err)
		}
	})

	t.Run("Invalid Json file, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/invalidjsonfile.json", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid Json file, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: The specified file: ./files/invalidjsonfile.json is not a valid JSON file") == -1 {
			t.Error("Invalid Json file, testrecord-publish: directory should have failed ", out, err)
		}
	})

	t.Run("Invalid XML file, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/invalidjunit.xml", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid XML file, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: Failed to upload ./files/invalidjunit.xml to DLMS") == -1 {
			t.Error("Invalid XML file, publishtestrecord: directory should have failed ", out, err)
		}
	})

	t.Run("Invalid XML file, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/invalidjunit.xml", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid XML file, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: Failed to upload ./files/invalidjunit.xml to DLMS") == -1 {
			t.Error("Invalid XML file, testrecord-publish: directory should have failed ", out, err)
		}
	})

	// upload a file with Byte order mark
	t.Run("Successful upload BOM XML file, publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/devops-insights-dlms-dynamic-2.xml", "--type", "dynamicsecurityscan", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload BOM XML file, publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload BOM XML file, publishtestrecord: should succeed", out, err)
		}
	})

	t.Run("Successful upload BOM XML file, testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/devops-insights-dlms-dynamic-2.xml", "--type", "dynamicsecurityscan", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload BOM XML file, testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload BOM XML file, testrecord-publish: should succeed", out, err)
		}
	})
}