package repository

import (
	"log"
	"sync"
)

type userRepo struct{}

func NewUserRepository() UserRepo {
	return &userRepo{}
}

var instance *RepositoryManager
var once sync.Once

type RepositoryManager struct {
	UserRepo UserRepo
}

func InitRepositoryManager() {
	once.Do(func() {
		instance = &RepositoryManager{
			UserRepo: NewUserRepository(),
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
