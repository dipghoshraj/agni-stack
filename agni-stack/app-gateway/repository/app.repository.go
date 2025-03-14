package repository

import (
	graphmodel "app-gateway/graph/model"
	dbmodel "app-gateway/repository/model"
	"context"
)

type AppRepo interface {
	CreateApp(ctx context.Context, input graphmodel.App) (*dbmodel.App, error)
}

func (ar *appRepo) CreateApp(ctx context.Context, input graphmodel.App) (*dbmodel.App, error) {

	return nil, nil
}
