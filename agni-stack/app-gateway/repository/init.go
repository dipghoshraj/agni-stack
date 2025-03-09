package repository

import (
	dbmodel "app-gateway/repository/model"
	"context"
	"log"
	"sync"
)

type dbRepo struct{}

type DbRepo interface {
	CreateUser(ctx context.Context, user dbmodel.User) (*dbmodel.User, error)
	GetUser(ctx context.Context, id int64, fileds []string) (*dbmodel.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dbmodel.User, error)
	GetProjectsByUserID(ctx context.Context, user_id int64) ([]*dbmodel.Project, error)
	GetAppsByUserID(ctx context.Context, user_id int64) ([]*dbmodel.App, error)
}

func NewRepository() DbRepo {
	return &dbRepo{}
}

var instance *RepositoryManager
var once sync.Once

type RepositoryManager struct {
	DbRepo DbRepo
}

/*
using singletone pattern for repository manager
to avide the multiple dependency injection of repository manager
to keep the single instance of repository manager

make code more readable and maintainable
*/

func InitRepositoryManager() {
	once.Do(func() {
		instance = &RepositoryManager{
			DbRepo: NewRepository(),
		}
		log.Println("RepositoryManager initialized")
	})
}

func GetRepositoryManager() *RepositoryManager {
	if instance == nil {
		log.Fatal("RepositoryManager is not initialized")
	}

	return instance
}
