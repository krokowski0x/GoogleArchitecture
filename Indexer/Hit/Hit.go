package hit

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Hit struct {
	Word  string
	Count int
}

func Parse(r io.Reader) (map[string]int, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := hitNodes(doc)
	var words string
	for _, node := range nodes {
		words += text(node)
	}

	hits := generateHits(words)

	return hits, nil
}

func generateHits(wordsString string) map[string]int {
	hits := make(map[string]int)

	words := strings.Split(wordsString, " ")

	for _, word := range words {
		if _, ok := hits[word]; ok {
			hits[word]++
		} else {
			hits[word] = 1
		}

	}

	return hits
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

func hitNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, hitNodes(c)...)
	}
	return ret
}
