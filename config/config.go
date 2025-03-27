package config

import (
	"errors"
	"os"
	"strconv"
)

// Config
// Summary: This is structure which defines Config
type Config struct {
	Env    string
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		Sslmode  string
	}
	LogLevel              string
	ZapLogLevel           string
	KeycloakBaseURL       string
	KeycloakRealm         string
	KeycloakClientID      string
	KeycloakClientSecret  string
	KeycloakAdminUserName string
	KeycloakAdminPassword string
	TokenIntrospect       string
	EnableIpRestriction   bool
}

var (
	ErrEnvNotDefined    = errors.New("GO_ENV not defined")
	ErrReadConfigFile   = errors.New("config file read error")
	ErrConfigFileFormat = errors.New("config file formant error")
)

// NewConfig
// Summary: This is function which is used to get the configuration from environment variables
// output: (*Config) pointer of Config struct
// output: (error) error object
func NewConfig() (*Config, error) {
	var cfg Config

	var err error

	current := &cfg

	current.Env = os.Getenv("GO_ENV")
	current.Server.Port = os.Getenv("SERVER_PORT")

	current.Database.Host = os.Getenv("DB_HOST")
	current.Database.Port = os.Getenv("DB_PORT")
	current.Database.User = os.Getenv("DB_USER")
	current.Database.Password = os.Getenv("DB_PASSWORD")
	current.Database.Database = os.Getenv("DB_DATABASE")
	current.Database.Sslmode = os.Getenv("DB_SSLMODE")
	current.LogLevel = os.Getenv("ECHO_LOG_LEVEL")
	current.ZapLogLevel = os.Getenv("ZAP_LOG_LEVEL")
	cfg.KeycloakClientID = os.Getenv("KEYCLOAK_CLIENTID")
	cfg.KeycloakClientSecret = os.Getenv("KEYCLOAK_CLIENTSECRET")
	cfg.KeycloakRealm = os.Getenv("KEYCLOAK_REALM")
	cfg.KeycloakBaseURL = os.Getenv("KEYCLOAK_BASEURL")
	cfg.KeycloakAdminUserName = os.Getenv("KEYCLOAK_ADMIN")
	cfg.KeycloakAdminPassword = os.Getenv("KEYCLOAK_PASSWORD")
	cfg.TokenIntrospect = os.Getenv("TOKEN_INTROSPECT")

	if current.EnableIpRestriction, err = strconv.ParseBool(os.Getenv("ENABLE_IP_RESTRICTION")); err != nil {
		return nil, ErrReadConfigFile
	}

	return current, nil
}
