package checks

import (
	"fmt"
	"math"
	"strconv"
)

func expectedInt(v any) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case int64:
		return int(val), nil
	case float64:
		if math.Trunc(val) != val {
			return 0, fmt.Errorf("expected integer, got %v", val)
		}
		return int(val), nil
	case string:
		i, err := strconv.Atoi(val)
		if err != nil {
			return 0, fmt.Errorf("expected integer, got %q", val)
		}
		return i, nil
	default:
		return 0, fmt.Errorf("expected integer, got %T", v)
	}
}

func expectedString(v any) (string, error) {
	if val, ok := v.(string); ok {
		return val, nil
	}

	return "", fmt.Errorf("expected string, got %T", v)
}
