package announcerd_test

import (
	"testing"

	"github.com/jdockerty/announcerd/pkg/announcerd"
	"github.com/stretchr/testify/assert"
)

type MockAnnouncer struct{}

func (a *MockAnnouncer) AnnounceViaWebhook(string, string) error { return nil }

func TestParseAnnouncement(t *testing.T) {

	tests := map[string]struct {
		input    string
		expected string
	}{
		"simple announcement parsing":                                 {input: "hello\nworld\nannouncement=\"this is my announcement!\"\ntest message", expected: "this is my announcement!"},
		"announcement with a # returns blank":                         {input: "#announcement=\"testing, one two three.\" another message on the same line", expected: ""},
		"announcement should return blank, even with preceeding text": {input: "description above the announcement\n#announcement=\"this should return blank\"", expected: ""},
		"typo in the announcement prefix will cause a blank return":   {input: "an example\nament=\"my announcement\"", expected: ""},
		"first announcement in a message is used":                     {input: "announcement=\"first announcement!\"\nannouncement=\"second announcement\"", expected: "first announcement!"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := announcerd.ParseAnnouncement(tc.input)
			if got != tc.expected {
				t.Fatalf("Got: %s, expected: %s", got, tc.expected)
			}
		})
	}
}

func TestIsValidAnnouncement(t *testing.T) {

	tests := map[string]struct {
		input    string
		expected bool
	}{
		"should be valid": {input: "an announcement", expected: true},
		"is invalid":      {input: "", expected: false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := announcerd.IsValidAnnouncement(tc.input)
			if got != tc.expected {
				t.Fatalf("Got: %v, expected: %v", got, tc.expected)
			}
		})
	}
}

func TestAnnounceToSlack(t *testing.T) {

	announcementMsg := "this is my announcement!"

	announcer := &MockAnnouncer{}

	err := announcer.AnnounceViaWebhook("https://slack.com", announcementMsg)
	assert.Nil(t, err)
}
