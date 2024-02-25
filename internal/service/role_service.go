package service

import (
	"management-api/internal/model"
	"management-api/internal/repository"
)

type IRoleService interface {
	GetAllRole() ([]model.Role, error)
	GetRoleById(id string) (*model.Role, error)
	UpsertRole(role *model.Role) error
	DeleteRole(id string) error
}

type RoleService struct {
	repository repository.IRoleRepository
}

func (r RoleService) GetAllRole() ([]model.Role, error) {
	return r.repository.GetAll()
}

func (r RoleService) GetRoleById(id string) (*model.Role, error) {
	return r.repository.GetById(id)
}

func (r RoleService) UpsertRole(role *model.Role) error {
	return r.repository.Upsert(role)
}

func (r RoleService) DeleteRole(id string) error {
	return r.repository.Delete(id)
}

func NewRoleService(repository repository.IRoleRepository) IRoleService {
	return &RoleService{repository: repository}
}
