package main

import (
	"errors"
	"strings"
)

func ParseURL(rawURL string) (scheme, host string, err error) {
	if rawURL == "" {
		return "", "", errors.New("empty URL")
	}

	colonIdx := strings.Index(rawURL, "://")
	if colonIdx == -1 {
		return "", "", errors.New("missing scheme separator '://'")
	}

	scheme = rawURL[:colonIdx]
	rest := rawURL[colonIdx+3:]

	slashIdx := strings.Index(rest, "/")
	if slashIdx == -1 {
		host = rest
	} else {
		host = rest[:slashIdx]
	}

	if host == "" {
		return "", "", errors.New("empty host")
	}

	return scheme, host, nil
}
