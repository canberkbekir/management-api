package repository

import (
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"management-api/internal/model"
	"management-api/internal/util"
)

type IRoleRepository interface {
	GetAll() ([]model.Role, error)
	GetById(id string) (*model.Role, error)
	Upsert(role *model.Role) error
	Delete(id string) error
}

type RoleRepository struct {
	cbClient *gocb.Cluster
}

func (r RoleRepository) GetAll() ([]model.Role, error) {
	roles := make([]model.Role, 0)
	result, err := r.cbClient.Query("SELECT * FROM roles", &gocb.QueryOptions{})
	if err != nil {
		util.Logger.Error().Err(err).Msg("Error querying users")
		return nil, err
	}

	for result.Next() {
		var role model.Role
		if err := result.Row(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r RoleRepository) GetById(id string) (*model.Role, error) {
	getResult, err := r.cbClient.Bucket("roles").DefaultCollection().Get(id, &gocb.GetOptions{})

	if err != nil {
		return nil, err
	}

	if getResult == nil {
		return &model.Role{}, nil
	}

	var role model.Role
	if err := getResult.Content(&role); err != nil {
		return nil, err
	}
	return &role, nil
}

func (r RoleRepository) Upsert(role *model.Role) error {
	id := uuid.New().String()
	role.ID = id
	_, err := r.cbClient.Bucket("roles").DefaultCollection().Upsert("role_"+id, role, &gocb.UpsertOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (r RoleRepository) Delete(id string) error {
	_, err := r.cbClient.Bucket("roles").DefaultCollection().Remove("role_"+id, &gocb.RemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

func NewRoleRepository(cbClient *gocb.Cluster) IRoleRepository {
	return &RoleRepository{cbClient: cbClient}
}
