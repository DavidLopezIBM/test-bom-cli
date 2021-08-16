package deployment

import (
	"errors"
	"os"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
)

const (
	name        = "deployment-scan"
	description = "Scan the deployment"
	// PluginName - Name of the plugin
	PluginName = "cra"

	// // PluginName - Description of the plugin
	// PluginDescription = "Integrate with DevOps Insights service"
)

var aliases = []string{"d"}

// Command - Defines the struct that implements the Deployment function
type Command struct{}

// Run - process the bom-generate request
func (cx *Command) Run(c *cli.Context) error {
	return errors.New("error!")
}

// GetCliCommand - get the command with options for cli util
func (cx *Command) GetCliCommand() cli.Command {
	return cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   description,
		Action: func(c *cli.Context) error {
			ui := terminal.NewStdUI()
			validationErr := validateCliArguments()
			if validationErr != nil {
				ui.Failed("%s: %s failed, %s", PluginName, name, terminal.CommandColor(validationErr.Error()))
				os.Exit(1)
			}

			err := cx.Run(c)
			if err != nil {
				cli.ShowSubcommandHelp(c)
				os.Exit(1)
			}
			return nil
		},
		Flags: cliFlags,
	}
}

// GetOptions - get the options for this options
func (cx *Command) GetOptions() plugin.Command {
	return plugin.Command{
		Namespace:   PluginName,
		Name:        name,
		Aliases:     aliases,
		Description: description,
		Usage:       "ibmcloud cra deployment-scan --branch BRANCH --repositoryurl REPOSITORYURL --commitid COMMITID --toolchainid TOOLCHAINID [--joburl JOBURL]",
		Flags:       helpMenuFlags,
	}
}
