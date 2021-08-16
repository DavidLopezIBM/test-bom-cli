package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func spawn(args string) (string, error) {
	//cmdOut, err := exec.Command("bash", "-c", args).Output()
	cmdOut, err := exec.Command("bash", "-c", args).CombinedOutput()
	return string(cmdOut), err
}

func TestValidate(t *testing.T) {

	t.Run("doi help", func(t *testing.T) {
		out, err := spawn("ibmcloud doi")
		if err != nil {
			t.Fatal("TestDoiHelp: failed", err)
		}
		if strings.Index(out, "doi - Integrate with DevOps Insights service") == -1 {
			t.Error("TestDoiHelp: Invalid help text response")
		}
	})

	t.Run("doi sub command", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help k")
		if err == nil {
			t.Fatal("TestDoiHelp: failed", out)
		}
		// if err != nil {
		// 	if strings.Index(out, "'k' is not a registered command. See 'ibmcloud doi help'.") == -1 {
		// 		t.Error("TestInvalidSubCommand: Invalid error text response")
		// 	}
		// }
	})

	t.Run("publishbuildrecord", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help publishbuildrecord")
		if err != nil {
			t.Fatal("TestPublishBuildRecord: failed", out)
		}
		if strings.Index(out, "buildrecord-publish, publishbuildrecord, b - Publish build record to DevOps Insights") == -1 {
			t.Error("TestPublishBuildRecord: Invalid help text response")
		}
	})

	t.Run("publishbuildrecord short", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help b")
		if err != nil {
			t.Fatal("TestPublishBuildRecordShort: failed", out)
		}
		if strings.Index(out, "buildrecord-publish, publishbuildrecord, b - Publish build record to DevOps Insights") == -1 {
			t.Error("TestPublishBuildRecordShort: Invalid help text response")
		}
	})

	t.Run("buildrecord-publish", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help buildrecord-publish")
		if err != nil {
			t.Fatal("TestPublishBuildRecord: failed", out)
		}
		if strings.Index(out, "buildrecord-publish, publishbuildrecord, b - Publish build record to DevOps Insights") == -1 {
			t.Error("TestPublishBuildRecord: Invalid help text response")
		}
	})

	t.Run("publishtestrecord", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help publishtestrecord")
		if err != nil {
			t.Fatal("TestPublishTestRecord: failed", out)
		}
		if strings.Index(out, "testrecord-publish, publishtestrecord, p - Publish test record to DevOps Insights") == -1 {
			t.Error("TestPublishTestRecord: Invalid help text response")
		}
	})

	t.Run("publishtestrecord short", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help p")
		if err != nil {
			t.Fatal("TestPublishTestRecordShort: failed", out)
		}
		if strings.Index(out, "testrecord-publish, publishtestrecord, p - Publish test record to DevOps Insights") == -1 {
			t.Error("TestPublishTestRecordShort: Invalid help text response")
		}
	})

	t.Run("testrecord-publish", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help testrecord-publish")
		if err != nil {
			t.Fatal("TestPublishTestRecord: failed", out)
		}
		if strings.Index(out, "testrecord-publish, publishtestrecord, p - Publish test record to DevOps Insights") == -1 {
			t.Error("TestPublishTestRecord: Invalid help text response")
		}
	})

	t.Run("publishdeployrecord", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help publishdeployrecord")
		if err != nil {
			t.Fatal("TestPublishDeployRecord: failed", out)
		}
		if strings.Index(out, "deployrecord-publish, publishdeployrecord, d - Publish deploy record to DevOps Insights") == -1 {
			t.Error("TestPublishDeployRecord: Invalid help text response")
		}
	})

	t.Run("publishdeployrecord short", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help d")
		if err != nil {
			t.Fatal("TestPublishDeployRecordShort: failed", out)
		}
		if strings.Index(out, "deployrecord-publish, publishdeployrecord, d - Publish deploy record to DevOps Insights") == -1 {
			t.Error("TestPublishDeployRecordShort: Invalid help text response")
		}
	})

	t.Run("deployrecord-publish", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help deployrecord-publish")
		if err != nil {
			t.Fatal("TestPublishDeployRecord: failed", out)
		}
		if strings.Index(out, "deployrecord-publish, publishdeployrecord, d - Publish deploy record to DevOps Insights") == -1 {
			t.Error("TestPublishDeployRecord: Invalid help text response")
		}
	})

	t.Run("evaluategate", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help evaluategate")
		if err != nil {
			t.Fatal("TestEvaluateGate: failed", out)
		}
		if strings.Index(out, "gate-evaluate, evaluategate, g - Evaluate DevOps Insights gate policy") == -1 {
			t.Error("TestEvaluateGate: Invalid help text response")
		}
	})

	t.Run("evaluategate short", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help g")
		if err != nil {
			t.Fatal("TestEvaluateGateShort: failed", out)
		}
		if strings.Index(out, "gate-evaluate, evaluategate, g - Evaluate DevOps Insights gate policy") == -1 {
			t.Error("TestEvaluateGateShort: Invalid help text response")
		}
	})

	t.Run("gate-evaluate", func(t *testing.T) {
		out, err := spawn("ibmcloud doi help gate-evaluate")
		if err != nil {
			t.Fatal("TestEvaluateGate: failed", out)
		}
		if strings.Index(out, "gate-evaluate, evaluategate, g - Evaluate DevOps Insights gate policy") == -1 {
			t.Error("TestEvaluateGate: Invalid help text response")
		}
	})

	if os.Getenv("PIPELINE_ENV") == "" {
		os.Unsetenv("TOOLCHAIN_ID")
		t.Run("env fail 1", func(t *testing.T) {
			out, _ := spawn("ibmcloud doi publishtestrecord --filelocation ./test/mochatest.json --type unitest")
			if strings.Index(out, "doi: testrecord-publish failed, toolchain ID value not provided") == -1 {
				t.Error("TestUploadSingleFileFail1: should have failed with TOOLCHAIN_ID missing, got ", out)
			}
		})

		t.Run("env fail 2", func(t *testing.T) {
			out, _ := spawn("ibmcloud doi testrecord-publish --filelocation ./test/mochatest.json --type unitest")
			if strings.Index(out, "doi: testrecord-publish failed, toolchain ID value not provided") == -1 {
				t.Error("TestUploadSingleFileFail1: should have failed with TOOLCHAIN_ID missing, got ", out)
			}
		})

		t.Run("env fail 3", func(t *testing.T) {
			os.Setenv("TOOLCHAIN_ID", "aa05b2f4-d99a-43b2-a6dd-0d3328079858")
			os.Setenv("LOGICAL_APP_NAME", "someappname")
			out, _ := spawn("ibmcloud doi publishtestrecord --filelocation ./test/mochatest.json --type unitest")
			if strings.Index(out, "doi: Environment variable \"LOGICAL_APP_NAME\" is prohibited.") == -1 {
				t.Error("TestUploadSingleFileFail2: should have failed with LOGICAL_APP_NAME prohibited, got ", out)
			}
		})

		t.Run("env fail 4", func(t *testing.T) {
			os.Setenv("TOOLCHAIN_ID", "aa05b2f4-d99a-43b2-a6dd-0d3328079858")
			os.Setenv("LOGICAL_APP_NAME", "someappname")
			out, _ := spawn("ibmcloud doi testrecord-publish --filelocation ./test/mochatest.json --type unitest")
			if strings.Index(out, "doi: Environment variable \"LOGICAL_APP_NAME\" is prohibited.") == -1 {
				t.Error("TestUploadSingleFileFail2: should have failed with LOGICAL_APP_NAME prohibited, got ", out)
			}
		})

		t.Run("env fail 5", func(t *testing.T) {
			os.Setenv("TOOLCHAIN_ID", "aa05b2f4-d99a-43b2-a6dd-0d3328079858")
			os.Setenv("BUILD_PREFIX", "someprefix")
			os.Unsetenv("LOGICAL_APP_NAME")
			out, _ := spawn("ibmcloud doi publishtestrecord --filelocation ./test/mochatest.json --type unitest")
			if strings.Index(out, "doi: Environment variable \"BUILD_PREFIX\" is prohibited.") == -1 {
				t.Error("TestUploadSingleFileFail2: should have failed with BUILD_PREFIX prohibited, got ", out)
			}
		})

		t.Run("env fail 6", func(t *testing.T) {
			os.Setenv("TOOLCHAIN_ID", "aa05b2f4-d99a-43b2-a6dd-0d3328079858")
			os.Setenv("BUILD_PREFIX", "someprefix")
			os.Unsetenv("LOGICAL_APP_NAME")
			out, _ := spawn("ibmcloud doi testrecord-publish --filelocation ./test/mochatest.json --type unitest")
			if strings.Index(out, "doi: Environment variable \"BUILD_PREFIX\" is prohibited.") == -1 {
				t.Error("TestUploadSingleFileFail2: should have failed with BUILD_PREFIX prohibited, got ", out)
			}
		})
	}
}
