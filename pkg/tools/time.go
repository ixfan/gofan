package tools

import "time"

// Now 当前时间
func Now() *time.Time {
	mTime := time.Now()
	return &mTime
}
