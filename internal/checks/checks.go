package checks

import (
	"fmt"
	"github.com/m1sol/api-tester/internal/config"
	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
)

type Status string

const (
	StatusPassed  Status = "passed"
	StatusFailed  Status = "failed"
	StatusSkipped Status = "skipped"
)

type Result struct {
	Type     string
	Status   Status
	Path     string
	Expected any
	Actual   any
	Message  string
}

type Check interface {
	Run(doc *jsonutil.Document, response executor.Result) Result
}

func Build(cfg config.CheckConfig) (Check, error) {
	switch cfg.Type {
	case "status_code":
		return StatusCodeCheck{Expected: cfg.Expected}, nil
	default:
		return nil, fmt.Errorf("unknown check type: %s", cfg.Type)
	}
}
