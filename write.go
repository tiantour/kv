package kv

import (
	"time"
)

// Write write
type Write struct{}

// NewWrite new write
func NewWrite() *Write {
	return &Write{}
}

// List list
func (w *Write) List(args map[string][]byte, ttl time.Duration) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(true)
	defer txn.Discard()
	for k, v := range args {
		var err error
		if ttl != 0 {
			err = txn.SetWithTTL([]byte(k), v, ttl)
		} else {
			err = txn.Set([]byte(k), v)
		}
		if err != nil {
			return err
		}
	}
	return txn.Commit(nil)
}

// Item item
func (w *Write) Item(k string, v []byte, ttl time.Duration) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(true)
	defer txn.Discard()
	var err error
	if ttl != 0 {
		err = txn.SetWithTTL([]byte(k), v, ttl)
	} else {
		err = txn.Set([]byte(k), v)
	}
	if err != nil {
		return err
	}
	return txn.Commit(nil)
}
