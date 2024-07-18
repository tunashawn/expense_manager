package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type AuthConfig struct {
	JWTSecretKey       string `envconfig:"JWT_SECRET_KEY" required:"true"`
	JWTTokenExpireTime int    `envconfig:"JWT_TOKEN_EXPIRE_TIME" required:"true"`
}

type MongoDBConfig struct {
	ConnectionURI string `envconfig:"MONGODB_CONNECTION_URI" required:"true"`
	DatabaseName  string `envconfig:"MONGODB_DATABASE_NAME" required:"true"`
}

func GetConfig(cfg any) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return errors.Wrap(err, "could not read auth_service config")
	}

	return nil
}
