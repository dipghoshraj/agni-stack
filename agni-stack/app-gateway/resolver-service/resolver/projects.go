package resolver

import (
	"app-gateway/database"
	"app-gateway/graph/model"
	dbmodel "app-gateway/repository/model"
	"context"
	"fmt"
)

func CreateProject(ctx context.Context, input model.ProjectInput) (*dbmodel.Project, error) {
	userID, ok := ctx.Value("user_id").(float64)
	if !ok || userID == 0 {
		return nil, fmt.Errorf("missing user ID")
	}

	project := dbmodel.Project{
		Name:        input.Name,
		OwnerID:     int64(userID),
		Description: *input.Description,
	}

	err := database.DB.Create(&project).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}
