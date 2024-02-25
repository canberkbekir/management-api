package repository

import (
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"management-api/internal/model"
	"management-api/internal/util"
)

type IUserRepository interface {
	GetAll() ([]model.User, error)
	GetById(id string) (*model.User, error)
	Upsert(user *model.User) error
	Delete(id string) error
}

type UserRepository struct {
	cbClient *gocb.Cluster
}

func (u UserRepository) GetAll() ([]model.User, error) {
	users := make([]model.User, 0)
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
	getResult, err := u.cbClient.Bucket("users").DefaultCollection().Get(id, &gocb.GetOptions{})

	if err != nil {
		return nil, err
	}

	if getResult == nil {
		return &model.User{}, nil
	}

	var user model.User
	if err := getResult.Content(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) Upsert(user *model.User) error {
	newUUID := uuid.New().String()
	id := "user_" + newUUID
	user.ID = newUUID
	user.Status = true

	_, err := u.cbClient.Bucket("users").DefaultCollection().Upsert(id, user, &gocb.UpsertOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(cbClient *gocb.Cluster) IUserRepository {
	return &UserRepository{
		cbClient: cbClient,
	}
}
