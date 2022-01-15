package config

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
