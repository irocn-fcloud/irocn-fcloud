package auth

import (
	"net/http"

	"irocn.com/irocn-fcloud/users"
)

// Auther is the authentication interface.
type Auther interface {
	// Auth is called to authenticate a request.
	Auth(r *http.Request, s *users.Storage, root string) (*users.User, error)
	// LoginPage indicates if this auther needs a login page.
	LoginPage() bool
}
