package golbStore

import (
	"errors"
)

// ErrEntryNotFound is returned when a blog entry could not be found in the store
var ErrEntryNotFound = errors.New("no Blog Entry with that ID")
