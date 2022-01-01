package model

import "strings"

type (
	SubmissionArray []*Submission
	SubmissionMap   map[string]*Submission
)

type Submission struct {
	Creator *Client
	Voters  ClientArray
	URL     string
}

func NewSubmission(creator *Client, url string) *Submission {
	return &Submission{
		Creator: creator,
		URL:     url,
	}
}

func (M SubmissionMap) HasVoted(client *Client) bool {
	for _, sub := range M {
		for _, vot := range sub.Voters {
			if vot == client {
				return true
			}
		}
	}
	return false
}

func (M SubmissionMap) ByURL(url string) (*Submission, bool) {
	for _, sub := range M {
		if strings.EqualFold(url, sub.URL) {
			return sub, true
		}
	}
	return nil, false
}

func (M SubmissionMap) HasSubmittedGIF(client *Client) (ok bool) {
	_, ok = M[client.Name]
	return
}

func (M SubmissionMap) AllExceptFrom(client *Client) (res SubmissionArray) {
	for name, sub := range M {
		if name == client.Name {
			continue
		}
		res = append(res, sub)
	}
	return
}

func (M SubmissionMap) AsArray() (res SubmissionArray) {
	res = make(SubmissionArray, len(M))
	i := 0
	for _, v := range M {
		res[i] = v
		i += 1
	}
	return
}

func (A SubmissionArray) URLs() (res []string) {
	res = make([]string, len(A))
	for i, v := range A {
		res[i] = v.URL
	}
	return
}

type SubmissionResult struct {
	URL     string   `json:"url"`
	Creator string   `json:"creator"`
	Voters  []string `json:"voters"`
}

func (A SubmissionArray) Results() (res []*SubmissionResult) {
	for _, sub := range A {
		voters := make([]string, len(sub.Voters))
		for i, voter := range sub.Voters {
			voters[i] = voter.Name
		}
		res = append(res, &SubmissionResult{
			URL:     sub.URL,
			Creator: sub.Creator.Name,
			Voters:  voters,
		})
	}
	return
}
