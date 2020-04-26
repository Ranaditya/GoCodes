package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}
	//creating channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			//delay for 2 sec
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	//fmt.Println("Now checking", link)
	_, err := http.Get(link)
	if err != nil {
		log.Println("link is down", link)
		c <- link
		return
	}
	log.Println(link, "is OK")
	c <- link
}
