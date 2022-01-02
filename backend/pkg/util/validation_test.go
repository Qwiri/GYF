package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURLHash(t *testing.T) {
	type test struct {
		url string
		// exh - expected hash
		exh string
		// exe - expected error
		exe string
	}
	tests := []test{
		{
			url: "https://giphy.com/test",
			exh: "giphy.com:test",
		},
		{
			url: "https://media.giphy.com/test",
			exh: "media.giphy.com:test",
		},
		{
			url: "https://media.giphy.com/test?query=ok",
			exh: "media.giphy.com:test",
		},
		{
			url: "https://media.giphy.com/test#anchor",
			exh: "media.giphy.com:test",
		},
		{
			url: "https://media.giphy.com/test#anchor?query=ok",
			exh: "media.giphy.com:test",
		},
		{
			url: "abc",
			exh: ":abc",
		},
	}
	for _, tst := range tests {
		hash, err := URLHash(tst.url)
		if err != nil || tst.exe != "" {
			assert.Equal(t, tst.exe, err.Error())
		}
		assert.Equal(t, tst.exh, hash)
	}
}
