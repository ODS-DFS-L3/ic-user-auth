package gocloak

import (
	"log"

	"github.com/Nerzal/gocloak/v13"
)

func NewClient(baseURL, realm, clientID, clientSecret string) (*gocloak.GoCloak, error) {
	client := gocloak.NewClient(baseURL)
	log.Println("client", client)
	return client, nil
}
