package report

import (
	"fmt"
	"github.com/m1sol/api-tester/internal/checks"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func logWithColor(status checks.Status, format string, args ...any) {
	color := colorByStatus(status)
	level := strings.ToUpper(string(status))

	msg := fmt.Sprintf(format, args...)

	fmt.Printf("%s[%s]:%s %s\n",
		color,
		level,
		Reset,
		msg,
	)
}

func colorByStatus(status checks.Status) string {
	switch status {
	case checks.StatusPassed:
		return Green
	case checks.StatusSkipped:
		return Yellow
	case checks.StatusFailed:
		return Red
	default:
		return Reset
	}
}

func Print(results []checks.Result) {
	totalChecks := make(map[checks.Status]int)
	for _, res := range results {
		totalChecks[res.Status]++
		logWithColor(
			res.Status,
			"type=%s path=%s expected=%v actual=%v message=%s",
			res.Type,
			res.Path,
			res.Expected,
			res.Actual,
			res.Message,
		)
	}
	statuses := []checks.Status{
		checks.StatusPassed,
		checks.StatusSkipped,
		checks.StatusFailed,
	}
	fmt.Println("\nSummary:")
	for _, status := range statuses {
		count := totalChecks[status]

		logWithColor(status, "count=%d", count)
	}

}
