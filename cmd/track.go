/*
Copyright © 2026 notperplexed notperplexed404@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Username string

type GitHubEvent string

var filter GitHubEvent

const (
	Push         GitHubEvent = "PushEvent"
	Watch        GitHubEvent = "WatchEvent"
	Create       GitHubEvent = "CreateEvent"
	Pull         GitHubEvent = "PullRequestEvent"
	IssueComment GitHubEvent = "IssueCommentEvent"
	Issues       GitHubEvent = "IssuesEvent"
	Release      GitHubEvent = "ReleaseEvent"
)

type GithubEventData struct {
	Type      string  `json:"type"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	CreatedAt string  `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Commits []Commit `json:"commits"`
}

type Commit struct {
	Message string `json:"message"`
	Sha     string `json:"sha"`
	Author  Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"` // This is the "Commit Name" / Person
	Email string `json:"email"`
}

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "start the Activity Tracker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Track..")
		url := fmt.Sprintf("https://api.github.com/users/%s/events", Username)
		resp, err := http.Get(url)
		if err != nil {
			color.Red("ERROR: %s", err)
		}
		if resp.StatusCode != http.StatusOK {
			color.Yellow("User not found or API error: %d", resp.StatusCode)
			return
		}

		defer resp.Body.Close()
		var events []GithubEventData
		if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
			color.Red("Error decoding JSON: %v", err)
			return
		}

		for _, event := range events {

			if filter != "all" && event.Type != string(filter) {
				continue // Skip this event
			}
			t, err := time.Parse(time.RFC3339, event.CreatedAt)
			if err != nil {
				color.Red("ERROR: %s", err)
			}
			formattedTime := t.Format("January 04 2026, 12:00 PM")
			switch GitHubEvent(filter) {

			case Watch:
				color.HiCyan("→ Starred | %s at %s", event.Repo.Name, formattedTime)
			case Create:
				color.HiGreen("→ Created | %s at %s", event.Repo.Name, formattedTime)
			case Pull:
				color.HiYellow("→ PR Opened | %s at %s", event.Repo.Name, formattedTime)
			case Issues:
				color.HiRed("→ Issue | %s at %s", event.Repo.Name, formattedTime)
			case Push:
				color.HiGreen("→ Push | %s at %s", event.Repo.Name, formattedTime)
			case IssueComment:
				color.HiBlack("→ IssueComment | %s at %s", event.Repo.Name, formattedTime)
			default:
				// This catches everything else so your tool doesn't stay silent
				fmt.Printf("→ %s | %s at %s\n", event.Type, event.Repo.Name, formattedTime)
			}
		}

	},
}

func (e *GitHubEvent) String() string {
	return string(*e)
}

func (e *GitHubEvent) Set(v string) error {
	switch strings.ToLower(strings.TrimSpace(v)) {

	case "push", "pushevent":
		*e = Push

	case "watch", "star", "watchevent":
		*e = Watch

	case "pull", "pr", "pullrequest", "pullrequestevent":
		*e = Pull

	case "issue", "issues", "issuesevent":
		*e = Issues

	case "comment", "issuecomment", "issuecommentevent":
		*e = IssueComment

	case "create", "createevent":
		*e = Create

	case "release", "releaseevent":
		*e = Release

	default:
		return fmt.Errorf("invalid event type: %s", v)
	}

	return nil
}

func (e *GitHubEvent) Type() string {
	return "GitHubEvent"
}

func init() {
	rootCmd.AddCommand(trackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	trackCmd.Flags().StringVarP(&Username, "user", "u", "notperplexed", "Track Username")
	trackCmd.Flags().VarP(&filter, "type", "t", "Filter by event type")
}
