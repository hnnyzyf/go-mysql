package protocol

import (
	"crypto/sha1"
)

//read https://dev.mysql.com/doc/internals/en/secure-password-authentication.html
//SHA1( password ) XOR SHA1( "20-bytes random data from server" <concat> SHA1( SHA1( password ) ) )
func CalAuthResponse(key []byte, passwd string) []byte {
	h := sha1.New()
	h.Write([]byte(passwd))
	hash1 := h.Sum(nil)

	h.Reset()
	h.Write(hash1)
	hash2 := h.Sum(nil)

	h.Reset()
	h.Write(key)
	h.Write(hash2)
	hash3 := h.Sum(nil)

	for index := range hash3 {
		hash3[index] ^= hash1[index]
	}

	return hash3
}
