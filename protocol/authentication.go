package protocol

import (
	"crypto/sha1"

	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/hnnyzyf/go-mysql/util/random"
	"github.com/juju/errors"
)

type Method interface {
	//The Server use this method to create a Key
	CreateKey() []byte

	//The client use this method to create a response inforamtion
	EncodeKey([]byte, string) []byte
}

//Create a authentication method
func GetMethod(name uint8) (Method, error) {
	switch name {
	case mysql_native_password:
		return SPA, nil
	case ssl:
		return nil,nil 
	default:
		return nil, errors.Errorf("Invalid authentication method:%s", name)
	}
}

var SPA = &spa{}

//Represent the mysql_native_password method
//using a tested, crypto-graphic hashing function which isn't broken
//knowning the content of the hash in the mysql.user table isn't enough to authenticate against the MySQL Server.
type spa struct{}

//CreateKey return a 20 bytes random letters
//the mysql server store the SHA1(passwd)
func (a *spa) CreateKey() []byte {
	return random.String(20)
}

//EncodeKey use scramble and passwd to create a response
//read https://dev.mysql.com/doc/internals/en/secure-password-authentication.html
//SHA1( password ) XOR SHA1( "20-bytes random data from server" <concat> SHA1( SHA1( password ) ) )
func (a *spa) EncodeKey(scramble []byte, passwd string) []byte {
	//sha1(password)
	h := sha1.New()
	h.Write(hack.Slice(passwd))
	hash1 := h.Sum(nil)

	//sha1(sha1(password))
	h.Reset()
	h.Write(hash1)
	hash2 := h.Sum(nil)

	//SHA1( "20-bytes random data from server" <concat> SHA1( SHA1( password ) ) )
	h.Reset()
	h.Write(scramble)
	h.Write(hash2)
	hash3 := h.Sum(nil)

	//XOR
	for index := range hash3 {
		hash3[index] ^= hash1[index]
	}

	return hash3
}
