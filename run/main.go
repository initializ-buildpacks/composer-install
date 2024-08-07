package main

import (
	"os"

	"github.com/initializ-buildpacks/composer"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/chronos"
	"github.com/paketo-buildpacks/packit/v2/fs"
	"github.com/paketo-buildpacks/packit/v2/pexec"
	"github.com/paketo-buildpacks/packit/v2/sbom"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

type Generator struct{}

func (f Generator) Generate(dir string) (sbom.SBOM, error) {
	return sbom.Generate(dir)
}

func main() {
	logEmitter := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv(composer.BpLogLevel))
	phpVersionResolver := composer.NewPhpVersionResolver()
	options := composer.NewComposerInstallOptions()

	configExec := pexec.NewExecutable("composer")
	installExec := pexec.NewExecutable("composer")
	globalExec := pexec.NewExecutable("composer")
	checkPlatformReqsExec := pexec.NewExecutable("composer")

	packit.Run(
		composer.Detect(logEmitter, phpVersionResolver),
		composer.Build(
			logEmitter,
			options,
			configExec,
			installExec,
			globalExec,
			checkPlatformReqsExec,
			Generator{},
			os.Getenv("PATH"),
			fs.NewChecksumCalculator(),
			chronos.DefaultClock),
	)
}
