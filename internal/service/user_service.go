package service

import (
	"management-api/internal/model"
	"management-api/internal/repository"
)

type IUserService interface {
	GetAllUser() ([]model.User, error)
	GetUserById(id string) (*model.User, error)
	UpsertUser(user *model.User) error
	DeleteUser(id string) error
}

type UserService struct {
	repository repository.IUserRepository
}

func (u UserService) GetAllUser() ([]model.User, error) {
	return u.repository.GetAll()
}

func (u UserService) GetUserById(id string) (*model.User, error) {
	return u.GetUserById(id)
}

func (u UserService) UpsertUser(user *model.User) error {
	return u.repository.Upsert(user)
}

func (u UserService) DeleteUser(id string) error {
	return u.DeleteUser(id)
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{repository: repository}
}
