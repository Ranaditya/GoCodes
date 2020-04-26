package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args
	logger := log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
	if len(args) < 2 {
		logger.Println("Please provide file name in first argument")
		os.Exit(1)
	}
	logger.Println("File name", args[1])
	f, err := os.Open(args[1])
	if err != nil {
		logger.Println("Error is reading file", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
