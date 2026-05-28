package checks

import (
	"testing"

	"github.com/m1sol/api-tester/internal/config"
	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildSupportsMixedExpectedTypes(t *testing.T) {
	tests := []struct {
		name string
		cfg  config.CheckConfig
	}{
		{
			name: "status code",
			cfg:  config.CheckConfig{Type: "status_code", Expected: 200},
		},
		{
			name: "json path count",
			cfg:  config.CheckConfig{Type: "json_path_count", Path: "$.items[*]", Expected: 2},
		},
		{
			name: "json path type",
			cfg:  config.CheckConfig{Type: "json_path_type", Path: "$.status", Expected: "string"},
		},
		{
			name: "json path equals",
			cfg:  config.CheckConfig{Type: "json_path_equals", Path: "$.status", Expected: "ok"},
		},
		{
			name: "json path not empty",
			cfg:  config.CheckConfig{Type: "json_path_not_empty", Path: "$.items"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check, err := Build(tt.cfg)

			require.NoError(t, err)
			require.NotNil(t, check)
		})
	}
}

func TestJsonPathTypeCheck(t *testing.T) {
	doc := mustParse(t, `{"items":[{"value":10},{"value":20}],"status":"ok"}`)

	passed := JsonPathTypeCheck{
		Path:     "$.items[*].value",
		Expected: "number",
	}.Run(doc, executor.Result{})

	assert.Equal(t, StatusPassed, passed.Status)
	assert.Equal(t, []string{"number", "number"}, passed.Actual)

	failed := JsonPathTypeCheck{
		Path:     "$.status",
		Expected: "number",
	}.Run(doc, executor.Result{})

	assert.Equal(t, StatusFailed, failed.Status)
	assert.Equal(t, []string{"string"}, failed.Actual)
}

func TestJsonPathNotEmptyCheck(t *testing.T) {
	doc := mustParse(t, `{
		"items": [{"value": 10}],
		"emptyItems": [],
		"name": "report",
		"emptyName": "",
		"nullable": null
	}`)

	tests := []struct {
		name   string
		path   string
		status Status
	}{
		{name: "array with item", path: "$.items", status: StatusPassed},
		{name: "non empty string", path: "$.name", status: StatusPassed},
		{name: "empty array", path: "$.emptyItems", status: StatusFailed},
		{name: "empty string", path: "$.emptyName", status: StatusFailed},
		{name: "null", path: "$.nullable", status: StatusFailed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JsonPathNotEmptyCheck{Path: tt.path}.Run(doc, executor.Result{})

			assert.Equal(t, tt.status, result.Status)
		})
	}
}

func TestJsonPathEqualsCheck(t *testing.T) {
	doc := mustParse(t, `{
		"status": "ok",
		"count": 2,
		"enabled": true,
		"items": [{"value": 10}, {"value": 10}]
	}`)

	tests := []struct {
		name     string
		path     string
		expected any
		status   Status
	}{
		{name: "string matched", path: "$.status", expected: "ok", status: StatusPassed},
		{name: "number matched", path: "$.count", expected: 2, status: StatusPassed},
		{name: "bool matched", path: "$.enabled", expected: true, status: StatusPassed},
		{name: "wildcard values matched", path: "$.items[*].value", expected: 10, status: StatusPassed},
		{name: "string mismatched", path: "$.status", expected: "error", status: StatusFailed},
		{name: "wildcard mismatched", path: "$.items[*].value", expected: 20, status: StatusFailed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JsonPathEqualsCheck{Path: tt.path, Expected: tt.expected}.Run(doc, executor.Result{})

			assert.Equal(t, tt.status, result.Status)
		})
	}
}

func mustParse(t *testing.T, body string) *jsonutil.Document {
	t.Helper()

	doc, err := jsonutil.Parse([]byte(body))
	require.NoError(t, err)

	return doc
}
