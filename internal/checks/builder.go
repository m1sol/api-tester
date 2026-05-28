package checks

import (
	"fmt"
	"github.com/m1sol/api-tester/internal/config"
)

func Build(cfg config.CheckConfig) (Check, error) {
	switch cfg.Type {
	case "status_code":
		return StatusCodeCheck{Expected: cfg.Expected}, nil
	case "json_path_exists":
		return JsonPathCheck{Path: cfg.Path}, nil
	case "json_path_count":
		return JsonPathCountCheck{Path: cfg.Path, Expected: cfg.Expected}, nil
	default:
		return nil, fmt.Errorf("unknown check type: %s", cfg.Type)
	}
}
