package main

import "fmt"

const (
	SecondsPerMinute = 60
	SecondsPerHour = SecondsPerMinute * 60
	SecondsPerDay = SecondsPerHour * 24
)

func resolvetime(seconds int)(day int, hour int, minute int)  {
	day = seconds/SecondsPerDay
	hour = seconds/SecondsPerHour
	minute = seconds/SecondsPerMinute
	return
}

func main()  {
	fmt.Println(resolvetime(1000))
	_,hour, minute := resolvetime(18000)
	fmt.Println(hour,minute)

	day,_,_ := resolvetime(90000)
	fmt.Println(day)
}
