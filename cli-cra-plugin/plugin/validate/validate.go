package validate

import (
	"errors"
	"os"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

func ValidateParameters(context plugin.PluginContext) error {
	_, ok := os.LookupEnv("TOOLCHAIN_ID")

	if !ok {
		return errors.New("You must provide the TOOLCHAIN_ID environment variable.")
	}

	if !context.HasAPIEndpoint() {
		return errors.New("You must select an ibmcloud api endpoint.")
	}

	if !context.HasTargetedRegion() {
		return errors.New("You must select an ibmcloud region.")
	}

	return nil
}
