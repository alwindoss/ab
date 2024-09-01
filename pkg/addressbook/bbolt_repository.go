package addressbook

import "go.etcd.io/bbolt"

type Repository interface {
}

func NewBBoltRepository(db *bbolt.DB) Repository {
	return &bboltRepository{
		db: db,
	}
}

type bboltRepository struct {
	db *bbolt.DB
}
