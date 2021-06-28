package datetimeformat

import (
	"time"
)

func DateFormat(inputDate string) string {
	const (
		layoutISO = "2006-01-02"
		FORMAT    = "02/01/2006"
	)
	t, _ := time.Parse(layoutISO, inputDate)
	return t.Format(FORMAT)
}

func DateTimeFormat(inputDateTime string) string {
	const (
		layoutISO = "2006-01-02 15:04:05"
		FORMAT    = "02/01/2006 15:04"
	)
	t, _ := time.Parse(layoutISO, inputDateTime)
	return t.Format(FORMAT)
}
