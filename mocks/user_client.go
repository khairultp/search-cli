package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"search-cli/client"
)

type UserClient struct {
	mock.Mock
}

func (uc *UserClient) FetchUserData(ctx context.Context) ([]*client.UserDto, error) {
	args := uc.Called(ctx)

	var users []*client.UserDto
	if args.Get(0) != nil {
		users = args.Get(0).([]*client.UserDto)
	}

	return users, args.Error(1)
}
