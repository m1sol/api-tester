package checks

import (
	"fmt"
	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
)

type JsonPathCountCheck struct {
	Expected int
	Path     string
}

func (c JsonPathCountCheck) Run(
	doc *jsonutil.Document,
	response executor.Result,
) Result {
	nodes, exists := doc.PathIndex[c.Path]
	if !exists {
		return Result{
			Type:     "json_path_count",
			Status:   StatusFailed,
			Path:     c.Path,
			Expected: c.Expected,
			Actual:   nil,
			Message:  fmt.Sprintf("path %s not found", c.Path),
		}
	}

	actual := len(nodes)
	if actual != c.Expected {
		return Result{
			Type:     "json_path_count",
			Status:   StatusFailed,
			Path:     c.Path,
			Expected: c.Expected,
			Actual:   actual,
			Message:  fmt.Sprintf("json path count mismatched: expected %d, actual %d", c.Expected, actual),
		}
	}

	return Result{
		Type:     "json_path_count",
		Status:   StatusPassed,
		Path:     c.Path,
		Expected: c.Expected,
		Actual:   actual,
		Message:  "json path count matched",
	}
}
