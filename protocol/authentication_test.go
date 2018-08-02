package protocol

import (
	. "gopkg.in/check.v1"
)

type MyAuthSuite struct{}

var _ = Suite(&MyAuthSuite{})

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
func (s *MyAuthSuite) TestSPA(c *C) {
	m, err := NewMethod("mysql_native_password")
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
