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

func values(nodes []*jsonutil.Node) []any {
	result := make([]any, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, node.Value)
	}
	return result
}

func nodeValueEquals(node *jsonutil.Node, expected any) bool {
	switch node.Type {
	case jsonutil.TypeNull:
		return expected == nil
	case jsonutil.TypeString:
		val, ok := expected.(string)
		return ok && node.Value == val
	case jsonutil.TypeNumber:
		expectedNumber, ok := numberAsFloat64(expected)
		return ok && node.Value == expectedNumber
	case jsonutil.TypeBool:
		val, ok := expected.(bool)
		return ok && node.Value == val
	default:
		return false
	}
}

func numberAsFloat64(v any) (float64, bool) {
	switch val := v.(type) {
	case int:
		return float64(val), true
	case int64:
		return float64(val), true
	case float64:
		return val, true
	default:
		return 0, false
	}
}
