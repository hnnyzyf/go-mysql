package myparser

import "fmt"


//返回token类型
func (token *Tokener) Lex(lval *yySymType) int {
	tok,val:=token.Scan()
	//comment应由词法分析器获得
	if tok==COMMENT{
		tok,val=token.Scan()
	}
	//判断正确的类型
	switch tok{
		case NUMBER,IDENT,STRING:
			lval.val=val
	}
	return tok
}


//返回错误信息
func (token *Tokener) Error(err string) {
	token.LastError=err
	fmt.Println("ERROR:",err)
	fmt.Println("ERROR:",string(token.Lastchar),token.Position)
}