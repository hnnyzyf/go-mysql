package parser

//返回Lex
func (token *Tokener) Lex(lval *yySymType) int {
	tok, val := token.Scan()
	//comment应由词法分析器获得
	if tok == COMMENT {
		token.Comment = append(token.Comment, val)
		tok, val = token.Scan()
	}
	//判断正确的类型
	switch tok {
	case IDENT, BuiltinCharacterIdent, BuiltinFucTimeAddIdent, BuiltinFucTimeSubIdent, BuiltinTimeUnitIdent,BuiltinCharacterName:
		lval.ident = val
	case PARAM_MARKER, DOUBLE_QUOTA_STRING, SINGLE_QUOTA_STRING:
		lval.str = val
	case BIT_NUMBER, HEX_NUMBER, INTEGER, FLOAT:
		lval.val = val
	}

	return tok
}

//返回错误信息
func (token *Tokener) Error(err string) {
	token.LastError = err
}

//add yyGetToken
func yyGetToken(character int, token int, lval *yySymType) string {
	switch character {
	case IDENT, BuiltinCharacterIdent, BuiltinFucTimeAddIdent, BuiltinFucTimeSubIdent, BuiltinTimeUnitIdent,BuiltinCharacterName:
		return lval.ident
	case PARAM_MARKER, DOUBLE_QUOTA_STRING, SINGLE_QUOTA_STRING:
		return lval.str
	case BIT_NUMBER, HEX_NUMBER, INTEGER, FLOAT:
		return lval.val
	}
	return yyTokname(token)
}