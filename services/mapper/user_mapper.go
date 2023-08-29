package mapper

import (
	"search-cli/client"
	"search-cli/model"
)

func ToUsers(userDtos []*client.UserDto) []*model.User {
	var users []*model.User

	for _, item := range userDtos {
		for _, friend := range item.Friends {
			user := &model.User{
				Name:     friend.Name,
				Salary:   item.Balance,
				IsActive: item.IsActive,
				Tags:     item.Tags,
			}
			users = append(users, user)
		}
	}

	return users
}
