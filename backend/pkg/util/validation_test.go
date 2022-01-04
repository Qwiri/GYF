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

func TestIsAllowed(t *testing.T) {
	exa := []string{
		"https://media.tenor.com/images/45719ed0533eaed76ede0f844781e962/tenor.gif",
		"https://media.tenor.com/images/2543da98e10b0429bf74fddfb99e1117/tenor.gif",
		"https://media.tenor.com/images/39facd5b200e21220dbaa440e66c98b9/tenor.gif",
		"https://media3.giphy.com/media/QYwEBOGAyCCV56WqlT/giphy.gif?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
		"https://media4.giphy.com/media/3ohhwJ1t8NzJDcJTVu/giphy.gif?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
		"https://media2.giphy.com/media/3otPotoOjtt6409ORG/giphy.gif?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
	}
	exd := []string{
		"https://fake.tenor.com/images/39facd5b200e21220dbaa440e66c98b9/tenor.gif",
		"https://fake.giphy.com/media/3ohhwJ1t8NzJDcJTVu/giphy.gif?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
		"https://medya2.giphy.com/media/3otPotoOjtt6409ORG/giphy.gif?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
		"https://media.tenor.com/images/45719ed0533eaed76ede0f844781e962/tenor-abc.",
		"https://media.tenor.com/images/45719ed0533eaed76ede0f844781e962/tenor.png",
		"https://media.tenor.com/images/45719ed0533eaed76ede0f844781e962/tenor.mp4",
		"https://media3.giphy.com/media/QYwEBOGAyCCV56WqlT/giphy.gyf?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
		"https://media3.giphy.com/media/QYwEBOGAyCCV56WqlT/giphy.png?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
		"https://media3.giphy.com/media/QYwEBOGAyCCV56WqlT/giphy.mp4?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
		"https://media3.giphy.com/media/QYwEBOGAyCCV56WqlT/giphy.mov?cid=df4eeb68tyb2fkp73qli92eyu0ilynjuk62ex2vfuc3pk5kw&rid=giphy.gif&ct=g",
	}

	for _, a := range exa {
		allowed, err := IsURLAllowed(a)
		assert.Nil(t, err)
		assert.True(t, allowed)
	}
	for _, a := range exd {
		disallowed, err := IsURLAllowed(a)
		assert.Nil(t, err)
		assert.False(t, disallowed)
	}
}
