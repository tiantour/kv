package kv

import (
	"log"
	"runtime"

	"github.com/dgraph-io/badger"
	"github.com/tiantour/conf"
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
	c := conf.NewConf().KV
	if c.Path == "" {
		c.Path = "/tmp/badger"
	}
	opts := badger.DefaultOptions
	opts.Dir = c.Path
	opts.ValueDir = c.Path
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
