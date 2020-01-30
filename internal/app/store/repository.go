package store

import (
	"github.com/semka95/go-rest/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
