package repository

import (
	"app-gateway/repository/database"
	dbmodel "app-gateway/repository/model"
	"context"
	"fmt"
	"slices"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user dbmodel.User) (*dbmodel.User, error)
	GetUser(ctx context.Context, id int64, fileds []string) (*dbmodel.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dbmodel.User, error)
}

func (r *userRepo) CreateUser(ctx context.Context, user dbmodel.User) (*dbmodel.User, error) {

	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64, fields []string) (*dbmodel.User, error) {
	user := dbmodel.User{}
	query := database.DB.Where("users.id = ?", id)

	if slices.Contains(fields, "projects") {
		query = query.Joins("Projects")
	}

	if slices.Contains(fields, "apps") {
		query = query.Preload("Apps")
	}

	err := query.Omit("password").First(&user).Error
	if err != nil {
		fmt.Printf("err : %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*dbmodel.User, error) {
	user := dbmodel.User{}
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
