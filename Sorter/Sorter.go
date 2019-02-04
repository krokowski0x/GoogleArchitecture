package sorter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func SortBarrels() {
	document, err := ioutil.ReadFile("ForwardIndex.html")
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("InvertedIndex.json")
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
