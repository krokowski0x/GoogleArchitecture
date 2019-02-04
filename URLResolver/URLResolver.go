package urlresolver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./link"
)

func ParseLinks(url string) {
	document, err := ioutil.ReadFile("Anchors.json")
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("Links.json")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	var anchors []Anchor
	err := json.Unmarshal(document, &anchors)
	if err != nil {
		fmt.Println("error:", err)
	}

	links, err := link.Parse(anchors)
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(links)

	outFile.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
