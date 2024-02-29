package repository

import (
	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"management-api/internal/model"
	"management-api/internal/util"
	"time"
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
	key := "user_" + id
	getResult, err := u.cbClient.Bucket("users").DefaultCollection().Get(key, &gocb.GetOptions{})

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
	// Generate a new ID if it doesn't exist
	if user.ID == "" {
		user.ID = uuid.New().String()
		user.Status = 0
		user.CreatedAt = time.Now()
	}

	// Update the timestamp
	user.UpdatedAt = time.Now()

	// Hash the password
	if user.Password != "" {
		password, err := util.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = password
	}

	_, err := u.cbClient.Bucket("users").DefaultCollection().Upsert("user_"+user.ID, user, &gocb.UpsertOptions{})
	return err
}

func (u UserRepository) Delete(id string) error {
	remove, err := u.cbClient.Bucket("users").DefaultCollection().Remove("user_"+id, &gocb.RemoveOptions{})
	if err != nil {
		return err
	}
	util.Logger.Debug().Msgf("Removed user: %s", remove.Cas())
	return nil
}

func NewUserRepository(cbClient *gocb.Cluster) IUserRepository {
	return &UserRepository{
		cbClient: cbClient,
	}
}
