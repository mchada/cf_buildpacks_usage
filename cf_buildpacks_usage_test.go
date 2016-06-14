package main_test

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rahul-kj/cf_buildpacks_usage"
)

var _ = Describe("Cloud Foundry Buildpack Usage Command", func() {
	Describe(".Run", func() {
		var fakeCliConnection *fakes.FakeCliConnection
		var callBuildpackUsageCommandPlugin *BuildpackUsage

		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			callBuildpackUsageCommandPlugin = &BuildpackUsage{}
		})

		It("calls the buildpack-usage command with no arguments", func() {
			fakeAppsResponse := []string{"{\"total_pages\":1,\"total_results\":2,\"resources\":[{\"entity\":{\"name\":\"app1\",\"buildpack\":null,\"detected_buildpack\":\"Node.js\"}},{\"entity\":{\"name\":\"app2\",\"buildpack\":\"Java\",\"detected_buildpack\":null}}]}"}
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(fakeAppsResponse, nil)
			output := io_helpers.CaptureOutput(func() {
				callBuildpackUsageCommandPlugin.Run(fakeCliConnection, []string{"buildpack-usage"})
			})

			Expect(output[1]).To(Equal("Following is the table of apps and buildpacks app deployments"))
			Expect(output[3]).To(Equal("-------------------------------"))
			Expect(output[4]).To(Equal("| app1 - Node.js |"))
			Expect(output[5]).To(Equal("| app2 - Java |"))
			Expect(output[6]).To(Equal("-------------------------------"))
		})
	})
})
