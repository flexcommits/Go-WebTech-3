package main

import (
	"fmt"
)

func main() {

	var (
		startTimeHour   int
		startTimeMinute int
		endTimeHour     int
		endTimeMinute   int
		restTime        int
	)

	fmt.Println("Enter start time")
	_, err := fmt.Scanf("%d:%d %d:%d %d", &startTimeHour, &startTimeMinute, &endTimeHour, &endTimeMinute, &restTime)

	if err == nil {
		fmt.Println(ifValid(startTimeHour, startTimeMinute, endTimeHour, endTimeMinute, restTime))
	} else {
		fmt.Println("Enter correctly")
	}
}

func ifValid(startTimeHour, startTimeMinute, endTimeHour, endTimeMinute, restTime int) string {
	workTime := (endTimeHour-startTimeHour)*60 + (endTimeMinute - startTimeMinute)

	var minRestTime int
	switch {
	case workTime >= 481:
		minRestTime = 60
	case workTime >= 361:
		minRestTime = 45
	default:
		minRestTime = 0
	}

	if restTime < minRestTime {
		return fmt.Sprintf("%d分未満の休憩登録は不可", minRestTime)
	} else {
		return ""
	}
}
