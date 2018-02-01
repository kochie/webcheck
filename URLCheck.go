package webcheck

import (
	"regexp"
)

var URLregexp *regexp.Regexp = regexp.MustCompile("[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")

func IsValidURL(url string) (bool) {
	return URLregexp.MatchString(url)
}
