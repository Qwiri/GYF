package model

type Submission struct {
	Creator *Client
	URL     string
}

func NewSubmission(creator *Client, url string) *Submission {
	return &Submission{creator, url}
}
