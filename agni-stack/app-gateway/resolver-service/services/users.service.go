package services

import (
	"app-gateway/graph/model"
	repository "app-gateway/repository"
	dbmodel "app-gateway/repository/model"
	"app-gateway/utils"
	"context"
	"fmt"
	"slices"
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
	fields := GetFields(ctx)

	user_object := &model.User{
		ID:    strconv.FormatInt(user.ID, 10),
		Name:  user.Name,
		Email: user.Email,
	}

	if slices.Contains(fields, "projects") {
		projects, err := getProjects(ctx, user.ID)
		if err != nil {
			return user_object, nil
		}
		user_object.Projects = projects
	}

	if slices.Contains(fields, "apps") {
		apps, err := getAppss(ctx, user.ID)
		if err != nil {
			return nil, err
		}
		user_object.Apps = apps
	}
	return user_object, nil
}

func getProjects(ctx context.Context, user_id int64) ([]*model.BasicProject, error) {
	project, err := repository.GetRepositoryManager().UserRepo.GetProjectsByUserID(ctx, user_id)
	if err != nil {
		return nil, err
	}

	graphqlProjects := make([]*model.BasicProject, len(project))

	for i, dbModelproj := range project {

		graphqlProjects[i] = &model.BasicProject{
			ID:          strconv.FormatInt(dbModelproj.ID, 10),
			Name:        dbModelproj.Name,
			Description: dbModelproj.Description,
		}
	}

	return graphqlProjects, nil
}

func getAppss(ctx context.Context, user_id int64) ([]*model.BasicApp, error) {
	apps, err := repository.GetRepositoryManager().UserRepo.GetAppsByUserID(ctx, user_id)
	if err != nil {
		return nil, err
	}
	graphApps := make([]*model.BasicApp, len(apps))

	for i, dbModelApp := range apps {
		graphApps[i] = &model.BasicApp{
			ID:          strconv.FormatInt(dbModelApp.ID, 10),
			Name:        dbModelApp.Name,
			Description: &dbModelApp.Description,
			Image:       &dbModelApp.Image,
		}
	}

	return graphApps, nil
}
