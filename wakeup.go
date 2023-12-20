package wakeup

import (
	"fmt"
	"log"
	"time"
)

const (
	hourMinute       string = "%d:%d"
	hourMinuteSecond string = "%d:%d:%d"
)

func Parse(what string) (hour, minute, second int, err error) {
	if _, err = fmt.Sscanf(what, hourMinuteSecond, &hour, &minute, &second); err == nil {
		return
	}
	if _, err = fmt.Sscanf(what, hourMinute, &hour, &minute); err == nil {
		return
	}
	return
}

func SleepFrom(what string, logger ...*log.Logger) error {
	hour, minute, second, err := Parse(what)
	if err != nil {
		return err
	}
	Sleep(hour, minute, second, logger...)
	return nil
}

func Sleep(hour int, minute int, second int, logger ...*log.Logger) {
	t := Time(hour, minute, second)
	for _, log := range logger {
		log.Printf("sleeping until %s | %s", t.Format(time.DateTime), time.Until(t).String())
	}
	time.Sleep(time.Until(t))
}

func Duration(hour int, minute int, second int) time.Duration {
	return duration(time.Now(), hour, minute, second)
}

func Time(hour int, minute int, second int) time.Time {
	now := time.Now()
	return now.Add(duration(now, hour, minute, second))
}

func duration(now time.Time, hour int, minute int, second int) time.Duration {
	duration := time.Second * time.Duration(
		(hour-now.Hour())*60*60+
			(minute-now.Minute())*60+
			(second-now.Second()),
	)
	if duration < 0 {
		duration += 24 * time.Hour
	}
	return duration
}
