package announcerd_test

import (
	"testing"

	"github.com/jdockerty/announcerd/pkg/announcerd"
)

func TestParseAnnouncement(t *testing.T) {

	tests := map[string]struct {
		input    string
		expected string
	}{
		"simple announcement parsing":         {input: "hello\nworld\nannouncement=\"this is my announcement!\"\ntest message", expected: "this is my announcement!"},
		"announcement with a # returns blank": {input: "#announcement=\"testing, one two three.\" another message on the same line", expected: ""},
        "announcement should return blank, even with preceeding text": {input: "description above the announcement\n#announcement=\"this should return blank\"", expected: ""},
        "typo in the announcement prefix will cause a blank return": {input: "an example\nament=\"my announcement\"", expected: ""},
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
