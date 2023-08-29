package repo

import (
	"context"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"search-cli/model"
)

type UserRepoCsv struct {
	filePath string
}

func NewUserRepoCsv(filePath string) *UserRepoCsv {
	return &UserRepoCsv{filePath: filePath}
}

func (u *UserRepoCsv) FindAll(ctx context.Context) ([]*model.User, error) {
	file, err := os.Open(u.filePath)
	if err != nil {
		return nil, fmt.Errorf("%w. %s", ErrReadUserDB, err.Error())
	}
	defer file.Close()

	var users []*model.User
	if err = gocsv.Unmarshal(file, &users); err != nil {
		return nil, fmt.Errorf("%w. %s", ErrReadUserDB, err.Error())
	}

	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	return users, nil
}

func (u *UserRepoCsv) SaveBulk(ctx context.Context, users []*model.User) error {
	file, err := os.Create(u.filePath)
	if err != nil {
		return fmt.Errorf("%w. %s", ErrPersistUsers, err.Error())
	}
	defer file.Close()

	if err = gocsv.Marshal(users, file); err != nil {
		return fmt.Errorf("%w. %s", ErrPersistUsers, err.Error())
	}

	return nil
}
