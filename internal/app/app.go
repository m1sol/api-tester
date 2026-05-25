package app

import (
	"context"
	"fmt"
	"github.com/m1sol/api-tester/internal/config"
	"github.com/m1sol/api-tester/internal/executor"
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

	result := exec.Execute(ctx, suite.Request)

	fmt.Printf("SUITE: %+v\n", result)
	return nil
}
