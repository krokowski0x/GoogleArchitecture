package storeserver

import (
	"crypto/sha256"
	"log"
	"os"
)

func SaveToDB(document []byte) [32]byte {
	docID := sha256.Sum256(document)
	outFile, err := os.Create("Repository.html")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	outFile.Write(document)
	if err != nil {
		log.Fatal(err)
	}

	return docID
}
