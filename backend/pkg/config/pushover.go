package config

import (
	"github.com/apex/log"
	"github.com/gregdel/pushover"
)

type PushoverConfig struct {
	app     *pushover.Pushover
	rec     *pushover.Recipient
	Server  string
	Enable  bool
	UserKey string
	AppKey  string
}

func (p *PushoverConfig) Init() {
	p.app = pushover.New(p.AppKey)
	p.rec = pushover.NewRecipient(p.UserKey)
}

func (p *PushoverConfig) sendSync(message, title string) {
	if !p.Enable {
		return
	}
	log.Debugf("[Pushover] sending title: %s, message: %s", title, message)
	if _, err := p.app.SendMessage(pushover.NewMessageWithTitle(message, p.Server+": "+title), p.rec); err != nil {
		log.WithError(err).Warnf("cannot send pushover message")
	}
}

func (p *PushoverConfig) Send(message, title string) {
	go p.sendSync(message, title)
}
