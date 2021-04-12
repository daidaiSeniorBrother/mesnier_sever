package utils

import (
	"github.com/satori/go.uuid"
)

/**
Version 1,基于 timestamp 和 MAC address (RFC 4122)
Version 2,基于 timestamp, MAC address 和 POSIX UID/GID (DCE 1.1)
Version 3, 基于 MD5 hashing (RFC 4122)
Version 4, 基于 random numbers (RFC 4122)
Version 5, 基于 SHA-1 hashing (RFC 4122)
*/
func UUID() string {
	return uuid.NewV4().String()
}
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
