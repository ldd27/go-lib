package tool

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(s string) string {
	t := md5.New()
	io.WriteString(t, s)
	return fmt.Sprintf("%x", t.Sum(nil))
}
