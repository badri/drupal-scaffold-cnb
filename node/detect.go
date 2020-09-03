package node

import (
	"fmt"
	"github.com/cloudfoundry/packit"
	"os"
	"path/filepath"
)

const BuildpackYML = `---
php:
  version: 7.4.*
  webserver: nginx
  webdirectory: web
`

type BuildPlanMetadata struct {
	Build  bool `toml:"build"`
	Launch bool `toml:"launch"`
}

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		file, err := os.Create(filepath.Join(context.WorkingDir, "buildpack.yml"))
		if err != nil {
			return packit.DetectResult{}, fmt.Errorf("failed to create buildpack.yml: %w", err)
		}
		defer file.Close()

		_, err = file.WriteString(BuildpackYML)
		if err != nil {
			file.Close()
			return packit.DetectResult{}, fmt.Errorf("failed to write buildpack.yml: %w", err)
		}
		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "drupal-scaffold"},
				},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "drupal-scaffold",
					},
					{
						Name:    "node",
						Version: "~10",
						Metadata: BuildPlanMetadata{
							Build:  true,
							Launch: true,
						},
					},
				},
			},
		}, nil
	}
}
