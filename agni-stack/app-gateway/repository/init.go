package repository

import (
	"log"
	"sync"
)

var instance *RepositoryManager
var once sync.Once

type userRepo struct{}
type projRepo struct{}

type RepositoryManager struct {
	UserRepo    UserRepo
	ProjectRepo ProjectRepo
}

func NewUserRepository() UserRepo {
	return &userRepo{}
}

func NewProjectRepository() ProjectRepo {
	return &projRepo{}
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
