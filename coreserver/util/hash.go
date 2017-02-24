package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

// MD5sum hash
func MD5sum(content string) string {
	h := md5.New()
	io.WriteString(h, content)
	return fmt.Sprintf("%x", h.Sum(nil))
}
