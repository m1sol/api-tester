package checks

import (
	"fmt"
	"github.com/m1sol/api-tester/internal/config"
)

func Build(cfg config.CheckConfig) (Check, error) {
	switch cfg.Type {
	case "status_code":
		expected, err := expectedInt(cfg.Expected)
		if err != nil {
			return nil, err
		}
		return StatusCodeCheck{Expected: expected}, nil
	case "json_path_exists":
		return JsonPathCheck{Path: cfg.Path}, nil
	case "json_path_count":
		expected, err := expectedInt(cfg.Expected)
		if err != nil {
			return nil, err
		}
		return JsonPathCountCheck{Path: cfg.Path, Expected: expected}, nil
	case "json_path_type":
		expected, err := expectedString(cfg.Expected)
		if err != nil {
			return nil, err
		}
		return JsonPathTypeCheck{Path: cfg.Path, Expected: expected}, nil
	case "json_path_not_empty":
		return JsonPathNotEmptyCheck{Path: cfg.Path}, nil
	case "json_path_equals":
		return JsonPathEqualsCheck{Path: cfg.Path, Expected: cfg.Expected}, nil
	default:
		return nil, fmt.Errorf("unknown check type: %s", cfg.Type)
	}
}
