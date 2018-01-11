package kv

import "github.com/dgraph-io/badger"

// Key key
type Key struct{}

// NewKey new key
func NewKey() *Key {
	return &Key{}
}

// All all
func (k *Key) All(size int) map[string][]byte {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(false)
	defer txn.Discard()
	opts := badger.DefaultIteratorOptions
	opts.PrefetchSize = size
	it := txn.NewIterator(opts)
	data := map[string][]byte{}
	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()
		k := item.Key()
		v, err := item.Value()
		if err != nil {
			continue
		}
		data[string(k)] = v
	}
	return data
}

// Keys keys
func (k *Key) Keys(size int) []*string {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(false)
	defer txn.Discard()
	opts := badger.DefaultIteratorOptions
	opts.PrefetchSize = size
	opts.PrefetchValues = false
	it := txn.NewIterator(opts)
	var data []*string
	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()
		k := item.Key()
		temp := string(k)
		data = append(data, &temp)
	}
	return data
}

// Prefix prefix
func (k *Key) Prefix(key string, size int) map[string][]byte {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(false)
	defer txn.Discard()
	opts := badger.DefaultIteratorOptions
	opts.PrefetchSize = size
	it := txn.NewIterator(opts)
	data := map[string][]byte{}
	for it.Seek([]byte(key)); it.ValidForPrefix([]byte(key)); it.Next() {
		item := it.Item()
		k := item.Key()
		v, err := item.Value()
		if err != nil {
			continue
		}
		data[string(k)] = v
	}
	return data
}

// Delete delete
func (k *Key) Delete(args ...string) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(true)
	defer txn.Discard()
	for _, k := range args {
		err := txn.Delete([]byte(k))
		if err != nil {
			continue
		}
	}
	return txn.Commit(nil)
}

// Unique unique
func (k *Key) Unique(key string, width uint64) (uint64, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	seq, err := po.DB.GetSequence([]byte(key), width)
	if err != nil {
		return 0, err
	}
	defer seq.Release()
	return seq.Next()
}
