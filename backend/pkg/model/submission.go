package model

type Submission struct {
	Creator *Client
	Voters  []*Client
	URL     string
}

//goland:noinspection GoUnusedExportedFunction
func NewSubmission(creator *Client, url string) *Submission {
	return &Submission{
		Creator: creator,
		Voters:  nil,
		URL:     url,
	}
}
