package config

import (
	"fmt"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v47/github"
)

type AnnouncerdConfiguration interface {
	New(appId, installationId int64, githubAppKeyFile string) error
}

type Config struct {
	Client *github.Client
}

func New(appId, installationId int64, githubAppKeyFile string) (*Config, error) {

	if githubAppKeyFile == "" {
		return nil, fmt.Errorf("GitHub App private key file expected, got empty string")
	}
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appId, installationId, githubAppKeyFile)
	if err != nil {
		return nil, err
	}


    client := github.NewClient(&http.Client{Transport: itr})

    return &Config{
        Client: client,
    }, nil
}
