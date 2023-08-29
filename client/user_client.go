package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var ErrNotAccessiblePartial = errors.New("some api are not accessible")
var ErrNotAccessible = errors.New("all api are not accessible")

type UserClient interface {
	FetchUserData(ctx context.Context) ([]*UserDto, error)
}

type userClient struct {
	apiUrls []string
}

func NewUserClient(apiUrls []string) UserClient {
	return &userClient{apiUrls: apiUrls}
}

func (uc *userClient) FetchUserData(ctx context.Context) ([]*UserDto, error) {
	var callErrors []error
	var users []*UserDto
	for _, url := range uc.apiUrls {
		result, err := uc.fetchUserData(url)
		if err != nil {
			callErrors = append(callErrors, err)
			continue
		}

		users = append(users, result...)
	}

	var errStatus error

	if len(callErrors) == len(uc.apiUrls) {
		errStatus = ErrNotAccessible
	} else if len(callErrors) != 0 {
		errStatus = ErrNotAccessiblePartial
	}

	return users, errStatus
}

func (uc *userClient) fetchUserData(url string) ([]*UserDto, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users []*UserDto
	if err = json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil
}
