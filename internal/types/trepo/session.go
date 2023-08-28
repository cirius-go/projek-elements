package trepo

import "github.com/projek-elements/internal/model"

// Session represents the session repository.
type Session interface {
	Common[model.Session]
}
