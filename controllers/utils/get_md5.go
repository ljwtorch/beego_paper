package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(pwd string) string {
	//返回一个哈希，用h接收
	h := md5.New()
	h.Write([]byte(pwd))

	return hex.EncodeToString(h.Sum(nil))
}
