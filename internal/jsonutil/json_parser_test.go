package jsonutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseValidJSON(t *testing.T) {
	body := []byte(`{"data":{"total":100}}`)

	doc, err := Parse(body)

	require.NoError(t, err)
	require.NotNil(t, doc.Root)
}

func TestParseInvalidJSON(t *testing.T) {
	body := []byte(`{"data":`)

	doc, err := Parse(body)

	require.Error(t, err)
	require.Nil(t, doc)
	assert.Contains(t, err.Error(), "parse json")
}

func TestWalkNestedObject(t *testing.T) {
	body := []byte(`{"data":{"total":{"revenue":100}}}`)

	doc, err := Parse(body)

	require.NoError(t, err)

	assert.Equal(t, []Node{
		{Path: "$", Type: TypeObject},
		{Path: "$.data", Type: TypeObject},
		{Path: "$.data.total", Type: TypeObject},
		{Path: "$.data.total.revenue", Type: TypeNumber, Value: 100.0},
	}, doc.Nodes)
}

func TestWalkArrayObject(t *testing.T) {
	body := []byte(`{
		"data": {
			"total": {
				"revenue": 100
			},
			"rows": [
				{"revenue": 10},
				{"revenue": 20}
			]
		}
	}`)

	doc, err := Parse(body)

	require.NoError(t, err)

	assert.Equal(t, []Node{
		{
			Path: "$",
			Type: TypeObject,
		},
		{
			Path: "$.data",
			Type: TypeObject,
		},
		{
			Path: "$.data.total",
			Type: TypeObject,
		},
		{
			Path:  "$.data.total.revenue",
			Type:  TypeNumber,
			Value: float64(100),
		},
		{
			Path: "$.data.rows",
			Type: TypeArray,
		},
		{
			Path: "$.data.rows[0]",
			Type: TypeObject,
		},
		{
			Path:  "$.data.rows[0].revenue",
			Type:  TypeNumber,
			Value: float64(10),
		},
		{
			Path: "$.data.rows[1]",
			Type: TypeObject,
		},
		{
			Path:  "$.data.rows[1].revenue",
			Type:  TypeNumber,
			Value: float64(20),
		},
	}, doc.Nodes)

	require.Contains(t, doc.PathIndex, "$.data.rows[*].revenue")

	assert.Len(t, doc.PathIndex["$.data.rows[*].revenue"], 2)

	assert.Equal(
		t,
		float64(10),
		doc.PathIndex["$.data.rows[*].revenue"][0].Value,
	)

	assert.Equal(
		t,
		float64(20),
		doc.PathIndex["$.data.rows[*].revenue"][1].Value,
	)
}
