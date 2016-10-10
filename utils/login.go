package utils

import (
	files "../lib"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"io/ioutil"
)

type Config struct {
	ClientID     string `json:"ClientID"`
	ClientSecret string `json:"ClientSecret"`
}

func getConfig() Config {
	raw, err := ioutil.ReadFile("./utils/token.json")

	if err != nil {
		panic(err)
	}

	var config Config

	json.Unmarshal(raw, &config)

	return config
}

func Login() {
	config := getConfig()
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"repo"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)
	fmt.Print("Code: ")

	var code string

	if _, err := fmt.Scan(&code); err != nil {
		panic(err)
	}

	token, err := conf.Exchange(ctx, code)

	if err != nil || !token.Valid() {
		panic("Invalid token")
	}

	files.Write(token.AccessToken, "token")
}
