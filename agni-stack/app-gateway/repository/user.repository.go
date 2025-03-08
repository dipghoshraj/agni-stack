package repository

import (
	"app-gateway/repository/database"
	dbmodel "app-gateway/repository/model"
	"context"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user dbmodel.User) (*dbmodel.User, error)
	GetUser(ctx context.Context, id int64) (*dbmodel.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dbmodel.User, error)
	GetProjectsByUserID(ctx context.Context, user_id int64) ([]*dbmodel.Project, error)
	GetAppsByUserID(ctx context.Context, user_id int64) ([]*dbmodel.App, error)
}

func (r *userRepo) CreateUser(ctx context.Context, user dbmodel.User) (*dbmodel.User, error) {

	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*dbmodel.User, error) {
	user := dbmodel.User{}
	err := database.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
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

func (r *userRepo) GetProjectsByUserID(ctx context.Context, user_id int64) ([]*dbmodel.Project, error) {
	var projects []*dbmodel.Project

	err := database.DB.Where("owner_id = ?", user_id).Find(&projects).Error
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *userRepo) GetAppsByUserID(ctx context.Context, user_id int64) ([]*dbmodel.App, error) {
	var app []*dbmodel.App

	err := database.DB.Where("owner_id = ?", user_id).Find(&app).Error
	if err != nil {
		return nil, err
	}

	return app, nil
}
