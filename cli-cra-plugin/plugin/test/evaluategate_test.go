package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestEvaluateGate(t *testing.T) {
	setEnvironment()

	t.Run("Missing policy evaluategate", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: policy flag not specified") == -1 {
				t.Error("Missing policy: should have failed, got ", out, err)
			}
		} else {
			t.Error("Missing policy: Should have failed")
		}
	})

	t.Run("Missing policy gate-evaluate", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate").CombinedOutput()
		if err != nil {
			out := string(cout)
			if strings.Index(out, "doi: policy flag not specified") == -1 {
				t.Error("Missing policy: should have failed, got ", out, err)
			}
		} else {
			t.Error("Missing policy: Should have failed")
		}
	})

	t.Run("Delete Policies", func(t *testing.T) {
		err := DeletePolicy(GetToolchainID(), "gopolicy")
		if err != nil {
			t.Error("Delete mocha Policy: Failed ", err)
		}
		err1 := DeletePolicy(GetToolchainID(), "goistanbulpolicy")
		if err1 != nil {
			t.Error("Delete istanbul Policy: Failed ", err1)
		}
	})

	t.Run("Invalid policy evaluategate", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid policy: Failed ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: Failed to evaluate gate, Policy: gopolicy not found.") == -1 {
			t.Error("Invalid policy: should have failed, got ", out, err)
		}
	})

	t.Run("Invalid policy gate-evaluate", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("Invalid policy: Failed ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: Failed to evaluate gate, Policy: gopolicy not found.") == -1 {
			t.Error("Invalid policy: should have failed, got ", out, err)
		}
	})

	t.Run("Create Mocha Policy", func(t *testing.T) {
		err := CreateMochaPolicy(GetToolchainID(), "gopolicy")
		if err != nil {
			t.Error("Create Mocha Policy: Failed ", err)
		}
	})

	t.Run("Create Istanbul Policy", func(t *testing.T) {
		err := CreateIstanbulPolicy(GetToolchainID(), "goistanbulpolicy")
		if err != nil {
			t.Error("Create Istanbul Policy: Failed ", err)
		}
	})

	t.Run("evaluategate no test uploaded", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("evaluategate no test uploaded: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  false") == -1 {
			t.Error("evaluategate no test uploaded: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate no test uploaded", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("gate-evaluate no test uploaded: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  false") == -1 {
			t.Error("gate-evaluate no test uploaded: should have failed, got ", out, err)
		}
	})

	t.Run("Successful upload json publishtestrecord", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "publishtestrecord", "--filelocation", "./files/mochatest.json", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload json publishtestrecord: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload json publishtestrecord: should have succeeded, got ", out, err)
		}
	})

	t.Run("Successful upload json testrecord-publish", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "testrecord-publish", "--filelocation", "./files/mochatest.json", "--type", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("Successful upload json testrecord-publish: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "OK") == -1 {
			t.Error("Successful upload json testrecord-publish: should have succeeded, got ", out, err)
		}
	})

	t.Run("evaluategate success", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("evaluategate success: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("evaluategate success: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate success", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("gate-evaluate success: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("gate-evaluate success: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate success with additional CA cert", func(t *testing.T) {
		os.Setenv("ADDITIONAL_CA_CERT", "../files/localhost.pem")
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("gate-evaluate success: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("gate-evaluate success: should have failed, got ", out, err)
		}
	})

	os.Setenv("ADDITIONAL_CA_CERT", "")
	t.Run("evaluategate fail", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "goistanbulpolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("evaluategate fail: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  false") == -1 {
			t.Error("evaluategate fail: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate fail", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "goistanbulpolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("gate-evaluate fail: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  false") == -1 {
			t.Error("gate-evaluate fail: should have failed, got ", out, err)
		}
	})

	t.Run("evaluategate fail invalid policy with forcedecision", func(t *testing.T) {
		_, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "invalidpolicy", "--forcedecision", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			if strings.Index(err.Error(), "exit status 1") == -1 {
				t.Error("evaluategate fail force decision: should have failed, got ", err)
			}
		}
	})

	t.Run("gate-evaluate fail invalid policy with forcedecision", func(t *testing.T) {
		_, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "invalidpolicy", "--forcedecision", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			if strings.Index(err.Error(), "exit status 1") == -1 {
				t.Error("gate-evaluate fail force decision: should have failed, got ", err)
			}
		}
	})

	t.Run("evaluategate fail invalid policy without forcedecision", func(t *testing.T) {
		_, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "invalidpolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			if strings.Index(err.Error(), "exit status 1") == 0 {
				t.Error("evaluategate fail without force decision: should not have failed, got ", err)
			}
		}
	})

	t.Run("gate-evaluate fail invalid policy without forcedecision", func(t *testing.T) {
		_, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "invalidpolicy", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			if strings.Index(err.Error(), "exit status 1") == 0 {
				t.Error("gate-evaluate fail without force decision: should not have failed, got ", err)
			}
		}
	})

	t.Run("evaluategate fail force decision", func(t *testing.T) {
		_, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "goistanbulpolicy", "--forcedecision", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			if strings.Index(err.Error(), "exit status 1") == -1 {
				t.Error("evaluategate fail force decision: should have failed, got ", err)
			}
		}
	})

	t.Run("gate-evaluate fail force decision", func(t *testing.T) {
		_, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "goistanbulpolicy", "--forcedecision", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			if strings.Index(err.Error(), "exit status 1") == -1 {
				t.Error("gate-evaluate fail force decision: should have failed, got ", err)
			}
		}
	})

	t.Run("evaluategate success force decision", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "gopolicy", "--forcedecision", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("evaluategate success force decision: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("evaluategate success force decision: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate success force decision", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--forcedecision", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("gate-evaluate success force decision: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("gate-evaluate success force decision: should have failed, got ", out, err)
		}
	})

	t.Run("evaluategate success with invalid ruletype", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "gopolicy", "--ruletype", "invalidstr", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("evaluategate success with invalid ruletype: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: Failed to evaluate gate, Policy does not have any rules after reduction") == -1 {
			t.Error("evaluategate success with invalid ruletype: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate success with invalid ruletype", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--ruletype", "invalidstr", "--buildnumber", "master:9", "--logicalappname", "doi-test").CombinedOutput()
		if err != nil {
			t.Error("gate-evaluate success with invalid ruletype: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "doi: Failed to evaluate gate, Policy does not have any rules after reduction") == -1 {
			t.Error("gate-evaluate success with invalid ruletype: should have failed, got ", out, err)
		}
	})

	t.Run("evaluategate success with valid ruletype", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "evaluategate", "--policy", "gopolicy", "--ruletype", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("evaluategate success with valid ruletype: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("evaluategate success with valid ruletype: should have failed, got ", out, err)
		}
	})

	t.Run("gate-evaluate success with valid ruletype", func(t *testing.T) {
		cout, err := exec.Command("ibmcloud", "doi", "gate-evaluate", "--policy", "gopolicy", "--ruletype", "unittest", "--buildnumber", "master:9", "--logicalappname", "doi-test").Output()
		if err != nil {
			t.Error("gate-evaluate success with valid ruletype: Failed to invoke the command ", err)
		}
		out := string(cout)
		if strings.Index(out, "Gate decision to proceed is:  true") == -1 {
			t.Error("gate-evaluate success with valid ruletype: should have failed, got ", out, err)
		}
	})
}
