package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/mail"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	token string = os.Getenv("GH_TOKEN")
	owner string = os.Getenv("GH_OWNER")
	repo  string = os.Getenv("GH_REPO")
	pipe  bool   = false
)

func init() {
	flag.StringVar(&token, "token", token, "The OAuth2 Token to authenticate against the GitHub API.")
	flag.StringVar(&owner, "owner", owner, "The owner of the GitHub repository to raise the input mail to as an issue.")
	flag.StringVar(&repo, "repo", repo, "The GitHub repository to raise the input mail to as an issue.")

	flag.BoolVar(&pipe, "pipe", pipe, "Pipe stdin through to stdout for chaining to an extra process.")

	flag.Parse()

	if token == "" || repo == "" {
		log.Fatalln("GitHub OAuth2 token and repo are both required.")
	}
}

func main() {

	var r io.Reader

	if pipe {
		r = bufio.NewReader(io.TeeReader(os.Stdin, os.Stdout))
	} else {
		r = bufio.NewReader(os.Stdin)
	}

	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}

	h := m.Header

	subject := h.Get("Subject")

	gh_body := fmt.Sprintf("Date: %s\nFrom: %s\nTo: %s\nSubject: %s\n\n%s", h.Get("Date"), h.Get("From"), h.Get("To"), subject, body)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	gh := github.NewClient(tc)

	_, _, err = gh.Issues.Create(owner, repo, &github.IssueRequest{
		Title: &subject,
		Body:  &gh_body,
	})
	if err != nil {
		log.Fatal(err)
	}

}

//  vim: set ts=4 sw=4 tw=0 et:
