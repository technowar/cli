package lib

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func Evt(orgName []string) {
	eventCommand := flag.NewFlagSet("event", flag.ExitOnError)
	eventUserFlag := eventCommand.String("username", "", "GitHub username")

	flag.Parse()

	eventCommand.Parse(orgName)

	if eventCommand.Parsed() {
		if *eventUserFlag == "" {
			fmt.Println("Usage of event:")
			fmt.Println("  -username string")
			fmt.Println("\tGitHub username")

			os.Exit(2)
		}

		var buffer bytes.Buffer
		var events []Events

		token, err := Read("token")

		if err == nil {
			buffer.WriteString("https://api.github.com/users/" + *eventUserFlag + "/events?access_token=" + token)
		} else {
			buffer.WriteString("https://api.github.com/users/" + *eventUserFlag + "/events")
		}

		events = GetEvents(buffer.String())

		if len(events) == 0 {
			fmt.Println("No events found")
		} else {
			list := EventList{events}

			list.IterateEvents(*eventUserFlag)
		}
	}
}
