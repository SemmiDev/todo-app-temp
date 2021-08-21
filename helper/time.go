package helper

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func timeFormat(hour int, minutes int) string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minutes, time.Now().Second(), time.Now().Nanosecond(), loc).
		Format(time.RFC850)
}

func TimeNow() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), loc).
		Format(time.RFC850)
}

func TaskTimeFormat(hour, minutes int) string {
	return timeFormat(hour, minutes)
}

func TaskDuration(startingAt, endsAt string) float64 {
	start, _ := time.Parse(time.RFC850, startingAt)
	ends, _ := time.Parse(time.RFC850, endsAt)
	return ends.Sub(start).Minutes()
}

func IsTaskExpired(endTask string) bool {
	now, _ := time.Parse(time.RFC850, TimeNow())
	end, _ := time.Parse(time.RFC850, endTask)
	return end.Sub(now).Minutes() < 0
}

func ExtractTodoTime(startingAt string, endsAt string) (starting string, ends string, duration float64, error error) {
	exStarting := strings.Split(startingAt, ":")
	exEnds := strings.Split(endsAt, ":")

	if len(exStarting) != 2 || len(exEnds) != 2 {
		return "", "", 0, errors.New("time format is not valid")
	}

	hour, err := strconv.Atoi(exStarting[0])
	if err != nil {
		return "", "", 0, errors.New("time format is not valid")
	}
	minute, err := strconv.Atoi(exStarting[1])
	if err != nil {
		return "", "", 0, errors.New("time format is not valid")
	}
	starting = TaskTimeFormat(hour, minute)

	hour, err = strconv.Atoi(exEnds[0])
	if err != nil {
		return "", "", 0, errors.New("time format is not valid")
	}

	minute, err = strconv.Atoi(exEnds[1])
	if err != nil {
		return "", "", 0, errors.New("time format is not valid")
	}
	ends = TaskTimeFormat(hour, minute)

	duration = TaskDuration(starting, ends)
	return
}
