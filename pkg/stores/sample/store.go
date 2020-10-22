package sample

// Store encapsulates operations to a data store (e.g. postgres, redis, ...).
type Store interface {
	Retrieve() (int, error)
}