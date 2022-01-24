package config

import "fmt"

var (
	Version   = "1.3.0"
	GitBranch = "unknown"
	GitCommit = "ffffff"
)

func FormattedVersion() string {
	return fmt.Sprintf("%s-%s:%s", Version, GitBranch, GitCommit)
}

type Initer interface {
	Init()
}

type Config struct {
	Pushover *PushoverConfig
}

var DefaultConfigs = Config{
	&PushoverConfig{
		Server:  "unknown",
		Enable:  false,
		UserKey: "",
		AppKey:  "",
	},
}

var Obj = &DefaultConfigs
