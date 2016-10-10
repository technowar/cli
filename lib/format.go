package lib

import (
	"fmt"
	"time"
)

type Parameter interface{}

type Repositories struct {
	Name              string
	Description       string
	Html_url          string
	Forks_count       int
	Open_issues_count int
	Stargazers_count  int
	Watchers_count    int
}

type Pull_Request struct {
	Title string
	Url   string
}

type Issue struct {
	Title string
}

type Payload struct {
	Action       string
	Ref          string
	Head         string
	Issue        Issue
	Pull_request Pull_Request
}

type Events struct {
	Type       string
	Created_at string
	Public     bool
	Repo       Repositories
	Payload    Payload
}

type RepoList struct {
	Repository []Repositories
}

type EventList struct {
	Event []Events
}

func (list *RepoList) IterateRepositories() {
	limit := 10

	if len(list.Repository) < 10 {
		limit = len(list.Repository)
	}

	for i := limit - 1; i >= 0; i-- {
		if list.Repository[i].Description == "" {
			list.Repository[i].Description = "No Description"
		}

		fmt.Printf("%s\n", list.Repository[i].Name)
		fmt.Printf("  Description: %s\n", list.Repository[i].Description)
		fmt.Printf("  Repository Page: %s\n", list.Repository[i].Html_url)
		fmt.Printf("  Watch: %d\n", list.Repository[i].Watchers_count)
		fmt.Printf("  Star: %d\n", list.Repository[i].Stargazers_count)
		fmt.Printf("  Fork: %d\n", list.Repository[i].Forks_count)
		fmt.Printf("  Issues: %d\n", list.Repository[i].Open_issues_count)
		fmt.Printf("\n")
	}
}

func (list *EventList) IterateEvents(eventUserFlag string) {
	for _, item := range list.Event {
		parse, _ := time.Parse(time.RFC3339, item.Created_at)
		date := parse.Format("January 2, 2006")

		switch item.Type {
		case "CommitCommentEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s CommitCommentEvent\n", eventUserFlag)
			fmt.Printf("\n")
		case "CreateEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s created a repository\n", eventUserFlag)
			fmt.Printf("  Repository name: %s\n", item.Repo.Name)
			fmt.Printf("\n")
		case "ForkEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s forked a repository\n", eventUserFlag)
			fmt.Printf("  Repository name: %s\n", item.Repo.Name)
			fmt.Printf("\n")
		case "IssueCommentEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s %s a comment\n", eventUserFlag, item.Payload.Action)
			fmt.Printf("  Repository name: %s\n", item.Repo.Name)
			fmt.Printf("  Title: %s\n", item.Payload.Issue.Title)
			fmt.Printf("\n")
		case "IssuesEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s %s an issue\n", eventUserFlag, item.Payload.Action)
			fmt.Printf("  Repository name: %s\n", item.Repo.Name)
			fmt.Printf("  Title: %s\n", item.Payload.Issue.Title)
			fmt.Printf("\n")
		case "PullRequestEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s %s a pull-request\n", eventUserFlag, item.Payload.Action)
			fmt.Printf("  Repository name: %s\n", item.Repo.Name)
			fmt.Printf("  Title: %s\n", item.Payload.Pull_request.Title)
			fmt.Printf("  Link: %s\n", item.Payload.Pull_request.Url)
			fmt.Printf("\n")
		case "PullRequestReviewCommentEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s %s a comment on a pull-request\n", eventUserFlag, item.Payload.Action)
			fmt.Printf("  Repository name: %s\n", item.Repo.Name)
			fmt.Printf("  Title: %s\n", item.Payload.Pull_request.Title)
			fmt.Printf("  Link: %s\n", item.Payload.Pull_request.Url)
			fmt.Printf("\n")
		case "PushEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s pushed a commit\n", eventUserFlag)
			fmt.Printf("  Repository name: %s\n", item.Repo.Name)
			fmt.Printf("  Branch name: %s\n", item.Payload.Ref)
			fmt.Printf("  Hash: %s\n", item.Payload.Head)
			fmt.Printf("\n")
		case "WatchEvent":
			fmt.Printf("[%s]\n", date)
			fmt.Printf("  %s watched a repository\n", eventUserFlag)
			fmt.Printf("  Repository name: %s\n", item.Repo.Name)
			fmt.Printf("\n")
		default:
			return
		}
	}
}
