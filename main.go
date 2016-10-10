package main

import (
	"./lib"
	"./utils"
	"os"
)

func main() {
	if arg := os.Args[1:]; len(arg) == 0 {
		utils.Usage()
		os.Exit(2)

		return
	}

	switch os.Args[1] {
	case "login":
		utils.Login()
	case "repo":
		lib.Repo(os.Args[2:])
	case "event":
		lib.Evt(os.Args[2:])
	default:
		utils.Usage()
		os.Exit(2)
	}
}
