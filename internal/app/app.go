package app

import (
	"context"
	"fmt"
	"github.com/m1sol/api-tester/internal/checks"
	"github.com/m1sol/api-tester/internal/config"
	"github.com/m1sol/api-tester/internal/executor"
	"github.com/m1sol/api-tester/internal/jsonutil"
	"github.com/m1sol/api-tester/internal/report"
	"os"
)

func Run() error {
	ctx := context.Background()

	path := "examples/aggregator.yaml"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	suite, err := config.LoadYaml(path)
	if err != nil {
		return fmt.Errorf("load suite: %w", err)
	}

	exec := executor.Executor{}

	response := exec.Execute(ctx, suite.Request)
	if response.Err != nil {
		return fmt.Errorf("execute request: %w", response.Err)
	}

	doc, err := jsonutil.Parse(response.Body)
	if err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	results := make([]checks.Result, 0, len(suite.Checks))

	for _, cfg := range suite.Checks {
		check, err := checks.Build(cfg)
		if err != nil {
			results = append(results, checks.Result{
				Type:     cfg.Type,
				Status:   checks.StatusSkipped,
				Message:  fmt.Sprintf("check is not supported yet: %v", err),
				Expected: cfg.Type,
			})
			continue
		}

		result := check.Run(doc, response)
		results = append(results, result)
	}

	report.Print(results)

	return nil
}
