package util

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
)

// MD5sum hash
func MD5sum(content string) string {
	h := md5.New()
	io.WriteString(h, content)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA256 hash
func SHA256(content string) string {
	h := sha256.New()
	io.WriteString(h, content)
	return fmt.Sprintf("%x", h.Sum(nil))
}
