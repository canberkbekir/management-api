package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	*ApplicationConfig
}

type ApplicationConfig struct {
	Server    ServerConfig    `yaml:"server"`
	Couchbase CouchbaseConfig `yaml:"couchbase"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type CouchbaseConfig struct {
	Host     string `yaml:"adresses"`
	Bucket   string `yaml:"bucket"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewConfig() *Config {
	applicationConfig := &ApplicationConfig{}
	applicationConfig.readApplicationConfig()
	return &Config{
		ApplicationConfig: applicationConfig,
	}
}

func (c *ApplicationConfig) readApplicationConfig() {
	// Read the application configuration from a file
	env, found := os.LookupEnv("ACTIVE_PROFILE")
	if !found {
		env = "local"
	}

	v := viper.New()
	v.SetTypeByDefaultValue(true)
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath("./internal/config/")

	err := v.ReadInConfig()
	if err != nil {
		log.Panic().Err(err).Msg("Error reading configuration")
	}

	sub := v.Sub(env)
	err = sub.Unmarshal(c)
	if err != nil {
		log.Panic().Err(err).Msg("Error unmarshalling configuration")
	}

}
