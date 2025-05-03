package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v50/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	token := os.Getenv("TOKEN")
	org := os.Getenv("GITHUB_ORG") // IF NEED other wise for a normal user no need it
	repoName := os.Getenv("REPO_NAME")
	// repoType := os.Getenv("REPO_PRIVATE")

	if token == "" || repoName == "" {
		log.Fatal("Github token and repo name must be set")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repo := &github.Repository{
		Name:        github.String(repoName),
		Private:     github.Bool(false),
		AutoInit:    github.Bool(true),
		Description: github.String("Automated repo creation with go"),
	}

	var newRepo *github.Repository
	var resp *github.Response
	// var err error

	if org != "" {
		newRepo, resp, err = client.Repositories.Create(ctx, org, repo)
	} else {
		newRepo, resp, err = client.Repositories.Create(ctx, "", repo)
	}

	if err != nil {
		log.Fatal("faild to create repo : %v\n Response: %+v", err, resp)
	}
	fmt.Printf("Repository created %s\n", newRepo.GetHTMLURL())
}
