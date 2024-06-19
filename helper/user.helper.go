package helper

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"unicode"
)

func CheckPassLength(pass string) bool {
	return len(pass) >= 8
}

func CheckEmail(data string) bool {
	return strings.Contains(data, "@") && strings.HasSuffix(data, ".com")
}

func CheckUsername(data string) bool {
	return !strings.Contains(data, " ")
}

func CheckFullName(data string) bool {
	for _, char := range data {
		if !('a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == ' ') {
			return false
		}
	}
	return true
}

func HasUpperLetter(password string) bool {
	for _, c := range password {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}

func HasLowerLetter(password string) bool {
	for _, c := range password {
		if unicode.IsLower(c) {
			return true
		}
	}
	return false
}

func HasNumber(password string) bool {
	for _, c := range password {
		if unicode.IsNumber(c) {
			return true
		}
	}
	return false
}

func HasSpecialChar(password string) bool {
	for _, c := range password {
		if unicode.IsPunct(c) || unicode.IsSymbol(c) {
			return true
		}
	}
	return false
}

func GenerateHash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
