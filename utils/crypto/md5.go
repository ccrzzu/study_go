package crypto

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func GetMD5UUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return MD5(base64.URLEncoding.EncodeToString(b))
}

func MD5(key string) (result string) {
	data := []byte(key)
	has := md5.Sum(data)
	//将[]byte转成16进制字符串
	result = fmt.Sprintf("%x", has)
	return result
}