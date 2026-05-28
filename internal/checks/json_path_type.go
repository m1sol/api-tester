package checks

import (
	"fmt"

	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
)

type JsonPathTypeCheck struct {
	Expected string
	Path     string
}

func (c JsonPathTypeCheck) Run(
	doc *jsonutil.Document,
	response executor.Result,
) Result {
	nodes, exists := doc.PathIndex[c.Path]
	if !exists {
		return Result{
			Type:     "json_path_type",
			Status:   StatusFailed,
			Path:     c.Path,
			Expected: c.Expected,
			Actual:   nil,
			Message:  fmt.Sprintf("path %s not found", c.Path),
		}
	}

	actual := nodeTypes(nodes)
	for _, node := range nodes {
		if string(node.Type) != c.Expected {
			return Result{
				Type:     "json_path_type",
				Status:   StatusFailed,
				Path:     c.Path,
				Expected: c.Expected,
				Actual:   actual,
				Message:  fmt.Sprintf("json path type mismatched: expected %s, actual %v", c.Expected, actual),
			}
		}
	}

	return Result{
		Type:     "json_path_type",
		Status:   StatusPassed,
		Path:     c.Path,
		Expected: c.Expected,
		Actual:   actual,
		Message:  "json path type matched",
	}
}
