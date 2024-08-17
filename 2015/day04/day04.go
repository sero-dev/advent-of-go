package day04

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

func Problem(secret string, prefix string) (uint, error) {
	var salt uint = 1

	for {
		saltString := fmt.Sprint(salt)
		checksum := md5.Sum([]byte(secret + saltString))
		hash := hex.EncodeToString(checksum[:])

		if strings.HasPrefix(hash, prefix) {
			break
		}

		if salt == ^uint(0) {
			return 0, errors.New("could not find salt that produced " + prefix + " prefix for MD5 hash")
		}

		salt++
	}

	return salt, nil
}
