package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(dayOfTheWeek(8, 3, 1992))
}

func dayOfTheWeek(day int, month int, year int) string {
	datefmt := "2/1/2006"
	date := strconv.Itoa(day) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(year)
	t, _ := time.Parse(datefmt, date)

	return t.Weekday().String()
}
