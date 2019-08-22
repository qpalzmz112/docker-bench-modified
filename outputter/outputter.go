package outputter

import (
	"github.com/aquasecurity/bench-common/check"
)

type Outputter interface {
	Output(controls *check.Controls, summary check.Summary) error
}

type Config struct {
	Console
	JSON
	JSONFormat bool
}

func BuildOutputter(controls *check.Controls, summary check.Summary, config *Config) Outputter {
	if (summary.Fail > 0 || summary.Warn > 0 || summary.Pass > 0 || summary.Info > 0) && config.JSONFormat {
		return NewJSON(config.JSON.Filename)
	} else {
		return NewConsole(config.Console.NoRemediations, config.Console.IncludeTestOutput)
	}
}
