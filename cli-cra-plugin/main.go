package main

import (
	"os"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	craCliPlugin "github.ibm.com/oneibmcloud/cli-cra-plugin/plugin"
)

func main() {
	defer os.Exit(0)
	plugin.Start(new(craCliPlugin.CraCliPlugin))
}
