package golbStoreMgo

import (
	"github.com/cryptix/golbStore"
	"labix.org/v2/mgo"
)

// defaults
const (
	blogCollName = "blogEntries"
	blogDbName   = "blog"
)

// MgoBlog is the the Mgo version of a golbStore
type MgoBlog struct {
	dbName, collName string
	s                *mgo.Session
}

func (m MgoBlog) getCollection() (*mgo.Collection, *mgo.Session) {
	s := m.s.Clone()

	return s.DB(m.dbName).C(m.collName), s
}

// Options pass the Db- and CollectionName to the Mgo driver
type Options struct {
	DbName         string
	CollectionName string
}

// NewStore creates a new mgo based golbStore
func NewStore(session *mgo.Session, o *Options) golbStore.GolbStorer {
	b := MgoBlog{s: session}

	if o == nil {
		b.collName = blogCollName
		b.dbName = blogDbName
	} else {
		if o.CollectionName != "" {
			b.collName = o.CollectionName
		} else {
			b.collName = blogCollName
		}

		if o.DbName != "" {
			b.dbName = o.DbName
		} else {
			b.dbName = blogDbName
		}
	}

	return &b
}
