package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

// BuildpackUsage represents Buildpack Usage CLI interface
type BuildpackUsage struct{}

// AppSearchResults represents top level attributes of JSON response from Cloud Foundry API
type AppSearchResults struct {
	TotalResults int                  `json:"total_results"`
	TotalPages   int                  `json:"total_pages"`
	Resources    []AppSearchResources `json:"resources"`
}

// AppSearchResources represents resources attribute of JSON response from Cloud Foundry API
type AppSearchResources struct {
	Entity AppSearchEntity `json:"entity"`
}

// AppSearchEntity represents entity attribute of resources attribute within JSON response from Cloud Foundry API
type AppSearchEntity struct {
	Name              string `json:"name"`
	Buildpack         string `json:"buildpack"`
	DetectedBuildpack string `json:"detected_buildpack"`
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
		res := c.GetAppData(cli)
		c.PrintBuildpacks(res)
	}
}

// PrintBuildpacks prints the buildpack data to console
func (c BuildpackUsage) PrintBuildpacks(res AppSearchResults) {
	fmt.Println("")

	fmt.Printf("Following is the table of apps and buildpacks app deployments\n\n")

	fmt.Println("-------------------------------")
	for _, val := range res.Resources {
		bp := val.Entity.Buildpack
		if bp == "" {
			bp = val.Entity.DetectedBuildpack
		}
		fmt.Printf("| %s - %s |\n", val.Entity.Name, bp)
	}
	fmt.Println("-------------------------------")

}

// GetAppData requests all of the Application data from Cloud Foundry
func (c BuildpackUsage) GetAppData(cli plugin.CliConnection) AppSearchResults {
	var res AppSearchResults
	res = c.UnmarshallAppSearchResults("/v2/apps?order-direction=asc&results-per-page=100", cli)

	if res.TotalPages > 1 {
		for i := 2; i <= res.TotalPages; i++ {
			apiUrl := fmt.Sprintf("/v2/apps?order-direction=asc&page=%v&results-per-page=100", strconv.Itoa(i))
			tRes := c.UnmarshallAppSearchResults(apiUrl, cli)
			res.Resources = append(res.Resources, tRes.Resources...)
		}
	}

	return res
}

func (c BuildpackUsage) UnmarshallAppSearchResults(apiUrl string, cli plugin.CliConnection) AppSearchResults {
	var tRes AppSearchResults
	cmd := []string{"curl", apiUrl}
	output, _ := cli.CliCommandWithoutTerminalOutput(cmd...)
	json.Unmarshal([]byte(strings.Join(output, "")), &tRes)

	return tRes
}
