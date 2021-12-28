package model

type Topic struct {
	Description string
	Submissions map[*Client]*Submission
}

func NewTopic(description string) *Topic {
	return &Topic{
		Description: description,
		Submissions: make(map[*Client]*Submission),
	}
}
