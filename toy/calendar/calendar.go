package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year, month, _ := now.Date()
	d1 := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	d2 = d2.Add(-24 * time.Hour)

	weekday := d1.Weekday()
	_, week := d1.ISOWeek()
	fmt.Println("\n周\t 日 一 二 三 四 五 六  ")
	fmt.Printf("%d\t%*d", week, int((weekday+1)*3), 1)
	for i := 2; i <= d2.Day(); i++ {
		if weekday++; weekday%7 == 0 {
			week++
			fmt.Printf("\n%d\t", week)
		}
		fmt.Printf("%3d", i)
	}
}
