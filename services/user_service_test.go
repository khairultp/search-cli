package services

import (
	"context"
	"reflect"
	"search-cli/client"
	"search-cli/mocks"
	"search-cli/model"
	"search-cli/repo"
	"search-cli/services/mapper"
	"testing"
)

func Test_userServiceImpl_FetchUsers(t *testing.T) {
	ctx := context.TODO()

	userDtos := []*client.UserDto{
		{IsActive: true, Balance: "$2,633.92", Tags: []string{"pariatur", "qui"}, Friends: []*client.FriendDto{{Name: "Koch Valdez"}}},
		{IsActive: true, Balance: "$3,817.74", Tags: []string{"enim", "qui"}, Friends: []*client.FriendDto{{Name: "Velazquez Moody"}}},
	}
	userClient1 := &mocks.UserClient{}
	userClient1.On("FetchUserData", ctx).Return(userDtos, nil)

	userClient2 := &mocks.UserClient{}
	userClient2.On("FetchUserData", ctx).Return(userDtos, client.ErrNotAccessiblePartial)

	userClient3 := &mocks.UserClient{}
	userClient3.On("FetchUserData", ctx).Return(nil, client.ErrNotAccessible)

	users := mapper.ToUsers(userDtos)
	userRepo1 := &mocks.UserRepo{}
	userRepo1.On("SaveBulk", ctx, users).Return(nil)

	userRepo2 := &mocks.UserRepo{}
	userRepo2.On("SaveBulk", ctx, users).Return(repo.ErrPersistUsers)

	type fields struct {
		userClient client.UserClient
		userRepo   repo.UserRepo
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"success", fields{userClient1, userRepo1}, args{ctx}, false},
		{"success with error partial", fields{userClient2, userRepo1}, args{ctx}, false},
		{"error save to csv", fields{userClient1, userRepo2}, args{ctx}, true},
		{"error api not accessible", fields{userClient3, nil}, args{ctx}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userServiceImpl{
				userClient: tt.fields.userClient,
				userRepo:   tt.fields.userRepo,
			}
			if err := u.FetchUsers(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("FetchUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userServiceImpl_FindByTags(t *testing.T) {
	ctx := context.TODO()

	userDtos := []*client.UserDto{
		{IsActive: true, Balance: "$2,633.92", Tags: []string{"pariatur", "qui"}, Friends: []*client.FriendDto{{Name: "Koch Valdez"}}},
		{IsActive: true, Balance: "$3,817.74", Tags: []string{"enim", "qui"}, Friends: []*client.FriendDto{{Name: "Velazquez Moody"}}},
	}

	users := mapper.ToUsers(userDtos)
	userRepo1 := &mocks.UserRepo{}
	userRepo1.On("FindAll", ctx).Return(users, nil)

	userRepo2 := &mocks.UserRepo{}
	userRepo2.On("FindAll", ctx).Return(nil, repo.ErrUserNotFound)

	userRepo3 := &mocks.UserRepo{}
	userRepo3.On("FindAll", ctx).Return(nil, repo.ErrReadUserDB)

	pariaturUser := []*model.User{
		{Name: "Koch Valdez", Salary: "$2,633.92", Tags: []string{"pariatur", "qui"}, IsActive: true},
	}

	type fields struct {
		userClient client.UserClient
		userRepo   repo.UserRepo
	}
	type args struct {
		ctx  context.Context
		tags []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.User
		wantErr bool
	}{
		{"found", fields{nil, userRepo1}, args{ctx, []string{"pariatur"}}, pariaturUser, false},
		{"not found by tags", fields{nil, userRepo1}, args{ctx, []string{"pizza"}}, nil, true},
		{"empty users db", fields{nil, userRepo2}, args{ctx, []string{"pariatur"}}, nil, true},
		{"error read users db", fields{nil, userRepo3}, args{ctx, []string{"enim"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserService(tt.fields.userClient, tt.fields.userRepo)
			got, err := u.FindByTags(tt.args.ctx, tt.args.tags...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}
