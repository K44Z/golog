package golog

import (
	"fmt"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	Bold   = "\033[1m"
)

func getSpeedColor(duration time.Duration) string {
	switch {
	case duration < 50*time.Millisecond:
		return Green
	case duration < 200*time.Millisecond:
		return Yellow
	case duration < 500*time.Millisecond:
		return Red
	default:
		return Red + Bold
	}
}

func getStatusColor(status int) string {
	switch {
	case status >= 500:
		return Red + Bold
	case status >= 400:
		return Yellow
	case status >= 300:
		return Blue
	case status >= 200:
		return Green
	default:
		return Gray
	}
}

func getMethodColor(method string) string {
	switch method {
	case "GET":
		return Blue
	case "POST":
		return Green
	case "PUT":
		return Yellow
	case "DELETE":
		return Red
	case "PATCH":
		return Purple
	default:
		return Gray
	}
}

func formatDuration(d time.Duration) string {
	switch {
	case d < time.Microsecond:
		return fmt.Sprintf("%.0fns", float64(d.Nanoseconds()))
	case d < time.Millisecond:
		return fmt.Sprintf("%.0fµs", float64(d.Nanoseconds())/1000)
	case d < time.Second:
		return fmt.Sprintf("%.2fms", float64(d.Nanoseconds())/1e6)
	default:
		return fmt.Sprintf("%.2fs", d.Seconds())
	}
}
