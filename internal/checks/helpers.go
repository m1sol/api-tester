package checks

import (
	"github.com/m1sol/api-tester/internal/jsonutil"
	"strings"
)

func nodeTypes(nodes []*jsonutil.Node) []string {
	types := make([]string, 0, len(nodes))
	for _, node := range nodes {
		types = append(types, string(node.Type))
	}
	return types
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
