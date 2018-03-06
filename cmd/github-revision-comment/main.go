/*
Command 'github-revision-comment' comments on GitHub pull request.

  $ github-revision-comment ORG REPO REVISION COMMENT

For example, if you want to comment on https://github.com/tcnksm/ghr pull request number 987

  $ github-revision-comment tcnksm ghr 987 "This is test comment"

To use this command, you need to prepare GitHub API Token and set it via GITHUB_TOKEN
env var.

To install, use go get,

  $ go get github.com/tcnksm/misc/cmd/github-revision-comment

*/
package main

import (
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

const EnvToken = "GITHUB_TOKEN"

func main() {
	if len(os.Args) < 5 {
		log.Fatal("[Usage] github-revision-comment ORG REPO REVISION BODY")
	}

	token := os.Getenv(EnvToken)
	if len(token) == 0 {
		log.Fatal("You need GitHub API token via GITHUB_TOKEN env var")
	}

	owner, repo, revision := os.Args[1], os.Args[2], os.Args[3]

	body := strings.Join(os.Args[4:], " ")

	// Construct github HTTP client
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	if _, _, err := client.Repositories.CreateComment(context.Background(), owner, repo, revision, &github.RepositoryComment{
		Body: &body,
	}); err != nil {
		log.Fatalf("[ERROR] Failed to create comment: %s", err)
	}

	log.Printf("[INFO] Successfully created a comment!")
}
