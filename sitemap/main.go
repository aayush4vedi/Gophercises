package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

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
	pages := getPages(*urlFlag)
	for _, p := range pages {
		fmt.Println(p)
	}
}

func getPages(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return getHrefs(resp.Body, base)
}

func getHrefs(r io.Reader, base string) []string {
	var hrefs []string
	links, _ := link.Parse(r)
	/*
		Possible links:
		A. partial links: `/path`
		B. full links: starting with `http` or `https`
		C. Anything else-not to be included in the list: fragments-`#signup`, `mailto@xyz.com`
	*/
	for _, l := range links {
		switch {
		//for A.
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		//for B.
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	return hrefs
}
