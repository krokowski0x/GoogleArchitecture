package link

type Link struct {
	To   string
	From string
}

func Parse(anchors []Anchor, url string) ([]Link, error) {
	var links []Link
	for _, anchor := range anchors {
		links = append(links, Link{
			To:   url + anchor.Href,
			From: url,
		})
	}
	return links, nil
}
