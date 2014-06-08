package golbStoreMgo

import (
	"errors"
)

// ErrBadObjectID is returned when the ID is not a valid Mgo ObjectId
var ErrBadObjectID = errors.New("error while validating entry ID")
