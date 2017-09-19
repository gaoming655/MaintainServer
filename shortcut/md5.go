package shortcut

import (
	"crypto/md5"
	"fmt"
)

// 生成md5
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

