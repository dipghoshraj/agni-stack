package resolver

import (
	"app-gateway/database"
	"app-gateway/graph/model"
	dbmodel "app-gateway/resolver-service/model"
	"app-gateway/utils"
	"context"
)

func CreateUser(ctx context.Context, input model.UserInput) (*dbmodel.User, error) {
	hashpassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := dbmodel.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashpassword,
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
