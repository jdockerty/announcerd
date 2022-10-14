package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-github/v47/github"
	"github.com/jdockerty/announcerd/pkg/announcerd"
)

func PullRequestEventHandler(w http.ResponseWriter, req *http.Request, client *github.Client) {

	var event github.PullRequestEvent

	payload, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(payload, &event); err != nil {
		panic(err)
	}

	if action := event.GetAction(); action == "opened" || action == "reopened" {

        msg := announcerd.ParseAnnouncement(*event.PullRequest.Body)

        fmt.Println("PARSED BODY:", msg)
        // Send to slack
	}

}
