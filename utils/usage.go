package utils

import "fmt"

func Usage() {
	fmt.Println("usage: cli <command> [<args>]")
	fmt.Println("\nCommands available:")
	fmt.Println("  login   \tLogs user in")
	fmt.Println("  repo    \tRetrieve 10 latest repositories created")
	fmt.Println("  event   \tRetrieve user's event")
	fmt.Println("\nArguments available:")
	fmt.Println("  username \tGitHub username")
	fmt.Println("  type     \tGitHub flag")
	fmt.Println("    public  \tPublic flag \t[\"default\"]")
	fmt.Println("    private \tPrivate flag")
}
