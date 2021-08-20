package helper

import (
	"strconv"
	"strings"
	"time"
)

func timeFormat(day int, hour int, minutes int) string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(time.Now().Year(), time.Now().Month(), day, hour, minutes, 0, 0, loc).
		Format(time.RFC822)
}

func TimeNow() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), loc).
		Format(time.RFC822)
}

func TaskTimeFormat(day, hour, minutes int) string {
	return timeFormat(day, hour, minutes)
}

func TaskDuration(startingAt, endsAt string) float64 {
	start, _ := time.Parse(time.RFC822, startingAt)
	ends, _ := time.Parse(time.RFC822, endsAt)
	return ends.Sub(start).Minutes()
}

func IsTaskExpired(endTask string) bool {
	timenow, _ := time.Parse(time.RFC822, TimeNow())
	timetask, _ := time.Parse(time.RFC822, endTask)
	return timetask.Sub(timenow).Minutes() < 0
}

func ExtractTodoTime(startingAt string, endsAt string) (starting string, ends string, duration float64) {
	exStarting := strings.Split(startingAt, ":")
	exEnds := strings.Split(endsAt, ":")

	day, _ := strconv.Atoi(exStarting[0])
	hour, _ := strconv.Atoi(exStarting[1])
	minute, _ := strconv.Atoi(exStarting[2])
	starting = TaskTimeFormat(day, hour, minute)

	day, _ = strconv.Atoi(exEnds[0])
	hour, _ = strconv.Atoi(exEnds[1])
	minute, _ = strconv.Atoi(exEnds[2])
	ends = TaskTimeFormat(day, hour, minute)

	duration = TaskDuration(starting, ends)
	return
}
