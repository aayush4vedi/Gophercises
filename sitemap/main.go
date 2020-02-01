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
	maxDepth := flag.Int("depth", 10, "the maximum number of links deep to traverse")
	flag.Parse()
	pages := bfs(*urlFlag, *maxDepth)
	for _, p := range pages {
		fmt.Println(p)
	}
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{}) //could have used `map[string]bool` here but empty structs use less memory.TODO: READ: https://dave.cheney.net/2014/03/25/the-empty-struct
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url, _ := range q {
			if _, ok := seen[url]; ok { //TODO: add in notes
				continue
			}
			seen[url] = struct{}{}
			for _, link := range getPages(url) {
				nq[link] = struct{}{} // first `{}` is type & second `{}` is instantiating it
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func getPages(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		return []string{}
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
	return filter(hrefs, withPrefix(base))
}

// 4. filter out any links w/ a diff domain
func filter(links []string, keepFn func(string) bool) []string {
	var filteredLinks []string
	for _, l := range links {
		if keepFn(l) {
			filteredLinks = append(filteredLinks, l)
		}
	}
	return filteredLinks
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
