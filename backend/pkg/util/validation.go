package util

import (
	"fmt"
	"github.com/apex/log"
	"net/url"
	"regexp"
	"strings"
)

var allowedHostsRegex = []interface{}{
	regexp.MustCompile(`^media[0-9]*\.giphy\.com$`),
	regexp.MustCompile(`^media[0-9]*\.tenor\.com$`),
}

func IsURLAllowed(urlStr string) (bool, error) {
	u, err := url.Parse(urlStr)
	log.Debugf("checking if %s is allowed; host: %s, path: %s, raw: %s", urlStr, u.Host, u.Path, u.RawPath)
	if err != nil {
		return false, err
	}
	path := strings.TrimSpace(strings.ToLower(u.Path))
	if !strings.HasSuffix(path, ".gif") {
		return false, nil
	}

	host := strings.TrimSpace(strings.ToLower(u.Host))

	for _, allowed := range allowedHostsRegex {
		switch a := allowed.(type) {
		case string:
			if strings.EqualFold(a, host) {
				return true, nil
			}
		case *regexp.Regexp:
			if a.MatchString(host) {
				return true, nil
			}
		}
	}
	return false, nil
}

func URLHash(urlStr string) (string, error) {
	u, err := url.Parse(strings.ToLower(urlStr))
	if err != nil {
		return "", err
	}
	host, path := u.Host, u.Path
	path = strings.Trim(path, "/")
	return fmt.Sprintf("%s:%s", host, path), nil
}
