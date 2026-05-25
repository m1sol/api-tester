package jsonutil

import (
	"encoding/json"
	"fmt"
)

type ValueType string

const (
	TypeObject  ValueType = "object"
	TypeArray   ValueType = "array"
	TypeString  ValueType = "string"
	TypeNumber  ValueType = "number"
	TypeBool    ValueType = "bool"
	TypeNull    ValueType = "null"
	TypeUnknown ValueType = "unknown"
)

type Node struct {
	Path  string
	Type  ValueType
	Value any
}

type Document struct {
	Root      any
	Nodes     []Node
	PathIndex map[string][]*Node
}

func Parse(body []byte) (*Document, error) {
	var data any
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("parse json: %w", err)
	}

	doc := &Document{
		Root:      data,
		PathIndex: make(map[string][]*Node),
	}

	walk("$", "$", data, doc)

	return doc, nil
}

func walk(path string, indexPath string, v any, doc *Document) {
	switch val := v.(type) {
	case map[string]any:
		add(doc, indexPath, Node{Path: path, Type: TypeObject})
		for k, child := range val {
			walk(path+"."+k, indexPath+"."+k, child, doc)
		}

	case []any:
		add(doc, indexPath, Node{Path: path, Type: TypeArray})

		for i, child := range val {
			realPath := fmt.Sprintf("%s[%d]", path, i)
			wildcardPath := indexPath + "[*]"
			walk(realPath, wildcardPath, child, doc)
		}

	case nil:
		add(doc, indexPath, Node{Path: path, Type: TypeNull, Value: nil})
	case string:
		add(doc, indexPath, Node{Path: path, Type: TypeString, Value: val})
	case float64:
		add(doc, indexPath, Node{Path: path, Type: TypeNumber, Value: val})
	case bool:
		add(doc, indexPath, Node{Path: path, Type: TypeBool, Value: val})
	default:
		add(doc, indexPath, Node{Path: path, Type: TypeUnknown, Value: val})
	}
}

func add(doc *Document, indexPath string, node Node) {
	doc.Nodes = append(doc.Nodes, node)
	doc.PathIndex[indexPath] = append(doc.PathIndex[indexPath], &doc.Nodes[len(doc.Nodes)-1])
}
