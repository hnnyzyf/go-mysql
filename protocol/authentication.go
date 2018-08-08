package protocol

import (
	"crypto/sha1"

	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/hnnyzyf/go-mysql/util/hash"
	"github.com/hnnyzyf/go-mysql/util/random"
	"github.com/juju/errors"
)

type Method interface {
	//The Server use this method to create a Key
	CreateKey() []byte

	//The client use this method to create a response inforamtion
	EncodeKey([]byte, string) []byte

	//Name return the name of plugin name
	Name() string
}

var SPA = &securePassword{}
var OPA = &oldPassword{}
var CPA = &clearPassword{}

//Create a authentication method
func GetMethod(name string) (Method, error) {
	switch name {
	case "mysql_old_password":
		return OPA, nil
	case "mysql_native_password":
		return SPA, nil
	case "mysql_clear_password":
		return CPA, nil
	case "sha256_password":
		return nil, errors.Errorf("Protocol:the sha256_password is not implemented")
	default:
		return nil, errors.Errorf("Protocol:invalid authentication method:%s", name)
	}
}

//Represent the mysql_old_password method
//The hashing algorithm used for this auth method is broken as shown at http://sqlhack.com/ and CVE-2000-0981
type oldPassword struct{}

////CreateKey return a 20 bytes random letters
//if need,the server will return only 8 bytes
func (a *oldPassword) CreateKey() []byte {
	return random.String(20)
}

//EncodeKey use scramble and passwd to create a response
func (a *oldPassword) EncodeKey(scramble []byte, passwd string) []byte {
	scramble = scramble[:8]

	//calculate hash
	hw := hash.Hash(hack.Slice(passwd))
	hs := hash.Hash(scramble)

	//get random seed
	seed := hash.NewSeed(hw[0]^hs[0], hw[1]^hs[1])

	r := make([]byte, 8)

	for i := range r {
		r[i] = seed.Next() + 64
	}

	mask := seed.Next()
	for i := range r {
		r[i] ^= mask
	}

	return r
}

func (a *oldPassword) Name() string{
	return "mysql_old_password"
}

//Represent the mysql_native_password method
//using a tested, crypto-graphic hashing function which isn't broken
//knowning the content of the hash in the mysql.user table isn't enough to authenticate against the MySQL Server.
type securePassword struct{}

//CreateKey return a 20 bytes random letters
//the mysql server store the SHA1(passwd)
func (a *securePassword) CreateKey() []byte {
	return random.String(20)
}

//EncodeKey use scramble and passwd to create a response
//read https://dev.mysql.com/doc/internals/en/secure-password-authentication.html
//SHA1( password ) XOR SHA1( "20-bytes random data from server" <concat> SHA1( SHA1( password ) ) )
func (a *securePassword) EncodeKey(scramble []byte, passwd string) []byte {
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

func (a *securePassword) Name() string{
	return "mysql_native_password"
}

//Reprensent the mysql_clear_password method
//https://dev.mysql.com/doc/internals/en/clear-text-authentication.html
type clearPassword struct{}

func (a *clearPassword) CreateKey() []byte {
	return []byte{}
}

func (a *clearPassword) EncodeKey(scramble []byte, passwd string) []byte {
	return hack.Slice(passwd)
}

func (a *clearPassword) Name() string{
	return "mysql_clear_password"
}


//Reprensent the sha256_password method
//https://dev.mysql.com/doc/internals/en/sha256.html
type sha256Password struct{}

func (a *sha256Password) CreateKey() []byte {
	return []byte{}
}

func (a *sha256Password) EncodeKey(scramble []byte, passwd string) []byte {
	return hack.Slice(passwd)
}

func (a *sha256Password) Name() string{
	return "sha256_password"
}