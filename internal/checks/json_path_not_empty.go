package checks

import (
	"fmt"
	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
)

type JsonPathNotEmptyCheck struct {
	Path string
}

func (c JsonPathNotEmptyCheck) Run(
	doc *jsonutil.Document,
	response executor.Result,
) Result {
	nodes, exists := doc.PathIndex[c.Path]
	if !exists {
		return Result{
			Type:     "json_path_not_empty",
			Status:   StatusFailed,
			Path:     c.Path,
			Expected: "not empty",
			Actual:   nil,
			Message:  fmt.Sprintf("path %s not found", c.Path),
		}
	}

	for _, node := range nodes {
		if nodeIsEmpty(doc, node) {
			return Result{
				Type:     "json_path_not_empty",
				Status:   StatusFailed,
				Path:     c.Path,
				Expected: "not empty",
				Actual:   node.Value,
				Message:  fmt.Sprintf("json path %s is empty", node.Path),
			}
		}
	}

	return Result{
		Type:     "json_path_not_empty",
		Status:   StatusPassed,
		Path:     c.Path,
		Expected: "not empty",
		Actual:   values(nodes),
		Message:  "json path is not empty",
	}
}
