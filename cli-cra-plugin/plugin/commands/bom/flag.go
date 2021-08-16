package bom

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
)

// cliFlags - flags that are accepted by the cli
var cliFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "branch, B",
		Usage: "Required, The repository branch on which the build is being performed.",
	},
	cli.StringFlag{
		Name:  "repositoryurl, R",
		Usage: "Required, The url of the git repository.",
	},
	cli.StringFlag{
		Name:  "commitid, C",
		Usage: "Required, The git commit id.",
	},
	cli.StringFlag{
		Name:  "status, S",
		Usage: "Required, The build status. Acceptable values: pass/fail",
	},
	cli.StringFlag{
		Name:  "buildnumber, N",
		Usage: "Required, Any string that identifies the build.",
	},
	cli.StringFlag{
		Name:  "logicalappname, L",
		Usage: "Required, Name of the application",
	},
	cli.StringFlag{
		Name:  "toolchainid, I",
		Usage: "Required, This flag is optional if TOOLCHAIN_ID environment variable is set. The value of this flag overrides the value of environment variable if both are provided.",
	},
	cli.StringFlag{
		Name:  "joburl, J",
		Usage: "Optional, The url to the job's build logs",
	},
}

// helpMenuFlags - flags to display as part of the help menu
var helpMenuFlags = []plugin.Flag{
	{
		Name:        "branch, B",
		Description: "Required, The repository branch on which the build is being performed.",
		HasValue:    true,
	},
	{
		Name:        "repositoryurl, R",
		Description: "Required, The url of the git repository.",
		HasValue:    true,
	},
	{
		Name:        "commitid, C",
		Description: "Required, The git commit id.",
		HasValue:    true,
	},
	{
		Name:        "status, S",
		Description: "Required, The build status. Acceptable values: pass/fail",
		HasValue:    true,
	},
	{
		Name:        "buildnumber, N",
		Description: "Required, Any string that identifies the build.",
		HasValue:    true,
	},
	{
		Name:        "logicalappname, L",
		Description: "Required, Name of the application",
		HasValue:    true,
	},
	{
		Name:        "toolchainid, I",
		Description: "Required, This flag is optional if TOOLCHAIN_ID environment variable is set. The value of this flag overrides the value of environment variable if both are provided.",
		HasValue:    true,
	},
	{
		Name:        "joburl, J",
		Description: "Optional, The url to the job's build logs",
		HasValue:    true,
	},
}
