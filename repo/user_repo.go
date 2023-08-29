package repo

import (
	"context"
	"errors"
	"search-cli/model"
)

var ErrUserNotFound = errors.New("users is not found")
var ErrPersistUsers = errors.New("can not persist users data")
var ErrReadUserDB = errors.New("can not read users db")

type UserRepo interface {
	SaveBulk(ctx context.Context, users []*model.User) error
	FindAll(ctx context.Context) ([]*model.User, error)
}
