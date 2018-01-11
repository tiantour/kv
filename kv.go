package kv

import (
	"log"
	"runtime"

	"github.com/dgraph-io/badger"
)

var (
	conn chan int
	po   *pool
)

// pool pool
type pool struct {
	*badger.DB
}

func init() {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	cap := runtime.NumCPU()
	conn = make(chan int, cap)
	for i := 0; i < cap; i++ {
		conn <- 1
	}
	po = &pool{db}
}
