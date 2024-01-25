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
	var workTimeHours int = endTimeHour - startTimeHour
	var workTimeMinutes int = endTimeMinute - startTimeMinute

	var workTime = workTimeHours*60 + workTimeMinutes
	if (workTime >= 481) && (restTime < 60) {
		return "60分未満の休憩登録は不可"
	} else if (workTime >= 361) && (restTime < 45) {
		return "45分未満の休憩登録は不可"
	} else {
		return ""
	}
}
