package crawler

import (
	"io/ioutil"
	"net/http"
)

func Scrape(URLs []string) ([]byte, string) {
	resp, err := http.Get(URLs[1])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes, URLs[1]
}
