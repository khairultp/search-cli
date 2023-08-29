package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"search-cli/client"
	"search-cli/config"
	"search-cli/model"
	"search-cli/repo"
	"search-cli/services"
	"strings"
)

var ctx context.Context
var userService services.UserService
var argument *Argument

func init() {
	ctx = context.Background()

	apiUrlsStr := config.AppConfig.ApiUrls
	apiUrls := strings.Split(apiUrlsStr, ",")

	userClient := client.NewUserClient(apiUrls)
	userRepo := repo.NewUserRepoCsv(config.AppConfig.CsvFile)

	userService = services.NewUserService(userClient, userRepo)
}

func Run() {

	var err error

	argument, err = parseArgument()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info(fmt.Sprintf("application started with argument %s", argument))

	if argument.fetch {
		if err = fetchUser(ctx); err != nil {
			slog.Error(err.Error())
			return
		}
	}

	result, err := findByTags()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	printData(result)
}

func printData(result []*model.User) {
	slog.Info(fmt.Sprintf("found %v", len(result)))
	slog.Info(fmt.Sprintf("%s", result))

	for i := 0; i < len(result); i++ {
		fmt.Println(fmt.Sprintf("%d. %s \t %s", i+1, result[i].Name, result[i].Salary))
	}
}
