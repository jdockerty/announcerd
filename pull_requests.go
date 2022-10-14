package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-github/v47/github"
)

func PullRequestEventHandler(w http.ResponseWriter, req *http.Request) {

	var event github.PullRequestEvent

	payload, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(payload, &event); err != nil {
		panic(err)
	}

	if action := event.GetAction(); action == "opened" || action == "reopened" {
		fmt.Println("PR OPENED!")

		b, a, f := strings.Cut(*event.PullRequest.Body, "announcement=")
		if f {
			fmt.Printf("BEFORE: %s\n", b)
			fmt.Printf("AFTER: %s\n", a)
		} else {
			fmt.Println("announcement= not found in msg")
		}
	}

}
