package util

import "time"

func GetTodayWithTime() (startDay time.Time, endDay time.Time) {
	year, month, day := time.Now().Date()
	startDay = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	endDay = time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
	return
}
