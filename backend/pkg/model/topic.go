package model

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"strings"
)

type TopicArray []*Topic

type Topic struct {
	Description string        `json:"description,omitempty"`
	Submissions SubmissionMap `json:"submissions,omitempty"`
	Played      bool          `json:"played"`
}

func NewTopic(description string) *Topic {
	return &Topic{
		Description: description,
		Submissions: make(SubmissionMap),
	}
}

func (T TopicArray) Exists(topic string) bool {
	for _, t := range T {
		if strings.EqualFold(t.Description, topic) {
			return true
		}
	}
	return false
}

func (T *TopicArray) Add(topic string) {
	*T = append(*T, NewTopic(topic))
}

func (T *TopicArray) Delete(topic string) {
	c := *T
	*T = nil
	for _, t := range c {
		if !strings.EqualFold(t.Description, topic) {
			*T = append(*T, t)
		}
	}
}

func (T TopicArray) NextTopic() (*Topic, error) {
	for _, t := range T {
		if !t.Played {
			return t, nil
		}
	}
	return nil, gerrors.ErrNoTopicsLeft
}

func (T TopicArray) PlayedTopicsCount() (count int) {
	for _, t := range T {
		if t.Played {
			count += 1
		}
	}
	return
}
