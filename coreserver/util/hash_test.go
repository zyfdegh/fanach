package util

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMD5sum(t *testing.T) {
	var cases = []struct {
		in  string
		out string
	}{
		{in: "tom", out: "34b7da764b21d298ef307d04d8152dc5"},
		{in: "bob", out: "9f9d51bc70ef21ca5c14f307980a29d8"},
	}

	for _, c := range cases {
		got := MD5sum(c.in)
		assert.Equal(t, c.out, got)
	}
}
