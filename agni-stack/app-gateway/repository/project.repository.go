package repository

import (
	graphmodel "app-gateway/graph/model"
	"app-gateway/repository/database"
	dbmodel "app-gateway/repository/model"
	"context"
	"fmt"
)

type ProjectRepo interface {
}

func (pr *projRepo) CreateProject(ctx context.Context, input graphmodel.ProjectInput) (graphmodel.Project, error) {

	var graphproj graphmodel.Project

	userID, ok := ctx.Value("user_id").(float64)
	if !ok || userID == 0 {
		return graphproj, fmt.Errorf("missing user ID")
	}

	project := &dbmodel.Project{
		Name:        input.Name,
		OwnerID:     int64(userID),
		Description: *input.Description,
	}

	err := database.DB.Create(&project).Error

	if err = database.DB.Model(&project).Scan(&graphproj).Error; err != nil {
		return graphproj, err
	}

	return graphproj, nil
}
