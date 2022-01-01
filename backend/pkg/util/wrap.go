package util

func WrapStringArray(strings ...string) (res []interface{}) {
	res = make([]interface{}, len(strings))
	for i, v := range strings {
		res[i] = v
	}
	return
}
