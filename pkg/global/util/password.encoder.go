package util

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordEncode(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func PasswordVerify(hashVal, userPw string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
    if err != nil {
        return false
    } else {
        return true
    }
}