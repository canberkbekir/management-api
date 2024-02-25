package repository

import (
	"github.com/couchbase/gocb/v2"
	"management-api/internal/model"
	"management-api/internal/util"
)

type IUserRepostiory interface {
	GetAll() ([]model.User, error)
	GetById(id string) (*model.User, error)
	Upsert(user *model.User) error
	Delete(id string) error
}

type UserRepository struct {
	cbClient *gocb.Cluster
}

func (u UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	u.cbClient.Bucket("users").DefaultCollection()

	result, err := u.cbClient.Query("SELECT * FROM users", &gocb.QueryOptions{})
	if err != nil {
		util.Logger.Error().Err(err).Msg("Error querying users")
		return nil, err
	}

	for result.Next() {
		var user model.User
		if err := result.Row(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u UserRepository) GetById(id string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Upsert(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(cbClient *gocb.Cluster) IUserRepostiory {
	return &UserRepository{
		cbClient: cbClient,
	}
}
