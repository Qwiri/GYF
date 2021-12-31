package util

import (
	"fmt"
	"github.com/apex/log"
	"net/url"
	"strings"
)

var allowedHosts = []string{
	"giphy.com",
	"media.giphy.com",
	"tenor.com",
}

func IsAllowed(urlStr string) (bool, error) {
	u, err := url.Parse(urlStr)
	log.Infof("checking if %s is allowed; host: %s, path: %s, raw: %s", urlStr, u.Host, u.Path, u.RawPath)
	if err != nil {
		return false, err
	}
	for _, allowed := range allowedHosts {
		if strings.EqualFold(allowed, u.Host) {
			return true, nil
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
