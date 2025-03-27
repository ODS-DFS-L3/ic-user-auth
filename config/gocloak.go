package config

import (
	"os"
)

type GocloakConfig struct {
	Realm           string
	AdminUserName   string
	AdminPassword   string
	ClientID        string
	ClientSecret    string
	BaseURL         string
	TokenIntrospect string
}

func NewKeycloakClient(cfg *Config) *GocloakConfig {
	var keycloakRealm string
	keycloakRealm, ok := os.LookupEnv("KEYCLOAK_REALM")
	if !ok {
		keycloakRealm = cfg.KeycloakRealm
	}

	return &GocloakConfig{
		Realm:           keycloakRealm,
		ClientID:        cfg.KeycloakClientID,
		ClientSecret:    cfg.KeycloakClientSecret,
		BaseURL:         cfg.KeycloakBaseURL,
		AdminUserName:   cfg.KeycloakAdminUserName,
		AdminPassword:   cfg.KeycloakAdminPassword,
		TokenIntrospect: cfg.TokenIntrospect,
	}
}
