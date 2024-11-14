package gochannel

import (
	"strings"
	"github.com/tuhinal/go-experiment/go-basic/cryptit/encrypt"
)

func ToUppercase() string{
encryptedUppercaseName := strings.ToUpper(encrypt.Encrypt("Tuhin"))
return encryptedUppercaseName
}
func ToLowercase() string{
encryptedLowercaseName := strings.ToLower(encrypt.Encrypt("Tuhin"))
return encryptedLowercaseName
}