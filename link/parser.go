package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := dfs(doc)
	var links []Link
	for _, n := range nodes {
		links = append(links, buildLink(n))
	}
	return links, nil
}

func dfs(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, dfs(c)...)
	}
	return nodes
}

//get only the useful information from each node
func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	// return ret  //gives ugly string(with /n & multiple consecutive spaces)
	return strings.Join(strings.Fields(ret), " ")
}
