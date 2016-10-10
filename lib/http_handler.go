package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetRepositories(url string) []Repositories {
	response, err := http.Get(url)

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	repoDetails := new([]Repositories)
	err = json.NewDecoder(response.Body).Decode(repoDetails)

	if err != nil {
		fmt.Println("Invalid access token. Please authenticate again.")

		os.Exit(2)
	}

	return *repoDetails
}

func GetEvents(url string) []Events {
	response, err := http.Get(url)

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	evtDetails := new([]Events)
	err = json.NewDecoder(response.Body).Decode(evtDetails)

	if err != nil {
		fmt.Println("Invalid access token. Please authenticate again.")

		os.Exit(2)
	}

	return *evtDetails
}
