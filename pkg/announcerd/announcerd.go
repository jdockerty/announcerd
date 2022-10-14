package announcerd

import (
	"strings"
)

func ParseAnnouncement(s string) string {
	ss := strings.Split(s, "\n")

	var msg string
	announcementPrefix := "announcement="
	for _, v := range ss {

		if strings.HasPrefix(v, announcementPrefix) {
			msgWithQuotes := strings.Split(v, announcementPrefix)[1]

			v := strings.ReplaceAll(msgWithQuotes, "\"", "")
			msg = v
		}
	}

	return msg
}
