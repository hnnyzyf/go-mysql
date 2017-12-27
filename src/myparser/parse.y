//头文件
%{
package myparser
%}



//类型定义
%union {
  keyword       struct{}

}


%token LEX_ERROR
%token <keyword> ALL AS ASC
	BY
	DELETE DISTINCT DESC DUPLICATE DEFAULT
	EXISTS
	FROM FOR
	GROUP
	HAVING
	INTO INSERT
	KEY
	LIMIT LOCK
	NULL
	ORDER
	PRIMARY
	SELECT SET
	UPDATE
	VALUES
	WHERE

%%