package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Gophercises/link"
)

/*
	1. GET the webpage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links w/ a diff domain
	5. Find all pages (BFS)
	6. print out XML
*/
func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "link of website you want to parse for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	links, _ := link.Parse(resp.Body)
	for _, l := range links {
		fmt.Println(l)
	}
	/*
		Possible links:
		A. partial links: `/path`
		B. full links: starting with `http` or `https`
		C. Anything else-not to be included in the list: fragments-`#signup`, `mailto@xyz.com`
	*/
	//for A.
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	fmt.Printf("%+v\n", base)
}
