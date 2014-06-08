package golbStoreMgo

import (
	"github.com/cryptix/golbStore"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// Latest loads n Blogentries in the results slice
// (sorted descending by date)
func (b MgoBlog) Latest(n int, withText bool) (entries []*golbStore.Entry, err error) {
	coll, s := b.getCollection()
	defer s.Close()

	// supress text by default to make results shorter
	proj := bson.M{"text": 0}
	if withText {
		proj["text"] = 1
	}

	qry := coll.Find(nil).Select(proj).Sort("-_written")

	// add an Limit to the query if specified
	if n > 0 {
		qry = qry.Limit(n)
	}

	// run the query
	err = qry.All(&entries)
	if err != nil {
		return nil, err
	}

	return
}

// Get loads a single Entry
func (b MgoBlog) Get(id string) (e *golbStore.Entry, err error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrBadObjectID
	}

	coll, s := b.getCollection()
	defer s.Close()

	qry := bson.M{"_id": bson.ObjectIdHex(id)}
	err = coll.Find(qry).One(&e)
	switch {
	case err == nil:
		return

	case err == mgo.ErrNotFound:
		return nil, golbStore.ErrEntryNotFound

	default:
		return nil, err
	}
}

// Delete drops the entry with id from the store
func (b MgoBlog) Delete(id string) error {
	// validate the post id
	if !bson.IsObjectIdHex(id) {
		return ErrBadObjectID
	}
	entryID := bson.ObjectIdHex(id)

	coll, s := b.getCollection()
	defer s.Close()

	// Delete entry
	return coll.Remove(bson.M{"_id": entryID})
}

// Save updates an entry
func (b MgoBlog) Save(e *golbStore.Entry) error {
	if !bson.IsObjectIdHex(e.ID) {
		return ErrBadObjectID
	}
	entryID := bson.ObjectIdHex(e.ID)

	coll, s := b.getCollection()
	defer s.Close()

	// building the update bson manually is necessery because mgo/bson irgnores
	// the "ommitempty" tag and we don't want to update timestamp and username.
	// this requires MongoDB 2.4!

	_, err := coll.UpsertId(entryID, bson.M{
		"$set": bson.M{
			"text":   e.Text,
			"title":  e.Title,
			"teaser": e.Teaser,
		},
	})
	return err
}
