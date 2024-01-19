package inmemory

import (
	"errors"
	"sync"
	"time"
)

type Transaction struct {
	tx     sync.Mutex
	Status bool
}

func (transaction *Transaction) StartTransaction() error {

	transaction.tx.Lock()
	transaction.Status = true
	time.Sleep(2 * time.Second)
	return nil

}
func (transaction *Transaction) CommiTransaction() error {
	defer transaction.tx.Unlock()
	return nil
}
func (transaction *Transaction) RollbackTransaction() error {
	return errors.New("Rolling back")
}
