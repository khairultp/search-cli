package services

import (
	"context"
	"errors"
	"golang.org/x/exp/maps"
	"search-cli/client"
	"search-cli/model"
	"search-cli/repo"
	"search-cli/services/mapper"
	"strings"
)

type UserService interface {
	FetchUsers(ctx context.Context) error
	FindByTags(ctx context.Context, tags ...string) ([]*model.User, error)
}

type userServiceImpl struct {
	userClient client.UserClient
	userRepo   repo.UserRepo
}

func NewUserService(userClient client.UserClient, userRepo repo.UserRepo) UserService {
	return &userServiceImpl{
		userClient: userClient,
		userRepo:   userRepo,
	}
}

func (u *userServiceImpl) FetchUsers(ctx context.Context) error {
	userDtos, err := u.userClient.FetchUserData(ctx)
	if err != nil && errors.Is(err, client.ErrNotAccessible) {
		return err
	}

	users := mapper.ToUsers(userDtos)

	return u.userRepo.SaveBulk(ctx, users)
}

func (u *userServiceImpl) FindByTags(ctx context.Context, tags ...string) ([]*model.User, error) {
	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	filteredUsers := u.filterByTags(users, tags)

	if len(filteredUsers) == 0 {
		return nil, repo.ErrUserNotFound
	}

	return filteredUsers, nil
}

func (u *userServiceImpl) filterByTags(users []*model.User, tags []string) []*model.User {
	tagUsers := make(map[string][]*model.User)

	for _, user := range users {
		for _, tag := range user.Tags {
			if tag = strings.TrimSpace(tag); tag == "" {
				continue
			}

			tagUsers[tag] = append(tagUsers[tag], user)
		}
	}

	filtered := make(map[string]*model.User)

	for _, tag := range tags {
		usersByTag := tagUsers[tag]
		for _, user := range usersByTag {
			if _, ok := filtered[user.Name]; ok {
				continue
			}
			filtered[user.Name] = user
		}
	}

	return maps.Values(filtered)
}
