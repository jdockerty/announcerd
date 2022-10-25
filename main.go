package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-github/v47/github"
	"github.com/jdockerty/announcerd/pkg/announcerd"
	"github.com/jdockerty/announcerd/pkg/config"
)

var (
	healthy = []byte(`{
        "status": "ok"
    }`)
)

func PullRequestEventHandler(w http.ResponseWriter, req *http.Request, conf *config.Config) error {

	var event github.PullRequestEvent
	var announcer announcerd.Announcerd

	payload, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(payload, &event); err != nil {
		panic(err)
	}

	if action := event.GetAction(); action == "opened" || action == "reopened" {

		msg := announcerd.ParseAnnouncement(*event.PullRequest.Body)

		if announcerd.IsValidAnnouncement(msg) {
			err := announcer.AnnounceViaWebhook(conf, msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {

	c := &config.Config{}
	err := c.PopulateFromEnv()
	if err != nil {
		fmt.Printf("could not create configuration: %s\n", err)
		return
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(healthy)
	})

	http.HandleFunc("/pulls", func(w http.ResponseWriter, req *http.Request) {
		err := PullRequestEventHandler(w, req, c)
		if err != nil {
			c.Logger.Info().Msgf("unable to send announcement '%s'", err)
		}
	})

	listenAddr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	c.Logger.Info().Msgf("Listening on %s", listenAddr)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		fmt.Printf("could not start server on %s: %s\n", listenAddr, err)
		return
	}
}
