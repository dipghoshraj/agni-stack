package services

import (
	"app-gateway/graph/model"
	repository "app-gateway/repository"
	dbmodel "app-gateway/repository/model"
	"app-gateway/utils"
	"context"
	"fmt"
	"strconv"
)

func CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	hashpassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := dbmodel.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashpassword,
	}

	userData, err := repository.GetRepositoryManager().UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// userData :=

	return &model.User{
		ID:    strconv.FormatInt(userData.ID, 10),
		Name:  userData.Name,
		Email: userData.Email,
	}, nil
}

func LoginUser(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {

	user, err := repository.GetRepositoryManager().UserRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	valid_password := utils.CheckPassword(user.Password, input.Password)
	if !valid_password {
		return nil, fmt.Errorf("Invalid password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token: token,
		ID:    strconv.FormatInt(user.ID, 10),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func GetUser(ctx context.Context, id int64) (*model.User, error) {
	user, err := repository.GetRepositoryManager().UserRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    strconv.FormatInt(user.ID, 10),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
