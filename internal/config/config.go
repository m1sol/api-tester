package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Suite struct {
	Request RequestConfig `yaml:"request"`
	Checks  []CheckConfig `yaml:"checks"`
}

type RequestConfig struct {
	Method  string  `yaml:"method"`
	URL     string  `yaml:"url"`
	Headers Headers `yaml:"headers"`
	Body    Body    `yaml:"body"`
}

type CheckConfig struct {
	Type        string `yaml:"type"`
	Expected    any    `yaml:"expected,omitempty"`
	Path        string `yaml:"path,omitempty"`
	Left        string `yaml:"left,omitempty"`
	Right       string `yaml:"right,omitempty"`
	Aggregation string `yaml:"aggregation,omitempty"`
}

type Headers struct {
	Authorization string `yaml:"Authorization"`
	ContentType   string `yaml:"Content-Type"`
}

type Body struct {
	Type           string  `yaml:"type" json:"type"`
	Step           string  `yaml:"step" json:"step"`
	OrganizationId string  `yaml:"organizationId" json:"organizationId"`
	StartDate      string  `yaml:"startDate" json:"startDate"`
	EndDate        string  `yaml:"endDate" json:"endDate"`
	Fields         []Field `yaml:"fields" json:"fields"`
}

type Field struct {
	ID       string `yaml:"id" json:"id"`
	MetricID string `yaml:"metricId" json:"metricId"`
	Sort     int    `yaml:"sort" json:"sort"`
}

func LoadYaml(path string) (*Suite, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dataStr := string(data)
	dataStr = ResolveEnv(dataStr)
	data = []byte(dataStr)
	var cfg Suite

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg, nil
}

func ResolveEnv(s string) string {
	for {
		start := strings.Index(s, "${")
		if start == -1 {
			break
		}
		end := strings.Index(s[start:], "}")
		if end == -1 {
			break
		}
		end += start
		key := s[start+2 : end]
		val := os.Getenv(key)
		s = s[:start] + val + s[end+1:]
	}
	return s
}
