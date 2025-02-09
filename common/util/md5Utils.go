package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5encoded(pwd, salt string) string {
	m5 := md5.New()
	m5.Write([]byte(pwd))
	m5.Write([]byte(salt))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}
func Md5Verify(pwd, salt, encodedPwd string) bool {
	encoded := Md5encoded(pwd, salt)
	if encodedPwd == encoded {
		return true
	}
	return false
}
