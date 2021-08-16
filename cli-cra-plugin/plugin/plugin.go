package plugin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"github.ibm.com/oneibmcloud/cli-cra-plugin/plugin/commands/deployment"
	"github.ibm.com/oneibmcloud/cli-cra-plugin/plugin/commands/example"
	"github.ibm.com/oneibmcloud/cli-cra-plugin/plugin/entitlement"
	"github.ibm.com/oneibmcloud/cli-cra-plugin/plugin/util"
	"github.ibm.com/oneibmcloud/cli-cra-plugin/plugin/validate"
	"github.ibm.com/oneibmcloud/cra-bom-cli/bom"
	"github.ibm.com/oneibmcloud/cra-bom-cli/common"
	"github.ibm.com/oneibmcloud/terraform-ca-go/terraform"
)

var bomCommand = new(bom.Command)
var deploymentCommand = new(deployment.Command)
var exampleCommand = new(example.Command)
var terraformCommand = new(terraform.TfCommand)

// var vulnerabilityCommand = new(vulnerability.Command)

// CraCliPlugin - Defines the plugin struct that implements the Dev plugin
type CraCliPlugin struct{}

// Run - Mandatory plugin method, used to implement the plugin logic
func (dp *CraCliPlugin) Run(context plugin.PluginContext, args []string) {
	craContext := util.CraContext{}
	craContext.ContextTrace = context.Trace()

	ui := terminal.NewStdUI()

	// user logged in?
	if context.IsLoggedIn() == false {
		ui.Failed(`%s: You are not logged in. Log in by running "%s".`, PluginName, terminal.CommandColor("ibmcloud login"))
		os.Exit(1)
	}

	err := validate.ValidateParameters(context)
	if err != nil {
		ui.Failed(`%s: %s`, PluginName, err)
		os.Exit(1)
	}

	// Get parameters
	craContext.IamToken = context.IAMToken()
	craContext.ToolchainId = os.Getenv("TOOLCHAIN_ID")
	apiEndpoint := context.APIEndpoint()
	region := context.CurrentRegion().Name

	err = entitlement.CheckEntitlement(&craContext, apiEndpoint, region)
	if err != nil {
		ui.Failed(`%s: %s`, PluginName, err)
		os.Exit(1)
	}

	fmt.Println("Context passed to commands:")
	s, _ := json.MarshalIndent(craContext, "", "\t")
	fmt.Print(string(s))
	fmt.Println("")

	bomCraContext := common.CraContext{}
	bomCraContext.ContextTrace = craContext.ContextTrace
	bomCraContext.IamToken = craContext.IamToken
	bomCraContext.ServiceUrls = craContext.ServiceUrls
	bomCraContext.ToolchainCrn = craContext.ToolchainCrn
	bomCraContext.ToolchainId = craContext.ToolchainId

	app := cli.NewApp()
	app.Name = PluginName
	app.Usage = PluginDescription
	app.Commands = []cli.Command{
		// New noun-verb cli commands
		// bomCommand.GetCliCommand(craContext),
		bomCommand.GetCliCommand(bomCraContext),
		deploymentCommand.GetCliCommand(),
		exampleCommand.GetCliCommand(),
		terraformCommand.TerraformCodeRiskAnalysis(),
		// vulnerabilityCommand.GetCliCommand(craContext),
		// vulnerabilityCommand.GetCliCommand(),
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)

}

// GetMetadata - Mandatory plugin method, used to return the plugin metadata
func (dp *CraCliPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: PluginName,
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 0,
		},
		Namespaces: []plugin.Namespace{
			plugin.Namespace{
				Name:        PluginName,
				Description: PluginDescription,
			},
		},
		Commands: []plugin.Command{
			// New noun-verb cli commands
			bomCommand.GetOptions(),
			deploymentCommand.GetOptions(),
			exampleCommand.GetOptions(),
			terraformCommand.GetOptions(),
			// vulnerabilityCommand.GetOptions(),
		},
	}
}
