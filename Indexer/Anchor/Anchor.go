package anchor

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Anchor struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Anchor, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := anchorNodes(doc)
	var anchors []Anchor
	for _, node := range nodes {
		anchors = append(anchors, buildAnchor(node))
	}
	return anchors, nil
}

func buildAnchor(n *html.Node) Anchor {
	var ret Anchor
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
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func anchorNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, anchorNodes(c)...)
	}
	return ret
}
