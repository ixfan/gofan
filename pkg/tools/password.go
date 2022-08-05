package tools

import (
	"encoding/base64"
	"github.com/gookit/goutil/strutil"
	"strings"
)

// PasswordKey 16位Key
var PasswordKey = "eg5jpqvoUoEFmaa8"

// EncryptPassword 加密密码
func EncryptPassword(password string) (string, error) {
	salt := strutil.RandomCharsV3(16)
	saltBytes, err := AesCFBEncrypt([]byte(salt), []byte(PasswordKey))
	if err != nil {
		return "", err
	}
	saltCrypt := base64.StdEncoding.EncodeToString(saltBytes)
	passwordBytes, err := AesCFBEncrypt([]byte(password+"."+salt), []byte(PasswordKey))
	if err != nil {
		return "", err
	}
	passwordCrypt := base64.StdEncoding.EncodeToString(passwordBytes)
	return passwordCrypt + "." + saltCrypt, nil
}

// DecryptPassword 解密
func DecryptPassword(passwordCrypt string) (string, error) {
	crypts := strutil.Split(passwordCrypt, ".")
	saltBytes, err := base64.StdEncoding.DecodeString(crypts[1])
	if err != nil {
		return "", err
	}
	salt, err := AesCFBDecrypt(saltBytes, []byte(PasswordKey))
	if err != nil {
		return "", err
	}
	passwordBytes, err := base64.StdEncoding.DecodeString(crypts[0])
	if err != nil {
		return "", err
	}
	passwordSalt, err := AesCFBDecrypt(passwordBytes, []byte(PasswordKey))
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(passwordSalt), "."+string(salt), ""), nil
}
