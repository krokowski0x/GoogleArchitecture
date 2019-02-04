package main

import (
	"fmt"

	crawler "./Crawler"
	indexer "./Indexer"
	storeserver "./StoreServer"
	urlserver "./URLServer"
)

func main() {
	urls := urlserver.GenerateURLs()
	document, url := crawler.Scrape(urls)
	docID := storeserver.SaveToDB(document)
	fmt.Println(url, docID)
	indexer.ParseAnchors()
	indexer.ParseHits()
	//urlresolver.AnchorsToLinks(url)
	//sorter.SortBarrels()
}
