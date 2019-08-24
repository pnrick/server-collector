package collector

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/v28/github"	// with go modules enabled
	"golang.org/x/oauth2"
	"context"
	"os"
	"reflect"
	// "strings"
)


// Notes: For generic and decoupling purpose, you might always want to inject these from the environment vaiables
// or some kinds of configuration yaml
const (
	// MY_ACCESS_TOKEN your github Personal Access Token
    MyAccessToken = "8992518d8cda5290ba387739837588662d6806e4"
    // TARGET_REPO_URL the target repo url you want to monitor
	TargetRepoURL = "https://github.com/lmchih/server-collector"
    // TARGET_BRANCH the target branch
    TargetBranch = "master"
)

var (
	sourceOwner = flag.String("source-owner", "lmchih", "Name of the owner (user or org) of the repo to monitor the latest commits.")
	sourceRepo = flag.String("source-repo", "server-collector", "name of repo to monitor the commits.")
	baseBranch = flag.String("base-branch", "master", "Name of the branch to monitor")

	client *github.Client
	ctx = context.Background()
)


func checkCommitTime() {

	flag.Parse()
	log.Printf("TargetRepoURL: %v\nTargetBranch: %s\n", TargetRepoURL, TargetBranch)

	// TODO: Get the latest commit
	ctx = context.Background()
	
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: MyAccessToken},
	)
	
	tc := oauth2.NewClient(ctx, ts)

	// get go-github client
	client = github.NewClient(tc)
	// fmt.Printf("github Client: %v\n", client)

	// list specific repository's commit info
	// var truncStr = strings.Split(TargetRepoURL, "https://github.com/")[1]
	// var owner = strings.Split(truncStr, "/")[0]
	// var repo = strings.Split(truncStr, "/")[1]
	// commitInfo, _, err := client.Repositories.ListCommits(ctx, owner, repo, nil)
	commitInfo, _, err := client.Repositories.ListCommits(ctx, *sourceOwner, *sourceRepo, nil)
	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
		os.Exit(1)
	}

	var lastCommit = commitInfo[0]
	// fmt.Printf("lastCommit: %v\n", reflect.TypeOf(lastCommit))
	var lastCommitDate = *lastCommit.Commit.Committer.Date
	fmt.Printf("lastDCommitDate: %v\n", reflect.TypeOf(lastCommitDate))
	
	// compare to commit time with now
	now := time.Now().UTC()
	fmt.Printf("Now: %v\n", now)
	since := time.Since(lastCommitDate)
	fmt.Printf("since: %v\n", since)
	// convert since to days
	days := int(since.Hours() / 24)
	fmt.Printf("days: %d\n", days)

	// if older than three days, terminate the server.
	if days >= 3 {
		terminate()
	}
}
