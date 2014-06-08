// Package golbStoreIpsum is an in-memory version of the blog that returns Lorem Ipsum-esque text for testing purpose
package golbStoreIpsum

import (
	"github.com/cryptix/golbStore"
)

type MemStore map[string]*golbStore.Entry

// IpsumBlog is the the version of a golbStore for testing
type IpsumBlog struct {
	store MemStore
}

// NewStore creates a new mgo based golbStore
func NewStore() golbStore.GolbStorer {
	var i IpsumBlog
	i.store = make(MemStore, 10)

	return &i
}
