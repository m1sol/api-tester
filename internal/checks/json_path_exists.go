package checks

import (
	"fmt"
	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
)

type JsonPathCheck struct {
	Path string
}

func (c JsonPathCheck) Run(
	doc *jsonutil.Document,
	response executor.Result,
) Result {
	if _, exists := doc.PathIndex[c.Path]; !exists {
		return Result{
			Type:     "json_path_exists",
			Status:   StatusFailed,
			Expected: c.Path,
			Actual:   nil,
			Message:  fmt.Sprintf("expected path %s not found", c.Path),
		}
	}

	return Result{
		Type:     "json_path_exists",
		Status:   StatusPassed,
		Expected: c.Path,
		Actual:   nil,
		Message:  "json path matched",
	}
}
