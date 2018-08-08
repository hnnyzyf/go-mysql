package protocol

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyAuthSuite struct{}

var _ = Suite(&MyAuthSuite{})

//use following scrip as another test
//from six import byte2int,int2byte
//import struct
//import io
//
//SCRAMBLE_LENGTH_323 = 8
//
//class RandStruct_323(object):
//
//    def __init__(self, seed1, seed2):
//        self.max_value = 0x3FFFFFFF
//        self.seed1 = seed1 % self.max_value
//        self.seed2 = seed2 % self.max_value
//
//    def my_rnd(self):
//        self.seed1 = (self.seed1 * 3 + self.seed2) % self.max_value
//        self.seed2 = (self.seed1 + self.seed2 + 33) % self.max_value
//        return float(self.seed1) / float(self.max_value)
//
//
//def scramble_old_password(password, message):
//    """Scramble for old_password"""
//    hash_pass = _hash_password_323(password)
//    hash_message = _hash_password_323(message[:SCRAMBLE_LENGTH_323])
//    hash_pass_n = struct.unpack(">LL", hash_pass)
//    hash_message_n = struct.unpack(">LL", hash_message)
//
//    rand_st = RandStruct_323(
//        hash_pass_n[0] ^ hash_message_n[0], hash_pass_n[1] ^ hash_message_n[1]
//    )
//    outbuf = io.BytesIO()
//    for _ in range(min(SCRAMBLE_LENGTH_323, len(message))):
//        outbuf.write(int2byte(int(rand_st.my_rnd() * 31) + 64))
//    extra = int2byte(int(rand_st.my_rnd() * 31))
//    out = outbuf.getvalue()
//    outbuf = io.BytesIO()
//    for c in out:
//        outbuf.write(int2byte(byte2int(c) ^ byte2int(extra)))
//    return outbuf.getvalue()
//
//def _hash_password_323(password):
//    nr = 1345345333
//    add = 7
//    nr2 = 0x12345671
//
//    # x in py3 is numbers, p27 is chars
//    for c in [byte2int(x) for x in password if x not in (' ', '\t', 32, 9)]:
//        nr ^= (((nr & 63) + add) * c) + (nr << 8) & 0xFFFFFFFF
//        nr2 = (nr2 + ((nr2 << 8) ^ nr)) & 0xFFFFFFFF
//        add = (add + c) & 0xFFFFFFFF
//
//    r1 = nr & ((1 << 31) - 1)  # kill sign bits
//    r2 = nr2 & ((1 << 31) - 1)
//    return struct.pack(">LL", r1, r2)
//
//scramble=["sstvy~&AHQZju^GVk!EX","ppqsvz$)EMWgr#DShxBU","##$^(AFLTbku%ESgv_Qi"]
//password=["123456\t78","12345 678","C0mpl!ca ted#PASS123"]
//
//for i in range(len(scramble)):
//    print scramble_old_password(password[i],scramble[i]).encode("hex")
func (s *MyAuthSuite) TestoldPassword(c *C) {
	m, err := GetMethod("mysql_old_password")
	if err != nil {
		c.Error(err)
	}

	data := []struct {
		scramble []byte
		password string
		key      []byte
	}{
		{[]byte("sstvy~&AHQZju^GVk!EX"), "123456\t78", []byte{0x46, 0x4a, 0x55, 0x5c, 0x5c, 0x5b, 0x41, 0x49}},
		{[]byte("ppqsvz$)EMWgr#DShxBU"), "12345 678", []byte{0x43, 0x53, 0x5b, 0x52, 0x4b, 0x43, 0x57, 0x4a}},
		{[]byte("##$^(AFLTbku%ESgv_Qi"), "C0mpl!ca ted#PASS123", []byte{0x4d, 0x5d, 0x4c, 0x53, 0x47, 0x5c, 0x55, 0x4d}},
	}

	for i := range data {
		c.Assert(m.EncodeKey(data[i].scramble, data[i].password), DeepEquals, data[i].key)
	}

}

//use following script as the another test
//#!/usr/bin/env python
//#coding=utf-8
//
//import hashlib
//
//def stringxor(str1,str2):
//    orxstr=""
//    for i in range(0,len(str1)):
//        rst=ord(list(str1)[i])^ord(list(str2)[i])
//        orxstr=orxstr+ chr(rst)
//    return orxstr
//
//def encode(s,p):
//    x1=hashlib.sha1(p).hexdigest()
//    x2=hashlib.sha1(x1.decode("hex")).hexdigest()
//    part2=s+x2.decode("hex")
//    x3=hashlib.sha1(part2).hexdigest()
//    return stringxor(x1.decode("hex"),x3.decode("hex"))
//
//scramble=["sstvy~&AHQZju^GVk!EX","ppqsvz$)EMWgr#DShxBU","##$^(AFLTbku%ESgv_Qi"]
//
//passwd="12345678"
//for s in scramble:
//    res=encode(s,passwd)
//    print [ord(j) for j in list(res)]
func (s *MyAuthSuite) TestsecurePasswprd(c *C) {
	m, err := GetMethod("mysql_native_password")
	if err != nil {
		c.Error(err)
	}

	data := []struct {
		scramble []byte
		password string
		key      []byte
	}{
		{[]byte("sstvy~&AHQZju^GVk!EX"), "12345678", []byte{0xa8, 0x4, 0x2b, 0xd9, 0x83, 0x17, 0x7f, 0x29, 0xed, 0x43, 0x77, 0xdc, 0x36, 0xd6, 0xae, 0xb8, 0x9e, 0x11, 0x9b, 0xc9}},
		{[]byte("ppqsvz$)EMWgr#DShxBU"), "12345678", []byte{0x32, 0xff, 0x25, 0xb, 0x15, 0x74, 0x4c, 0x19, 0x63, 0xab, 0xb9, 0x75, 0xf1, 0xe, 0xc3, 0xf2, 0x82, 0x89, 0xec, 0xa5}},
		{[]byte("##$^(AFLTbku%ESgv_Qi"), "12345678", []byte{0xd3, 0xf7, 0x90, 0xbf, 0xfb, 0xa8, 0x77, 0x8c, 0x14, 0xc7, 0x74, 0xf4, 0x7e, 0x50, 0x2f, 0x25, 0x73, 0xa, 0xb4, 0xd4}},
	}

	for i := range data {
		c.Assert(m.EncodeKey(data[i].scramble, data[i].password), DeepEquals, data[i].key)
	}
}
