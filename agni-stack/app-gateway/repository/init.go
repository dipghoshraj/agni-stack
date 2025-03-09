package repository

import (
	"log"
	"sync"
)

type userRepo struct{}
type projRepo struct{}

func NewUserRepository() UserRepo {
	return &userRepo{}
}

func NewProjectRepository() ProjectRepo {
	return &projRepo{}
}

var instance *RepositoryManager
var once sync.Once

type RepositoryManager struct {
	UserRepo    UserRepo
	ProjectRepo ProjectRepo
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
			UserRepo:    NewUserRepository(),
			ProjectRepo: NewProjectRepository(),
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
