package golbStoreIpsum

import (
	"github.com/cryptix/golbStore"
)

func (b IpsumBlog) Latest(n int, withText bool) ([]*golbStore.Entry, error) {
	return nil, nil
}

func (b IpsumBlog) Get(id string) (*golbStore.Entry, error) {
	e, found := b.store[id]
	if !found {
		return nil, golbStore.ErrEntryNotFound
	}

	return e, nil
}

func (b IpsumBlog) Delete(id string) error {
	return nil
}

func (b IpsumBlog) Save(e *golbStore.Entry) error {
	b.store[e.ID] = e

	return nil
}
