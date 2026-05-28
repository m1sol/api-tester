package checks

import (
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
