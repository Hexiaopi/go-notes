package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:./curl <url> <file>")
		os.Exit(-1)
	}
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
	defer file.Close()
	dest := io.MultiWriter(os.Stdout, file)
	io.Copy(dest, resp.Body)
	if err := resp.Body.Close(); err != nil {
		log.Println("ERROR:", err)
	}
}
