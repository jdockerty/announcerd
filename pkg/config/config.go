package config

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v47/github"
	"github.com/rs/zerolog"
)

const (
	ghAppIdEnv      = "ANNOUNCERD_GH_APP_ID"
	ghAppKeyFileEnv = "ANNOUNCERD_GH_APP_KEY_FILE"
	slackWebhookEnv = "ANNOUNCERD_SLACK_WEBHOOK"
	hostEnv         = "ANNOUNCERD_HOST"
	portEnv         = "ANNOUNCERD_PORT"
)

type Config struct {
	Client           *github.Client
	SlackWebhook     string
	GitHubAppKeyFile string
	GitHubAppId      int64
	Host             string
	Port             int64
	Logger           *zerolog.Logger
}

// PopulateFromEnv will fill in the configuration using the pre-defined environment variables.
// This is a simple wrapper for the NewFromSource function.
func (c *Config) PopulateFromEnv() error {
	appIdEnv, ok := os.LookupEnv(ghAppIdEnv)
	if !ok {
		return fmt.Errorf("%s is a required environment variable", ghAppIdEnv)
	}

	appId, err := strconv.ParseInt(appIdEnv, 10, 64)
	if err != nil {
		return fmt.Errorf("could not parse %s value to int64: %s\n", ghAppIdEnv, err)
	}

	ghAppKeyFilePath, ok := os.LookupEnv(ghAppKeyFileEnv)
	if !ok {
		return fmt.Errorf("%s is a required environment variable", ghAppKeyFileEnv)
	}

	slackWebhook, ok := os.LookupEnv(slackWebhookEnv)
	if !ok {
		return fmt.Errorf("%s is a required environment variable", slackWebhookEnv)
	}

	ghClient, err := NewGitHubClient(appId, 1, ghAppKeyFilePath)
	if err != nil {
		return err
	}

	c.Host = "localhost"
	host, ok := os.LookupEnv(hostEnv)
	if ok {
		c.Host = host
	}

	c.Port = 6000
	port, ok := os.LookupEnv(portEnv)
	if ok {
		p, err := strconv.ParseInt(port, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to parse %s into int64: %s\n", portEnv, err)
		}
		c.Port = p
	}

	c.GitHubAppId = appId
	c.GitHubAppKeyFile = ghAppKeyFilePath
	c.SlackWebhook = slackWebhook
	c.Client = ghClient

	appLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &appLogger

	c.Logger = &appLogger

	return nil
}

// NewGitHubClient returns a new GitHub client, implying use from a GitHub App.
func NewGitHubClient(appId, installationId int64, githubAppKeyFile string) (*github.Client, error) {

	if githubAppKeyFile == "" {
		return nil, fmt.Errorf("GitHub App private key file expected, got empty string")
	}

	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appId, installationId, githubAppKeyFile)
	if err != nil {
		return nil, err
	}

	client := github.NewClient(&http.Client{Transport: itr})

	return client, nil
}

// New will create a blank configuration.
func New() *Config { return &Config{} }
