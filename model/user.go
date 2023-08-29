package model

import "fmt"

type User struct {
	Name     string   `csv:"Name"`
	Salary   string   `csv:"Salary"`
	IsActive bool     `csv:"Is Active"`
	Tags     []string `csv:"Tags"`
}

func (u *User) String() string {
	return fmt.Sprintf("%s %s", u.Name, u.Salary)
}
