package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://www.google.com", "link of website you want to parse for")
	flag.Parse()

	fmt.Println(*urlFlag)

	/*
	1. GET the webpage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links w/ a diff domain
	5. Find all pages (BFS)
	6. print out XML
	*/
}
