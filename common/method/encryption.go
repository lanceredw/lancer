package method

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	hashes := md5.New()
	hashes.Write([]byte(str))

	hashInBytes := hashes.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}
