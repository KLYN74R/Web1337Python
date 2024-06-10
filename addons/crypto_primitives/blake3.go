package crypto_primitives

import (
	"encoding/hex"

	"lukechampine.com/blake3"
)

func Blake3Hash(msg string) string {

	msg_as_bytes := []byte(msg)

	blake3Hash := blake3.Sum256(msg_as_bytes)

	return hex.EncodeToString(blake3Hash[:])

}
