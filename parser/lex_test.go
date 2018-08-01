package parser

import (
	"regexp"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

//this section give the lex defination
var pattern = map[string]string{
	"letter":      "[a-zA-Z]",
	"digit":       "[0-9]",
	"identletter": "letter|digit|byte>1",
	"string1":     `"([^"]|"{2})*"`,
	"string2":     `'([^']|'{2})*'`,
	"identifier1": `_identletter*`,
	"identifier2": "`([^`]|`{2})*`",
	"identifier3": "identletter+",
	"integer1":    "[0-9]+",
	"integer2":    `([0-9]+\.?[0-9]*|[0-9]*\.?[0-9]+)E[+|-]?[0-9]+`,
	"decimal":     `[0-9]+\.[0-9]*|[0-9]*\.[0-9]+`,
	"comment":     `/\*.*?\*/`,
}

func (s *MySuite) TestString1(c *C) {
	r, err := regexp.Compile(pattern["string1"])
	if err != nil {
		c.Error(err)
	}

	res := []struct {
		str string
		ok  string
	}{
		{`"abcdefg"`, `"abcdefg"`},
		{`"
			dasdasd"`, `"
			dasdasd"`},
		{`""dsadasdasds"`, `""`},
		{`"dasd""asd"`, `"dasd""asd"`},
		{`"dasd''asd"`, `"dasd''asd"`},
	}

	for i := range res {
		c.Assert(string(r.Find([]byte(res[i].str))), Equals, res[i].ok)
	}
}

func (s *MySuite) TestString2(c *C) {
	r, err := regexp.Compile(pattern["string2"])
	if err != nil {
		c.Error(err)
	}

	res := []struct {
		str string
		ok  string
	}{
		{`'abcdefg'`, `'abcdefg'`},
		{`'
			dasdasd'`, `'
			dasdasd'`},
		{`''dsadasdasds'`, `''`},
		{`'dasd''asd'`, `'dasd''asd'`},
		{`'dasd''asd'`, `'dasd''asd'`},
	}

	for i := range res {
		c.Assert(string(r.Find([]byte(res[i].str))), Equals, res[i].ok)
	}
}

func (s *MySuite) TestIdentifier2(c *C) {
	r, err := regexp.Compile(pattern["identifier2"])
	if err != nil {
		c.Error(err)
	}

	res := []struct {
		str string
		ok  string
	}{
		{"`abcdefg`", "`abcdefg`"},
		{"`    dasdasd`", "`    dasdasd`"},
		{"``dsadasdasds`", "``"},
		{"`dasd``asd`", "`dasd``asd`"},
		{"`dasd``asd`", "`dasd``asd`"},
	}

	for i := range res {
		c.Assert(string(r.Find([]byte(res[i].str))), Equals, res[i].ok)
	}
}

func (s *MySuite) TestInteger1(c *C) {
	r, err := regexp.Compile(pattern["integer1"])
	if err != nil {
		c.Error(err)
	}

	res := []struct {
		str string
		ok  string
	}{
		{"123457676abcd", "123457676"},
		{"dasd123456", "123456"},
	}

	for i := range res {
		c.Assert(string(r.Find([]byte(res[i].str))), Equals, res[i].ok)
	}
}

func (s *MySuite) TestInteger2(c *C) {
	r, err := regexp.Compile(pattern["integer2"])
	if err != nil {
		c.Error(err)
	}

	res := []struct {
		str string
		ok  string
	}{
		{"asdasd123.3E+5dasdasd", "123.3E+5"},
		{"asdasd123.3E-5dasdasd", "123.3E-5"},
		{"asdasd123.3E5dasdasd", "123.3E5"},
		{"asdasd123.3E050dasdasd", "123.3E050"},
		{"asdasd123.3E-05dasdasd", "123.3E-05"},
		{"asdasd123.aa3E-05dasdasd", "3E-05"},
	}

	for i := range res {
		c.Assert(string(r.Find([]byte(res[i].str))), Equals, res[i].ok)
	}
}

func (s *MySuite) TestDecimal(c *C) {
	r, err := regexp.Compile(pattern["decimal"])
	if err != nil {
		c.Error(err)
	}

	res := []struct {
		str string
		ok  string
	}{
		{"ddd1.2aaaaaa", "1.2"},
		{"ddd.2ccc", ".2"},
		{"ddddsd2323.ccc", "2323."},
		{"ddddsd2xxxx323.ccc", "323."},
	}

	for i := range res {
		c.Assert(string(r.Find([]byte(res[i].str))), Equals, res[i].ok)
	}
}

func (s *MySuite) TestComment(c *C) {
	r, err := regexp.Compile(pattern["comment"])
	if err != nil {
		c.Error(err)
	}

	res := []struct {
		str string
		ok  string
	}{
		{"ddd/*213445436*/aa", "/*213445436*/"},
		{"ddd/*21344*/5436*/aa", "/*21344*/"},
		{"ddd/*21344/5436*/aa", "/*21344/5436*/"},
		{"ddd/*21344*5436*/aa", "/*21344*5436*/"},
	}

	for i := range res {
		c.Assert(string(r.Find([]byte(res[i].str))), Equals, res[i].ok)
	}
}
