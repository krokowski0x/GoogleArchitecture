package indexer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"./anchor"
	"./hit"
)

func ParseAnchors() {
	document, err := ioutil.ReadFile("Repository.html")
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("Anchors.json")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	r := strings.NewReader(string(document))

	anchors, err := anchor.Parse(r)
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(anchors)

	outFile.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

func ParseHits() {
	document, err := ioutil.ReadFile("Repository.html")
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("ForwardIndex.json")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	r := strings.NewReader(string(document))

	hits, err := hit.Parse(r)
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(hits)

	outFile.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
