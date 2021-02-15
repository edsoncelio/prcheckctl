package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

func countPullRequests(repos []*github.Repository, client *github.Client, ctx context.Context, prOpts *github.PullRequestListOptions, owner string, countPrsActual int) int {
	for _, r := range repos {
		prs, _, err := client.PullRequests.List(ctx, owner, *r.Name, prOpts)
		countPrsActual += len(prs)
		if err != nil {
			fmt.Println(err)
		}

	}
	return countPrsActual
}

func getAllPRs(user string, pool int) {

	if os.Getenv("GH_TOKEN") == "" {
		fmt.Println("Github token not found, export your token: export GH_TOKEN=<your_github_token>")
		os.Exit(1)
	}

	var countPrsActual int = 0
	var countPrsNew int = 0
	var startScript bool = true
	var countNew int = 0

	var prs []*github.PullRequest

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	owner := user

	prOpts := &github.PullRequestListOptions{}
	repoOpt := &github.RepositoryListOptions{Type: "owner"}
	for {

		countPrsNew = countPrsActual
		countPrsActual = 0
		repos, _, err := client.Repositories.List(ctx, owner, repoOpt)
		if err != nil {
			fmt.Println(err)
		}

		countPrsActual = countPullRequests(repos, client, ctx, prOpts, owner, countPrsActual)

		if countPrsActual != countPrsNew || startScript {
			fmt.Println("Showing open PRs...")

			if countPrsNew > 0 {
				prs = nil
				countNew = countPrsActual - countPrsNew
				for _, r := range repos {
					prs_temp, _, err := client.PullRequests.List(ctx, owner, *r.Name, prOpts)
					if err != nil {
						fmt.Println(err)
					}
					for _, new := range prs_temp {
						prs = append(prs, new)
					}

				}
				for _, new := range prs[0:countNew] {
					fmt.Printf("Title: %s - Creator: %s - Created at: %s\n", *new.Title, *new.User.Login, *new.CreatedAt)
				}

			} else {
				for _, r := range repos {
					prs, _, err := client.PullRequests.List(ctx, owner, *r.Name, prOpts)
					if err != nil {
						fmt.Println(err)
					}

					for _, pr := range prs {
						fmt.Printf("Title: %s - Creator: %s - Created at: %s\n", *pr.Title, *pr.User.Login, *pr.CreatedAt)

					}

				}
			}

		} else {
			log.Printf("No new PRs!")
		}

		fmt.Println("------------------")
		fmt.Println("Waiting ", pool, "seconds...")
		time.Sleep(time.Duration(pool) * time.Second)

		startScript = false

	}
}
