package golbStore

import (
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// struct for dbquery results
// mgo requires the string literal tags
type Entry struct {
	ObjId   bson.ObjectId `bson:"_id,omitempty"`
	Author  string        `bson:"author,omitempty"`
	Written time.Time     `bson:"written,omitempty"`
	Title   string        `bson:"title"`
	Text    string        `bson:"text"`
	Teaser  string        `bson:"teaser"`
}

// defaults
const (
	blogCollName = "blogEntries"
	blogDbName   = "blog"
)

type MgoBlog struct {
	dbName, collName string
	s                *mgo.Session
}

func (m MgoBlog) getCollection() (*mgo.Collection, *mgo.Session) {
	s := m.s.Clone()

	return s.DB(m.dbName).C(m.collName), s
}

type Options struct {
	DbName         string
	CollectionName string
}

func NewMgoBlog(session *mgo.Session, o *Options) *MgoBlog {
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
