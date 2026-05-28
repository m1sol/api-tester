package checks

import (
	"fmt"

	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
)

type JsonPathEqualsCheck struct {
	Expected any
	Path     string
}

func (c JsonPathEqualsCheck) Run(
	doc *jsonutil.Document,
	response executor.Result,
) Result {
	nodes, exists := doc.PathIndex[c.Path]
	if !exists {
		return Result{
			Type:     "json_path_equals",
			Status:   StatusFailed,
			Path:     c.Path,
			Expected: c.Expected,
			Actual:   nil,
			Message:  fmt.Sprintf("path %s not found", c.Path),
		}
	}

	actual := values(nodes)
	for _, node := range nodes {
		if !nodeValueEquals(node, c.Expected) {
			return Result{
				Type:     "json_path_equals",
				Status:   StatusFailed,
				Path:     c.Path,
				Expected: c.Expected,
				Actual:   actual,
				Message:  fmt.Sprintf("json path value mismatched: expected %v, actual %v", c.Expected, actual),
			}
		}
	}

	return Result{
		Type:     "json_path_equals",
		Status:   StatusPassed,
		Path:     c.Path,
		Expected: c.Expected,
		Actual:   actual,
		Message:  "json path value matched",
	}
}
