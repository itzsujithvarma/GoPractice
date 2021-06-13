package main

import (
	"fmt"
	"net/http"
	"time"
)

func checklink(link string, c chan string) {
	res, err := http.Get(link)

	if err != nil {

		fmt.Println(link + " is down and returned error")
		c <- link
		return
	}
	fmt.Println(link + " is up and returned status code " + fmt.Sprint(res.StatusCode))
	c <- link
}

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://yahoo.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checklink(link, c)
	}
	for {
		func() {
			time.Sleep(time.Second)
			go checklink(<-c, c)
		}()
	}
}
