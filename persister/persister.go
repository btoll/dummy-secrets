package persister

import (
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type Persister interface {
	Store(secret string) uuid.UUID
	Retrieve(id uuid.UUID) (string, error)
}

type MemPersister struct {
	storage map[string]string
	mux     sync.RWMutex
}

func NewMemPersister() *MemPersister {
	return &MemPersister{
		storage: make(map[string]string),
	}
}

func (mp *MemPersister) Store(secret string) uuid.UUID {
	id := uuid.NewV4()

	mp.mux.Lock()
	defer mp.mux.Unlock()

	mp.storage[id.String()] = secret

	return id
}

func (mp *MemPersister) Retrieve(id uuid.UUID) (string, error) {
	mp.mux.Lock()
	defer mp.mux.Unlock()

	secret, ok := mp.storage[id.String()]
	if !ok {
		return "", fmt.Errorf("error: a secret with id %v doesn't exist.", id)
	}
	delete(mp.storage, id.String())

	return secret, nil
}
