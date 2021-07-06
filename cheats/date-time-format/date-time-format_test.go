package datetimeformat

import (
	"testing"
)

func TestDateFormat(t *testing.T) {
	date := "2016-11-28"
	dateFormated := DateFormat(date)
	expected := "28/11/2016"
	if dateFormated != expected {
		t.Error("Received:", dateFormated)
		t.Error("Expected:", expected)
		t.Error("your date format is wrong", dateFormated)
	}
}

func TestDateTimeFormat(t *testing.T) {
	date := "2016-11-28 18:09:02"
	dateFormated := DateTimeFormat(date)
	expected := "28/11/2016 18:09"
	if dateFormated != expected {
		t.Error("Received:", dateFormated)
		t.Error("Expected:", expected)
		t.Error("your date time format is wrong", dateFormated)
	}
}

func TestDateTimeFormatISO8601(t *testing.T) {
	date := "2021-06-02T15:04:05+0700"
	dateFormated := DateTimeFormatISO8601(date)
	expected := "02/06/2021 15:04 WIB"
	if dateFormated != expected {
		t.Error("Received:", dateFormated)
		t.Error("Expected:", expected)
		t.Error("your date time format ISO8601 is wrong", dateFormated)
	}
}

func TestDateTimeFormatAMPMTo24(t *testing.T) {
	date := "2021-06-02 04:04:00PM"
	dateFormated := DateTimeFormatTo24(date)
	expected := "02-06-2021 16:04"
	if dateFormated != expected {
		t.Error("Received:", dateFormated)
		t.Error("Expected:", expected)
		t.Error("your date time format 24 is wrong", dateFormated)
	}
}
