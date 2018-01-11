package kv

// Read read
type Read struct{}

// NewRead new read
func NewRead() *Read {
	return &Read{}
}

// List list
func (r *Read) List(args ...string) map[string][]byte {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(false)
	defer txn.Discard()
	data := map[string][]byte{}
	for _, k := range args {
		item, err := txn.Get([]byte(k))
		if err != nil {
			continue
		}
		v, err := item.Value()
		if err != nil {
			continue
		}
		data[k] = v
	}
	return data
}

// Item item
func (r *Read) Item(k string) ([]byte, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte(k))
	if err != nil {
		return nil, err
	}
	return item.Value()
}
