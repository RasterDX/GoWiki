package db

import (
	"github.com/dgraph-io/badger"
	"log"
)

func CreateArticles() {
	// Open or create Badger database
	opts := badger.DefaultOptions
	opts.Dir = "../../tmp/articles"
	opts.ValueDir = "../../tmp/articles"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Your code here…
}
