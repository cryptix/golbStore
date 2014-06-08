package golbStore

// The GolbStorer interface represents an abstracted store
type GolbStorer interface {
	Latest(n int, withText bool) ([]*Entry, error)
	Get(id string) (*Entry, error)
	Save(e *Entry) error
	Delete(id string) error
}
