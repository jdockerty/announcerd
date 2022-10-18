package announcerd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jdockerty/announcerd/pkg/config"
)

// Announcer is a common interface for announcements, i.e. sending a message to a destination, such as Slack.
type Announcer interface {
	AnnounceViaWebhook(dest, msg string) error
}

type Announcerd struct{}

func (a *Announcerd) AnnounceViaWebhook(c *config.Config, msg string) error {

	m := make(map[string]string, 1)

	m["text"] = msg
	payload, err := json.Marshal(&m)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.SlackWebhook, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("unable to send POST to %s: %s\n", c.SlackWebhook, err)
	}

	if resp.StatusCode == 200 {
		c.Logger.Info().Msgf("Announced: '%s'", msg)
	} else {
		respMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		c.Logger.Error().Msgf("error when sending to Slack: %s", string(respMsg))
	}

	return nil
}

// IsValidAnnouncement is our simplistic criteria for an announcement being valid, i.e. not blank.
func IsValidAnnouncement(s string) bool {
	return s != ""
}

// ParseAnnouncement will perform rudimentry parsing of a body of text, this assumes that the announcement message
// itself resides on its own line and is not proceeded by anything else on the same line.
// TODO: This is very much a "it works for my use case" attempt, there must be an elegant way to do this in the broader sense.
func ParseAnnouncement(s string) string {
	ss := strings.Split(s, "\n")

	var msg string
	announcementPrefix := "announcement="
	for _, v := range ss {

		if strings.HasPrefix(v, announcementPrefix) {
			msgWithQuotes := strings.Split(v, announcementPrefix)[1]

			v := strings.ReplaceAll(msgWithQuotes, "\"", "")
			msg = v
			break
		}
	}

	return msg
}
