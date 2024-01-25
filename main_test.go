package main

import "testing"

type validTest struct {
	startTimeHour, startTimeMinute, endTimeHour, endTimeMinute, restTime int
	expected                                                             string
}

var validTests = []validTest{
	{2, 3, 8, 4, 34, "45分未満の休憩登録は不可"},
	{2, 3, 10, 4, 59, "60分未満の休憩登録は不可"},
	{2, 3, 4, 4, 34, ""},
}

func TestIfValid(t *testing.T) {

	for _, test := range validTests {
		if output := ifValid(test.startTimeHour, test.startTimeMinute, test.endTimeHour, test.endTimeMinute, test.restTime); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
