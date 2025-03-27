package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/Nerzal/gocloak/v13"
)

const (
	seedPath                  = "cmd/add_local_user/data/seed.csv"
	directAccessGrantsEnabled = true
)

type Operator struct {
	username     string
	password     string
	operatorID   string
	clientID     string
	clientSecret string
}

func main() {
	addLocal()
}

func addLocal() {
	ctx := context.Background()

	csvPath := seedPath
	client := gocloak.NewClient("http://localhost:4000")
	token, err := client.LoginAdmin(ctx, "admin", "password", "master")
	if err != nil {
		log.Fatalf("error logging in as admin: %v\n", err)
	} else {
		log.Println("Login Successed")
	}
	addOperatorFromCSV(ctx, client, token, csvPath)
}

func addOperatorFromCSV(ctx context.Context, client *gocloak.GoCloak, token *gocloak.JWT, csvPath string) {

	// read csv file
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV records: %v", err)
	}

	// create user by each record
	for _, record := range records {
		username := record[0]
		password := record[1]
		operatorID := record[2]
		clientID := record[3]
		clientSecret := record[4]
		operator := Operator{username, password, operatorID, clientID, clientSecret}
		addOperator(ctx, client, token, operator)
	}
}

func addOperator(ctx context.Context, client *gocloak.GoCloak, token *gocloak.JWT, operator Operator) {
	user := gocloak.User{
		Username: gocloak.StringP(operator.username),
		Email:    gocloak.StringP(operator.operatorID),
		Enabled:  gocloak.BoolP(true),
	}
	// create user
	userID, err := client.CreateUser(ctx, token.AccessToken, "master", user)
	if err != nil {
		log.Fatalf("Error creating user for User %s: %v", operator.operatorID, err)
	}
	fmt.Println("userID", userID)
	_, err = client.CreateClient(ctx, token.AccessToken, "master", gocloak.Client{
		Name:                      gocloak.StringP(operator.username),
		ClientID:                  gocloak.StringP(operator.clientID),
		Secret:                    gocloak.StringP(operator.clientSecret),
		DirectAccessGrantsEnabled: gocloak.BoolP(directAccessGrantsEnabled),
	})
	if err != nil {
		log.Fatalf("Error create Client %s: %v", userID, err)
	}
	err = client.SetPassword(ctx, token.AccessToken, userID, "master", operator.password, false)
	if err != nil {
		log.Fatalf("Error set Password %s: %v", userID, err)
	}
}
