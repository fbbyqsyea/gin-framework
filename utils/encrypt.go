package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(s string) string {
	b := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(b))
}
