package random

import (
	"math/rand"
)

var char = []byte("abcdefghijklmnopqrstuvwxyz!@~#$%^&*()_+ABCDEFGHIJKLMOPQRSTUVWXYZ")

//RandomString generate n bytes string
func String(n int) []byte {
	b := make([]byte, n)
	seed := rand.Int()
	for i := range b {
		seed = seed % len(char)
		b[i] = char[seed]
		seed = seed + i
	}
	return b
}
