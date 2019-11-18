package main

import (
	"fmt"
	"strconv"
)

func main() {
	var row, col int = 10, 5 //10行5列
	var seatNumber string
	for k := 0; k < row; k++ {
		for i, j := k*col, (k+1)*col-1; i < (k+1)*col && j >= k*col; i, j = i+1, j-1 {
			if k%2 == 0 {
				seatNumber += "\t" + strconv.Itoa(i)
			} else {
				seatNumber += "\t" + strconv.Itoa(j)
			}
		}
		fmt.Printf("%s\n", seatNumber)
		seatNumber = ""
	}
}
