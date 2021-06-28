package datetimeformat

import (
	"testing"
)

func TestDateFormat(t *testing.T) {
	date := "2016-11-28"
	dateFormated := DateFormat(date)
	if dateFormated != "28/11/2016" {
		t.Error("your date format is wrong", dateFormated)
	}
}

func TestDateTimeFormat(t *testing.T) {
	date := "2016-11-28 18:09:02"
	dateFormated := DateTimeFormat(date)
	if dateFormated != "28/11/2016 18:09" {
		t.Error("your date time format is wrong", dateFormated)
	}
}
