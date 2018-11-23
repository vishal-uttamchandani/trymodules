package main

import (
	"fmt"
	"trymodules/fetcher"
)

var urls = []string{
	"http://www.amazon.com",
	"http://www.apple.com",
	"http://www.microsoft.com",
	"http://golang.org",
}

func main() {
	f := fetcher.New()

	results := f.Fetch(urls...)

	for r := range results {
		fmt.Println(r)
	}
}
