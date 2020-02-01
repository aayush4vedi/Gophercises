package link

import "io"

type Link struct {
	Href string
	Text string
}

func Parse(io.Reader) ([]Link, error) {
	return nil, nil
}
