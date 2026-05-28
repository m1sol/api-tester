package checks

import (
	"fmt"
	"strings"

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

func nodeIsEmpty(doc *jsonutil.Document, node *jsonutil.Node) bool {
	switch node.Type {
	case jsonutil.TypeNull:
		return true
	case jsonutil.TypeString:
		return node.Value == ""
	case jsonutil.TypeArray, jsonutil.TypeObject:
		return !hasChildNode(doc, node.Path)
	default:
		return false
	}
}

func hasChildNode(doc *jsonutil.Document, path string) bool {
	for _, node := range doc.Nodes {
		if node.Path == path {
			continue
		}
		if strings.HasPrefix(node.Path, path+".") || strings.HasPrefix(node.Path, path+"[") {
			return true
		}
	}

	return false
}
