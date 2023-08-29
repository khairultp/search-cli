package cmd

import (
	"context"
)

func fetchUser(ctx context.Context) error {
	return userService.FetchUsers(ctx)
}
