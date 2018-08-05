package client

import (
	"fmt"
	"testing"
)

var data = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func read(b []byte) {
	l := len(b)
	b = data[:l]
}

func Test_copy(t *testing.T) {
	x := make([]byte, 5)

	read(x)

	fmt.Println(x)

}
