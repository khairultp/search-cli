package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"search-cli/model"
)

type UserRepo struct {
	mock.Mock
}

func (uc *UserRepo) SaveBulk(ctx context.Context, users []*model.User) error {
	args := uc.Called(ctx, users)
	return args.Error(0)
}

func (uc *UserRepo) FindAll(ctx context.Context) ([]*model.User, error) {
	args := uc.Called(ctx)

	var users []*model.User
	if args.Get(0) != nil {
		users = args.Get(0).([]*model.User)
	}

	return users, args.Error(1)
}
