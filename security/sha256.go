package security

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256HexEecode returns a hex encoded string of sha256 hashed string.
func Sha256HexEecode(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}
