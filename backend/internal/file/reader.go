package file

import (
	"fmt"
	"os"
	"time"
)

func ReadFile(filePath string) (*os.File, error) {
	io, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return io, nil
}

func parseTime(timeStr string) time.Time {
	// time formatter
	layout := "2006-01-02T15:04:05.999999999Z07:00"
	parsedTime, _ := time.Parse(layout, timeStr)
	return parsedTime
}

func allValuesEqual(m map[string]int) bool {
	last := -1
	for _, v := range m {
		if last != -1 && last != v {
			return false
		}
		last = v
	}
	return true
}
