package util

import (
	"time"
)

func GetCurrentTime(layout string) string {
	t := time.Now().Local()
	// const layout = "2006-01-02"
	return t.Format(layout)
}
