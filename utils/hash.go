package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, hashType string) (string, error) {
	switch hashType {
	case "bcrypt":
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		return string(hashedPassword), nil
	case "md5":
		hash := md5.Sum([]byte(password))
		return hex.EncodeToString(hash[:]), nil
	case "sha1":
		hash := sha1.Sum([]byte(password))
		return hex.EncodeToString(hash[:]), nil
	case "sha256":
		hash := sha256.Sum256([]byte(password))
		return hex.EncodeToString(hash[:]), nil
	default:
		return "", errors.New("unsupported hash type")
	}
}
