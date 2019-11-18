package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	_, err := readSlice("sw1.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func readSlice(infile string) (slice []string, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		slice = make([]string, 0)
		for i := 0; i < 14; i++ {
			line, isPrefix, err1 := br.ReadLine()
			if err1 != nil {
				if err1 != io.EOF {
					err = err1
				}
				break
			}
			if isPrefix {
				fmt.Println("A too long line,seems unexpected.")
				return
			}
			str := string(line)
			slice = append(slice, str)
		}
		tip, err2 := time.ParseDuration(slice[0] + "ms")
		if err2 != nil {
			err = err2
			break
		}
		for j := 1; j < len(slice); j++ {
			fmt.Println(slice[j])
		}
		time.Sleep(tip * 100)
		cmd := exec.Command("cls")
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
