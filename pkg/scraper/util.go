package scraper

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/indigo-sadland/ffuf/v2/pkg/ffuf"
)

func headerString(headers map[string][]string) string {
	val := ""
	for k, vslice := range headers {
		for _, v := range vslice {
			val += fmt.Sprintf("%s: %s\n", k, v)
		}
	}
	return val
}

func isActive(name string, activegroups []string) bool {
	return ffuf.StrInSlice(strings.ToLower(strings.TrimSpace(name)), activegroups)
}

func parseActiveGroups(activestr string) []string {
	retslice := make([]string, 0)
	for _, v := range strings.Split(activestr, ",") {
		retslice = append(retslice, strings.ToLower(strings.TrimSpace(v)))
	}
	return retslice
}

func GetParent(rawUrl string) string {
	var parent string

	u, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	// Remove the leading and trailing "/" then split on "/"
	trimmedPath := strings.Trim(u.Path, "/")
	segments := strings.Split(trimmedPath, "/")

	// Check if there is at least one segment
	if len(segments) > 1  {
		firstSegment := segments[0]
		parent = firstSegment
	} else {
		parent = "/"
	}

	return parent
}

func GetPort(rawUrl string) string {
	u, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	port := u.Port()
	if port == "" {
		protocol := regexp.MustCompile(`\bhttps?://(.*?)/?`).FindStringSubmatch(rawUrl)[0]
		if protocol == "http" {
			port = "80"
		} else {
			port = "443"
		}
	}
	return port
}
