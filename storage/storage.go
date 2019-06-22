package storage

import (
	"irocn.com/irocn-fcloud/auth"
	"irocn.com/irocn-fcloud/settings"
	"irocn.com/irocn-fcloud/share"
	"irocn.com/irocn-fcloud/users"
)

// Storage is a storage powered by a Backend whih makes the neccessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    *users.Storage
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
