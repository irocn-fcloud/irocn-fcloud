package bolt

import (
	"github.com/asdine/storm"
	"irocn.com/irocn-fcloud/errors"
)

func get(db *storm.DB, name string, to interface{}) error {
	err := db.Get("config", name, to)
	if err == storm.ErrNotFound {
		return errors.ErrNotExist
	}

	return err
}

func save(db *storm.DB, name string, from interface{}) error {
	return db.Set("config", name, from)
}
