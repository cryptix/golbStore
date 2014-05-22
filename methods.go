package golbStore

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// LatestEntries loads all Blogentries in the results slice
// (sorted descending by date)
func (b MgoBlog) LatestEntries(withText bool) (entries []*Entry, err error) {
	coll, s := b.getCollection()
	defer s.Close()

	proj := bson.M{"text": 0}
	if withText {
		proj["text"] = 1
	}

	err = coll.Find(nil).Select(proj).Sort("-_written").All(&entries)
	if err != nil {
		return nil, err
	}

	return
}

func (b MgoBlog) FindById(id string) (e *Entry, err error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrBadObjectId
	}

	coll, s := b.getCollection()
	defer s.Close()

	qry := bson.M{"_id": bson.ObjectIdHex(id)}
	err = coll.Find(qry).One(&e)
	switch {
	case err == nil:
		return

	case err == mgo.ErrNotFound:
		return nil, ErrEntryNotFound

	default:
		return nil, err
	}
}

func (b MgoBlog) Delete(id string) error {
	// validate the post id
	if !bson.IsObjectIdHex(id) {
		return ErrBadObjectId
	}
	entryId := bson.ObjectIdHex(id)

	coll, s := b.getCollection()
	defer s.Close()

	// Delete entry
	return coll.Remove(bson.M{"_id": entryId})
}

func (b MgoBlog) Save(e *Entry) error {
	coll, s := b.getCollection()
	defer s.Close()

	// building the update bson manually is necessery because mgo/bson irgnores
	// the "ommitempty" tag and we don't want to update timestamp and username.
	// this requires MongoDB 2.4!
	_, err := coll.UpsertId(e.ObjId, bson.M{
		"$set": bson.M{
			"text":   e.Text,
			"title":  e.Title,
			"teaser": e.Teaser,
		},
	})
	return err
}
