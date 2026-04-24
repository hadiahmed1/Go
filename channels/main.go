package main

import (
	"fmt"
	"net/http"
	"time"
)

type message struct {
	msg  string
	link string
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.com",
		"http://golang.org",
		"http://example.com",
		"http://exxxxample.com",
	}

	c := make(chan message)

	for _, l := range links {
		go checkLink(l, c)
	}
	for l := range c {
		go func (l message )  {
			time.Sleep(5 * time.Second)
			res := l // blocking call
			fmt.Println(res.msg)
			go checkLink(res.link, c)
		}(l)
	}
}

func checkLink(link string, c chan message) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link + " is Down")
		c <- message{msg: link + " is down", link: link}
	} else {
		// fmt.Println(link + " is UP & Running")
		c <- message{msg: link + " is UP & Running", link: link}
	}
}
