package kv

// Write write
type Write struct{}

// NewWrite new write
func NewWrite() *Write {
	return &Write{}
}

// List list
func (w *Write) List(args map[string][]byte) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(true)
	defer txn.Discard()
	for k, v := range args {
		err := txn.Set([]byte(k), v)
		if err != nil {
			return err
		}
	}
	return txn.Commit(nil)
}

// Item item
func (w *Write) Item(k string, v []byte) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	txn := po.DB.NewTransaction(true)
	defer txn.Discard()
	err := txn.Set([]byte(k), v)
	if err != nil {
		return err
	}
	return txn.Commit(nil)
}
