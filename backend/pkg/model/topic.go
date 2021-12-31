package model

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"strings"
)

type Topic struct {
	Description string                 `json:"description,omitempty"`
	Submissions map[string]*Submission `json:"submissions,omitempty"`
	Played      bool                   `json:"played"`
	CanSubmit   bool                   `json:"canSubmit"`
}

func NewTopic(description string) *Topic {
	return &Topic{
		Description: description,
		Submissions: make(map[string]*Submission),
	}
}

type Topics []*Topic

func (T Topics) Exists(topic string) bool {
	for _, t := range T {
		if strings.EqualFold(t.Description, topic) {
			return true
		}
	}
	return false
}

func (T *Topics) Add(topic string) {
	*T = append(*T, NewTopic(topic))
}

func (T *Topics) Delete(topic string) {
	c := *T
	*T = nil
	for _, t := range c {
		if !strings.EqualFold(t.Description, topic) {
			*T = append(*T, t)
		}
	}
}

func (T Topics) Next() (*Topic, error) {
	for _, t := range T {
		if !t.Played {
			return t, nil
		}
	}
	return nil, gerrors.ErrNoTopicsLeft
}

func (T Topics) PlayedCount() (count int) {
	for _, t := range T {
		if t.Played {
			count += 1
		}
	}
	return
}

func (t *Topic) Waiting(game *Game) []interface{} {
	var waiting = make([]interface{}, 0)
	for _, c := range game.Clients {
		if _, ok := t.Submissions[c.Name]; !ok {
			waiting = append(waiting, c.Name)
		}
	}
	return waiting
}
