package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

// BuildpackUsage represents Buildpack Usage CLI interface
type BuildpackUsage struct{}

// Metadata is the data retrived from the response json
type Metadata struct {
	GUID string `json:"guid"`
}

// GetMetadata provides the Cloud Foundry CLI with metadata to provide user about how to use buildpack-usage command
func (c *BuildpackUsage) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "buildpack-usage",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "buildpack-usage",
				HelpText: "Command to view buildpack usage in current CLI target context.",
				UsageDetails: plugin.Usage{
					Usage: "buildpack-usage\n   cf buildpack-usage",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(BuildpackUsage))
}

// Run is what is executed by the Cloud Foundry CLI when the buildpack-usage command is specified
func (c BuildpackUsage) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "buildpack-usage" {
		orgs := c.GetOrgs(cli)
		spaces := c.GetSpaces(cli)
		apps := c.GetAppData(cli)
		c.PrintBuildpacks(orgs, spaces, apps)
	}
}

// PrintBuildpacks prints the buildpack data to console
func (c BuildpackUsage) PrintBuildpacks(orgs map[string]string, spaces map[string]SpaceSearchEntity, apps AppSearchResults) {
	fmt.Println("")

	fmt.Printf("Following is the table of apps and buildpacks app deployments\n\n")

	fmt.Printf("| %10s | %30s | %30s | %10s | %250s |\n", "ORG", "SPACE", "APPLICATION", "STATE", "BUILDPACK")
	fmt.Printf("| %10s | %30s | %30s | %10s | %250s |\n", "-----", "-----", "-----", "-----", "-----")

	for _, val := range apps.Resources {
		bp := val.Entity.Buildpack
		if bp == "" {
			bp = val.Entity.DetectedBuildpack
		}

		space := spaces[val.Entity.SpaceGUID]
		spaceName := space.Name
		orgName := orgs[space.OrgGUID]

		fmt.Printf("| %10s | %30s | %30s | %10s | %250s |\n", orgName, spaceName, val.Entity.Name, val.Entity.State, bp)
	}

}
