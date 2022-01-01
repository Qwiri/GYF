package util

import "github.com/Qwiri/GYF/backend/pkg/model"

func WrapClientArray(clients ...*model.Client) (res []interface{}) {
	res = make([]interface{}, len(clients))
	for i, v := range clients {
		res[i] = v
	}
	return
}

func WrapStringArray(strings ...string) (res []interface{}) {
	res = make([]interface{}, len(strings))
	for i, v := range strings {
		res[i] = v
	}
	return
}

func WrapVoteResults(result ...*model.SubmissionResult) (res []interface{}) {
	res = make([]interface{}, len(result))
	for i, v := range result {
		res[i] = v
	}
	return
}
