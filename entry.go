package golbStore

import "time"

// Entry represents an entry in the blog - for all Storeage backends
type Entry struct {
	ID      string    `bson:"_"`
	Author  string    `bson:"author,omitempty"`
	Written time.Time `bson:"written,omitempty"`
	Title   string    `bson:"title"`
	Text    string    `bson:"text"`
	Teaser  string    `bson:"teaser"`
}
