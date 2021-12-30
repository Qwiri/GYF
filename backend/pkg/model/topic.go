package model

type Topic struct {
	Description string                 `json:"description,omitempty"`
	Submissions map[string]*Submission `json:"submissions,omitempty"`
}

func NewTopic(description string) *Topic {
	return &Topic{
		Description: description,
		Submissions: make(map[string]*Submission),
	}
}
