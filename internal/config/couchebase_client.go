package config

import (
	"github.com/couchbase/gocb/v2"
	"management-api/internal/util"
	"time"
)

func NewCouchebaseClient(config *CouchbaseConfig) (*gocb.Cluster, error) {

	cluster, err := gocb.Connect("couchbase://"+config.Host, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: config.Username,
			Password: config.Password,
		},
	})

	if err != nil {
		util.Logger.Fatal().Err(err).Msg("Error connecting to couchbase")
		return nil, err
	}

	err = cluster.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("Error connecting to Couchbase cluster")
		return nil, err
	}
	util.Logger.Info().Msg("Connected to couchbase")
	return cluster, nil
}
