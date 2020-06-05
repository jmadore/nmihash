// Package nmihash is used as a source cdoe example for a Go (golang) implementation of the examples found at:
// https://support.nmi.com/hc/en-gb/articles/115001683343-Implementation-Generating-a-Hash-Code
package nmihash

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

type constError string

func (err constError) Error() string {
	return string(err)
}

// Errors raised by package nmihash.
const (
	ErrInvalidParameters = constError("invalid parameters")
)

// CalculateHash return a SHA1 hash (in Base64 encoding) of the hashKey, terminalID, refNum, and amount provided as input.
func CalculateHash(hashKey string, terminalID string, refNum string, amount string) (string, error) {

	if hashKey == "" || terminalID == "" || refNum == "" || amount == "" {
		return "", ErrInvalidParameters
	}

	var s strings.Builder
	s.WriteString(hashKey)
	s.WriteString(terminalID)
	s.WriteString(refNum)
	s.WriteString(amount)

	data := []byte(s.String())
	hash := sha1.Sum(data)
	encoded := base64.StdEncoding.EncodeToString(hash[:])

	return encoded, nil
}
