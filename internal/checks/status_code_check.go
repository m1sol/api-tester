package checks

import (
	"fmt"
	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
)

type StatusCodeCheck struct {
	Expected int
}

func (c StatusCodeCheck) Run(
	doc *jsonutil.Document,
	response executor.Result,
) Result {
	if response.StatusCode != c.Expected {
		return Result{
			Type:     "status_code",
			Status:   StatusFailed,
			Expected: c.Expected,
			Actual:   response.StatusCode,
			Message:  fmt.Sprintf("expected status %d, got %d", c.Expected, response.StatusCode),
		}
	}

	return Result{
		Type:     "status_code",
		Status:   StatusPassed,
		Expected: c.Expected,
		Actual:   response.StatusCode,
		Message:  "status code matched",
	}
}
