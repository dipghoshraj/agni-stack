package repository

import (
	"app-gateway/repository/database"
	dbmodel "app-gateway/repository/model"
	"context"
	"fmt"
	"slices"
)

func (r *dbRepo) CreateUser(ctx context.Context, user dbmodel.User) (*dbmodel.User, error) {

	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *dbRepo) GetUser(ctx context.Context, id int64, fields []string) (*dbmodel.User, error) {
	user := dbmodel.User{}
	query := database.DB.Where("users.id = ?", id)

	if slices.Contains(fields, "projects") {
		query = query.Preload("Projects")
	}

	if slices.Contains(fields, "apps") {
		query = query.Preload("Apps")
	}

	err := query.First(&user).Error
	if err != nil {
		fmt.Printf("err : %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *dbRepo) GetUserByEmail(ctx context.Context, email string) (*dbmodel.User, error) {
	user := dbmodel.User{}
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *dbRepo) GetProjectsByUserID(ctx context.Context, user_id int64) ([]*dbmodel.Project, error) {
	var projects []*dbmodel.Project

	err := database.DB.Where("owner_id = ?", user_id).Find(&projects).Error
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *dbRepo) GetAppsByUserID(ctx context.Context, user_id int64) ([]*dbmodel.App, error) {
	var app []*dbmodel.App

	err := database.DB.Where("owner_id = ?", user_id).Find(&app).Error
	if err != nil {
		return nil, err
	}

	return app, nil
}
