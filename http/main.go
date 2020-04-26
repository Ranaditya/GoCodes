package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error happened, ", err)
		os.Exit(1)
	}
	lw := myWriter{}
	io.Copy(lw, resp.Body)
}

func (myWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("Number of bytes are written", len(bs))
	return len(bs), nil
}
