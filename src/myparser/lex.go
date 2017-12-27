package myparser

import (
	"bytes"
	"fmt"
	"strings"
)

//实现mysql的此法解析器
//定义关键子

//需要实现Lex方法和error方法
type Tokener struct {
	
}

//返回token类型
func (token *Tokener) Lex(v *yySymType) int {

}


//返回错误信息
func (token *Tokener) Error(err string) {

}
