package lib

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Repo(flagName []string) {
	repoCommand := flag.NewFlagSet("repo", flag.ExitOnError)
	repoUserFlag := repoCommand.String("username", "", "GitHub username")
	repoTypeFlag := repoCommand.String("type", "public", "Type of repository")

	flag.Parse()

	repoCommand.Parse(flagName)

	if repoCommand.Parsed() {
		if *repoUserFlag == "" {
			fmt.Println("Usage of repo:")
			fmt.Println("  -type string")
			fmt.Println("\tType of repository (default \"public\")")
			fmt.Println("  -username string")
			fmt.Println("\tGitHub username")

			os.Exit(2)
		}

		var buffer bytes.Buffer
		var repositories []Repositories

		token, err := Read("token")

		if err == nil && strings.ToLower(*repoTypeFlag) == "private" {
			buffer.WriteString("https://api.github.com/users/" + *repoUserFlag + "/repos?sort=created&access_token=" + token)
		} else {
			buffer.WriteString("https://api.github.com/users/" + *repoUserFlag + "/repos?sort=created")
		}

		repositories = GetRepositories(buffer.String())

		if len(repositories) == 0 {
			fmt.Println("No repositories found")
		} else {
			list := RepoList{repositories}

			list.IterateRepositories()
		}
	}
}
