package cmd

import (
	"log/slog"
	"search-cli/model"
	"strings"
)

func findByTags() ([]*model.User, error) {
	tags := strings.Split(argument.tags, ",")
	result, err := userService.FindByTags(ctx, tags...)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return result, nil
}
