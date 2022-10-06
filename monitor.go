package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	url := flag.String("u", "https://google.com", "enter url")
	flag.Parse()
	now := time.Now()

	res, err := http.Get(*url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	fmt.Println(res.Request.URL)
	fmt.Println(res.Status)
	fmt.Println(time.Since(now).Seconds(), "s")
}
