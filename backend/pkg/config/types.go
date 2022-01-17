package config

var (
	Version   = "1.3"
	GitBranch = "unknown"
	GitCommit = "ffffff"
)

type Initer interface {
	Init()
}

type Config struct {
	Pushover *PushoverConfig
}

var DefaultConfigs = Config{
	&PushoverConfig{
		UserKey: "",
		AppKey:  "",
	},
}

var Obj = &DefaultConfigs
