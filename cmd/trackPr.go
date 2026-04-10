/*
Copyright © 2026 notperplexed notperplexed404@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var state string
var mergedOnly bool
var openOnly bool

type PullRequest struct {
	Title    string  `json:"title"`
	State    string  `json:"state"`
	MergedAt *string `json:"merged_at"`
	User     struct {
		Login string `json:"login"`
	} `json:"user"`
	ID int `json:"id"`
}

func trim(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}

var reposotory string

// trackPrCmd represents the trackPr command
var trackPrCmd = &cobra.Command{
	Use:   "trackPr",
	Short: "Track PRS of reposotories",
	Long:  `Simple enter a username and a repo to track all prs made in said reposotory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Tracker..")
		url := fmt.Sprintf(
			"https://api.github.com/repos/%s/%s/pulls?state=all&per_page=100",
			Username, reposotory,
		)

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		var prs []PullRequest
		json.NewDecoder(resp.Body).Decode(&prs)

		seen := make(map[int]bool)

		filtered := []PullRequest{}

		for _, pr := range prs {
			if !seen[pr.ID] {
				seen[pr.ID] = true
				filtered = append(filtered, pr)
			}
		}

		for _, pr := range filtered {
			isMerged := pr.MergedAt != nil

			if state == "merged" && !isMerged {
				continue
			}
			if state == "open" && isMerged {
				continue
			}

			limit := 10
			if len(filtered) < limit {
				limit = len(filtered)
			}

			status := "❌"
			if pr.MergedAt != nil {
				status = "✅"
			}

			if pr.MergedAt != nil {
				color.HiGreen("→ %s %-50s %s", status, trim(pr.Title, 50), pr.User.Login)
			} else {
				color.HiBlack("→ %s %-50s %s", status, trim(pr.Title, 50), pr.User.Login)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(trackPrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trackPrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trackPrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	trackPrCmd.Flags().StringVarP(&Username, "username", "u", "octocat", "Name of the repo owner to track")
	trackPrCmd.Flags().StringVar(&state, "state", "all", "PR state: merged, open, all")
	trackPrCmd.Flags().StringVarP(&reposotory, "reposotory", "r", "Hello-World", "Name of the reposotory to track")

}
