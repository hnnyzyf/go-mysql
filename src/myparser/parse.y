//头文件
%{
package myparser

import AST "myast"

func SetParseTree(yylex interface{}, stmt AST.SelectNode) {
  yylex.(*Tokener).ASTree = stmt
}
%}



//类型定义
%union {
  keyword       string 	
  value  		string	
  node          AST.Node
  snode			AST.SelectNode
  expr 			AST.ExprNode
  list          AST.List
}


//token定义
%token <value> IDENT STRING NUMBER 
%token <keyword> AS ASC ALL ANY
	BY 
	COMMENT CHARACTER COLLATE
    DISTINCT DISTINCTROW DESC DAY DAY_HOUR DAY_MICROSECOND DAY_MINUTE DAY_SECOND
	EXISTS
	FROM FOR FALSE
	GROUP
	HAVING HOUR HOUR_MICROSECOND HOUR_MINUTE HOUR_SECOND 
	INTO IF INTERVAL
	LIMIT LOCK 
	MINUTE MINUTE_SECOND MODE MONTH MICROSECOND MINUTE_MICROSECOND
	NULL
	ORDER OFFSET 
	QUARTER
	SELECT SOME SHARE SECOND SECOND_MICROSECOND SET
	TO  TRUE
	UPDATE USING
	WHERE WEEK
	YEAR YEAR_MONTH
	'(' '~'

       
           
//优先级定义

%left <keyword> UNION EXCEPT 
%left <keyword> INTERSECT
%left <keyword> ','
%left <keyword> JOIN STRAIGHT_JOIN LEFT RIGHT INNER OUTER CROSS NATURAL USE
%left <keyword> ON
%left <keyword> OR
%left <keyword> AND
%right <keyword> NOT
%left <keyword> BETWEEN CASE WHEN THEN ELSE
%left <keyword> '=' '<' '>' IS LIKE IN GE LE NE IE LNE SL SR LEG AA OO
%left <keyworkd> OP
%left <keyword> '|'
%left <keyword> '&'
%left <keyword> '+' '-'
%left <keyword> '*' '/' '%'
%left <keyword> '^'
%right <keyword> UMINUS
%nonassoc <keyword> '.'
%left <keyword> END

//类型定义

%type<snode> SelectStmt select_no_parens select_with_parens simple_select select_clause

%type<list> opt_target_list target_list 

%type<expr> target_el

%type<value> opt_col_alias col_alias opt_table_alias table_alias ColLabel ColId

%type<expr> a_expr b_expr c_expr 

%type<list> in_expr 

%type<list> expr_list

%type<expr> func_expr func_arg_expr

%type<keyword> func_name 

%type<list> func_arg_list 

%type<keyword> opt_distinct 

%type<value> opt_character_set opt_collate_set

%type<value> time_expr time_type

%type<expr> case_expr case_arg case_default when_clause

%type<list> when_clause_list

%type<expr> columnref qualified_name

%type<value> indirection indirection_el  

%type<value> subquery_op subquery_type

%type<node> distinct_clause 

%type<keyword> all_or_distinct

%type<node> into_clause 

%type <value> into_table_name

%type<node> from_clause 

%type<list> from_list 

%type<expr> table_ref relation

%type<expr> joined_table join_qual name

%type<keyword> join_type  

%type<list> name_list 

%type<node> where_clause

%type<node> group_clause 

%type<list> group_by_list 

%type<expr> group_by_item

%type<node> having_clause

%type<node> opt_sort_clause sort_clause  

%type<list> sortby_list 

%type<expr> sortby  

%type<keyword> opt_asc_desc

%type<node> opt_for_locking_clause for_locking_clause

%type<node> opt_select_limit select_limit 

%type<list> limit_clause 

%type<expr> offset_clause 

%type<list> select_limit_value

%type<expr> select_offset_value


%start Start
%%


Start:
	SelectStmt
	{
		SetParseTree(yylex,$1)
	}


SelectStmt: 
		select_no_parens			%prec UMINUS
		| select_with_parens		%prec UMINUS
	


select_with_parens:
			'(' select_no_parens ')'				{ $$ = $2 }
			| '(' select_with_parens ')'			{ $$ = $2 }


select_no_parens:
			simple_select							{ $$ = $1 }
			| select_clause sort_clause        		
				{ 
					$1.(*AST.SimpleSelectStmt).Sortby = $2.(*AST.Sortclause)
					$$ = $1
				}
			| select_clause opt_sort_clause for_locking_clause opt_select_limit 
				{	
					if $2!=nil{
						$1.(*AST.SimpleSelectStmt).Sortby = $2.(*AST.Sortclause)
					}
					$1.(*AST.SimpleSelectStmt).Lock = $3.(*AST.Lockingclause)
					if $4!=nil{
						$1.(*AST.SimpleSelectStmt).Limit = $4.(*AST.Limitclause)
					}
					$$ = $1
				}
			| select_clause opt_sort_clause select_limit opt_for_locking_clause  
				{
					if $2!=nil{
						$1.(*AST.SimpleSelectStmt).Sortby = $2.(*AST.Sortclause)
					}
					if $4!=nil{
						$1.(*AST.SimpleSelectStmt).Lock = $4.(*AST.Lockingclause)
					}
					$1.(*AST.SimpleSelectStmt).Limit = $3.(*AST.Limitclause)
					$$ = $1
				}

select_clause:
			simple_select							{ $$ = $1 }
			| select_with_parens					{ $$ = $1 }


simple_select:
			SELECT opt_target_list into_clause from_clause where_clause group_clause having_clause 
				{	
					//创建一个新的Select结构体
					Select:=AST.NewSimpleSelectStmt()
					Select.NType=AST.AST_SELECT
					Select.Op=AST.SETOP_SIMPLE
					Select.Target = $2
					Select.Into = $3   
					Select.From = $4  
					Select.Where = $5    
					Select.Groupby = $6
					Select.Having = $7
					$$ = Select   
				}
			|SELECT distinct_clause target_list into_clause from_clause where_clause group_clause having_clause 
				{
					//创建一个新的Select结构体
					Select:=AST.NewSimpleSelectStmt()
					Select.NType=AST.AST_SELECT
					Select.Op=AST.SETOP_SIMPLE
					Select.Distinct = $2
					Select.Target = $3
					Select.Into = $4   
					Select.From = $5  
					Select.Where = $6    
					Select.Groupby = $7
					Select.Having = $8
					$$ = Select   
				}
			| select_clause UNION all_or_distinct select_clause
				{
					Stmt:=AST.NewSelectStmt()
					Stmt.NType = AST.AST_SELECT
					Stmt.Having=$2 
					Stmt.Op = AST.SETOP_UNION    
					Stmt.Left = $1
					Stmt.Right = $4
					$$ = Stmt
				}
			| select_clause INTERSECT all_or_distinct select_clause
				{
					Stmt:=AST.NewSelectStmt()
					Stmt.NType = AST.AST_SELECT
					Stmt.Having=$2 
					Stmt.Op = AST.SETOP_INTERSECT    
					Stmt.Left = $1
					Stmt.Right = $4
					$$ = Stmt
				}
			| select_clause EXCEPT all_or_distinct select_clause
				{
					Stmt:=AST.NewSelectStmt()
					Stmt.NType = AST.AST_SELECT
					Stmt.Having=$2 
					Stmt.Op = AST.SETOP_EXCEPT  
					Stmt.Left = $1
					Stmt.Right = $4
					$$ = Stmt
				}

/*****************************************************************************
 *
 *	distinct clause for SELECT
 *
 *****************************************************************************/

distinct_clause:
			DISTINCT 			
				{
					distinct:=AST.NewDistinctclause()
					$$ = distinct
				}
			|DISTINCTROW
				{
					distinct:=AST.NewDistinctclause()
					$$ = distinct
				}



all_or_distinct:
			ALL										
				{ 
					$$ = "ALL" 
				}
			| DISTINCT								
				{ 
					$$ = "DISTINCT" 
				}
			| /*EMPTY*/								
				{ 
					$$ = "" 
				}
		


/*****************************************************************************
 *
 *	target list for SELECT
 *
 *****************************************************************************/

opt_target_list: 
			target_list								
				{ 
					$$ = $1 
				}
			| /* EMPTY */							
				{ 
					$$ = nil
				}

target_list:
			target_el        						
				{ 
					$$=AST.NewList($1)

				}
			|target_list ',' target_el 				
				{ 
					$$ = append($1,$3)
				}


target_el:
            a_expr opt_col_alias     			
            	{ 
            		target:=AST.NewTargetExpr()
            		target.Tuple = $1
            		target.Alias = $2
            		$$ = target
            	}
            |'*' 									
            	{ 
            		target:=AST.NewTargetExpr()
            		star:=AST.NewStartExpr()
            		target.Tuple=star
            		$$ = target
            	}


/*****************************************************************************
 *
 *	table_name,column_name for SELECT
 *
 *****************************************************************************/


qualified_name:
			ColId
				{
					tb:=AST.NewTableExpr()
					tb.Table=$1
					$$=tb
				}			
			|ColId indirection
				{
					tb:=AST.NewTableExpr()
					tb.DB=$1
					tb.Table=$2
					$$=tb
				}


columnref:	
			ColLabel
				{
					col:=AST.NewColumnExpr()
					col.Column= $1
					
					$$ = col
				}
			| ColLabel indirection
				{
					col:=AST.NewColumnExpr()
					col.Table= $1
					col.Column= $2
					
					$$ = col
				}

indirection:
			indirection_el							
				{ 
					//返回字符串
					$$ = $1
				}
			| indirection indirection_el			
				{ 
					//返回字符串
					$$ = $1 + $2
				}


indirection_el:
			'.' ColLabel
				{
					//返回字符串
					$$ = $2
				}
			| '.' '*'
				{
					//返回字符串
					$$ = "*"
				}

/*****************************************************************************
 *
 *	alias clause for SELECT
 *
 *****************************************************************************/

opt_col_alias:
			col_alias 							
				{ 
					$$ = $1
				}
			|/*EMPTY*/    							
				{ 
					$$ = "" 
				}

col_alias:
			AS ColLabel 							
				{	
					//可以为INDET和STRING
					$$ = $2
				}
			|ColLabel   							
				{
					$$ = $1
				}

ColLabel:
			ColId 									
				{ 
					$$ = $1 
				}
			|STRING
				{
					$$ = $1 
				}
				

opt_table_alias:
			table_alias 							
				{ 
					$$ = $1
				}
			|/*EMPTY*/    							
				{ 
					$$ = "" 
				}

table_alias:
			AS ColId 							
				{	
					//可以为INDET和STRING
					$$ = $2
				}
			|ColId   							
				{
					$$ = $1
				}

ColId:
			IDENT
				{
					$$ = $1
				}

/*****************************************************************************
 *
 *	into clause for SELECT
 *
 *****************************************************************************/

 into_clause:
 			INTO into_table_name 							
 				{ 
 					Into:= AST.NewIntoclause()
 					Into.Tbname = $2
 					$$ = Into
 				}
 			|/*EMPTY*/								
 				{ 
 					$$ = nil
 				}

into_table_name:
		  ColId
		  		{
		  			$$ = $1
		  		}

/*****************************************************************************
 *
 *	from_clause for SELECT
 *
 *****************************************************************************/

from_clause:
			FROM from_list							
				{ 
					From:= AST.NewFromclause()
					From.Expr = $2
					$$ = From
				}
			| /*EMPTY*/								
				{ 
					$$ = nil 
				}
		
from_list:
			table_ref								
				{ 
					switch $1.(type){
						case *AST.SimpleTableExpr:
							$$ = AST.NewList($1)
						case *AST.SubTableExpr:
							$$ = AST.NewList($1)
						case *AST.JoinExpr:
							$$ = AST.NewList($1)
					}
				}
			| from_list ',' table_ref				
				{ 
					$$ = append($1,$3) 
				}
		
table_ref:
			relation opt_table_alias       
				{
					SimpleTable := AST.NewSimpleTableExpr()
					SimpleTable.Table = $1.(*AST.TableExpr)
					SimpleTable.Alias = $2
					$$ = SimpleTable
				}
			|select_with_parens table_alias        
				{
					SubTable := AST.NewSubTableExpr()
					SubTable.Stmt = $1
					SubTable.Alias = $2
					$$ = SubTable
				}
			|joined_table
				{
					$$ = $1
				}

relation:
			qualified_name    
				{
					$$ = $1
				}


/*****************************************************************************
 *
 *	where_clause for SELECT
 *
 *****************************************************************************/

where_clause:
			WHERE a_expr							
				{ 
					Where:=AST.NewWhereclause()
					Where.Expr = $2
					$$ = Where
				}
			| /*EMPTY*/								
				{ 
					$$ = nil
				}


/*****************************************************************************
 *
 *	group_clause for SELECT
 *
 *****************************************************************************/

group_clause:
			GROUP BY group_by_list					
				{ 
					Group:=AST.NewGroupclause()
					Group.Expr = $3
					$$ = Group
				}
			| /*EMPTY*/								
				{ 
					$$ = nil
				}


group_by_list:
			group_by_item							
				{ 
					$$ = AST.NewList($1)
				}
			| group_by_list ',' group_by_item		
				{ 
					$$ = append($1,$3)
				}
		

group_by_item:
			a_expr									
				{ 
					$$ = $1 
				}

/*****************************************************************************
 *
 *	having_clause for SELECT
 *
 *****************************************************************************/


having_clause:
			HAVING a_expr							
				{ 
					Having:=AST.NewHavingclause()
					Having.Expr = $2
					$$ = Having
				}
			| /*EMPTY*/								
				{ 
					$$ = nil 
				}



/*****************************************************************************
 *
 *	sort_clause for SELECT
 *
 *****************************************************************************/


opt_sort_clause:
			sort_clause 							
				{	
					$$ = $1
				}
			| /*EMPTY*/								
				{ 
					$$ = nil
				}

sort_clause:
			ORDER BY sortby_list					
				{ 
					Sort:=AST.NewSortclause()
					Sort.Expr = $3
					$$ = Sort
				}


sortby_list:
			sortby									
				{ 
					$$ = AST.NewList($1)
				}
			| sortby_list ',' sortby				
				{ 
					$$ = append($1,$3)
				}

sortby:
			a_expr opt_asc_desc						
				{ 
					SortBy:=AST.NewAscExpr()
					SortBy.Expr = $1
					SortBy.Asc = $2
					if SortBy.Asc ==""{
						SortBy.Not = false
					} else{
						SortBy.Not = true
					}
					$$ = SortBy
				}

	
opt_asc_desc: 
			ASC							
				{ 
					$$ = "ASC" 
				}
			| DESC							
				{
					$$ = "DESC" 
				}
			| /*EMPTY*/						
				{ 
					$$ = "" 
				}
		

/*****************************************************************************
 *
 *	Join_expr for SELECT
 *
 *****************************************************************************/
joined_table:
			'(' joined_table ')'
				{
					$$ = $2
				}
			| table_ref CROSS JOIN table_ref
				{
					Join:=AST.NewJoinExpr(AST.EXPR_JOIN_CROSS)
					Join.Jointype="CROSS"
					Join.Left=$1
					Join.Right=$4

  				    $$ = Join
				}
			| table_ref join_type JOIN table_ref join_qual
				{
        			Join:=AST.NewJoinExpr(AST.EXPR_JOIN)
					switch $2 {
						case "LEFT":
							Join.Ntype=AST.EXPR_JOIN_LEFT
							Join.Jointype="LEFT"
						case "RIGHT":
							Join.Ntype=AST.EXPR_JOIN_RIGHT
							Join.Jointype="RIGHT"
						case "INNER":
							Join.Ntype=AST.EXPR_JOIN_INNER
							Join.Jointype="INNER"
					}
					Join.Joinqual= $5
					Join.Left=$1
					Join.Right=$4

  				    $$ = Join
				}
			| table_ref JOIN table_ref join_qual
				{
					Join:=AST.NewJoinExpr(AST.EXPR_JOIN)
					Join.Jointype= ""
					Join.Joinqual= $4
					Join.Left=$1
					Join.Right=$3

  				    $$ = Join
				}
			| table_ref NATURAL join_type JOIN table_ref
				{
 					Join:=AST.NewJoinExpr(AST.EXPR_JOIN_NATURAL)
 					switch $2 {
						case "LEFT":
							Join.Ntype=AST.EXPR_JOIN_NATURAL_LEFT
							Join.Jointype="NATURAL LEFT"
						case "RIGHT":
							Join.Ntype=AST.EXPR_JOIN_NATURAL_RIGHT
							Join.Jointype="NATURAL RIGHT"
						case "INNER":
							Join.Ntype=AST.EXPR_JOIN_NATURAL_INNER
							Join.Jointype="NATURAL INNER"
					}
					Join.Left=$1
					Join.Right=$5

  				    $$ = Join
				}
			| table_ref NATURAL JOIN table_ref
				{
					Join:=AST.NewJoinExpr(AST.EXPR_JOIN_NATURAL)
					Join.Jointype= "NATURAL"
					Join.Left=$1
					Join.Right=$4

  				    $$ = Join
				}

join_type:
			LEFT 
				{ 
					$$ = "LEFT" 
				}
			|RIGHT 
				{ 
					$$ = "RIGHT" 
				}
			|INNER 
				{ 
					$$ = "INNER" 
				}

join_qual:
			ON a_expr 
				{ 
					On:=AST.NewOnExpr(AST.EXPR_JOIN_ON)
					On.Expr = $2

					$$ = On
				}	
			|USING '(' name_list ')'					
				{ 
					Use:=AST.NewUseExpr(AST.EXPR_JOIN_USING)
					Use.Expr = $3
					
					$$ = Use
				}


name_list:	
			name 	
				{ 	
					$$ = AST.NewList($1)
				}
			| name_list ',' name 
				{ 
					$$ = append($1,$3)
				}

name:
			ColLabel 
				{ 
					col:=AST.NewColumnExpr()
					col.Column=$1
					$$ = col
				}





/*****************************************************************************
 *
 *	lock_clause for SELECT
 *
 *****************************************************************************/

opt_for_locking_clause:
			for_locking_clause						
				{ 
					$$ = $1 
				}
			| /* EMPTY */							
				{ 
					$$ = nil 
				}

for_locking_clause:
			FOR UPDATE									
				{ 	
					Lock:=AST.NewLockingclause()
					Lock.Expr = "FOR UPDATE"
					$$ = Lock
				}
			IN SHARE MODE 						
				{ 	
					Lock:=AST.NewLockingclause()
					Lock.Expr = "IN SHARE MODE"
					$$ = Lock
				}


/*****************************************************************************
 *
 *	Limit for SELECT
 *
 *****************************************************************************/

 opt_select_limit:
			select_limit						
				{ 
					$$ = $1
				}
			| /* EMPTY */						
				{ 
					$$ = nil
				}


select_limit:
			limit_clause offset_clause			
				{ 
					Limit:=AST.NewLimitclause()
					Limit.Limit = $1
					Limit.Offset = $2
					$$ = Limit 
				}
			| offset_clause limit_clause		
				{ 
					Limit:=AST.NewLimitclause()
					Limit.Limit = $2
					Limit.Offset = $1
					$$ = Limit 
				}
			| limit_clause						
				{
					Limit:=AST.NewLimitclause()
					Limit.Limit = $1
					$$ = Limit 
				}
			| offset_clause					
				{
					Limit:=AST.NewLimitclause()
					Limit.Offset = $1
					$$ = Limit 
				}


limit_clause:
			LIMIT select_limit_value
				{ 
					$$ = $2
				}
			| LIMIT select_limit_value ',' select_offset_value
				{

					$$ = append($2,$4)
				}

select_limit_value:
			a_expr									
			{
				$$ = AST.NewList($1)
			}

select_offset_value:
			a_expr									
			{ 
				$$ = $1 
			}


offset_clause:
			OFFSET select_offset_value
			{ 
				$$ = $2 
			}


/*****************************************************************************
 *
 *	expr for SELECT
 *
 *****************************************************************************/
a_expr:
			c_expr 
				{ 
					$$ = $1 
				}
			|'+' a_expr %prec UMINUS
				{
					$$ = $2
				}
			|'-' a_expr %prec UMINUS
				{
					$$ = AST.DoValueExpr($2)
				}
			|a_expr '+' a_expr 		
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_PLUS)
					bool.Op="+"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}
			|a_expr '-' a_expr
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_MINUS)
					bool.Op="-"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}		
			|a_expr '*' a_expr   
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_MULTIPLY)
					bool.Op="*"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr '/' a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_DIVIDE)
					bool.Op="/"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr '%' a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_REMAINDER)
					bool.Op="%"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}    
			|a_expr '^' a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_XOR)
					bool.Op="^"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr '&' a_expr      
				{	
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_AND)
					bool.Op="&"
					bool.Left = $1
					bool.Right= $3
					$$ = bool
				}  
			|a_expr '|' a_expr      
				{	
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_OR)
					bool.Op="|"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}  
			|a_expr '<' a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_L)
					bool.Op="<"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr '>' a_expr    
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_G)
					bool.Op=">"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr '=' a_expr     
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_E)
					bool.Op="="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}  
			|a_expr AA a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_AA)
					bool.Op="&&"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}  
			|a_expr OO a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_OO)
					bool.Op="||"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				} 
			|a_expr SL a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_SL)
					bool.Op="<<"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}  
			|a_expr SR a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_SR)
					bool.Op=">>"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr LEG a_expr      
				{ 
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_LEG)
					bool.Op="<=>"
					bool.Left = $1
					bool.Right= $3
					$$ = bool
				}   
			|a_expr LNE a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_NE)
					bool.Op="!="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr NE a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_NE)
					bool.Op="<>"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr IE a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_IE)
					bool.Op="=="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr LE a_expr     
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_LE)
					bool.Op="<="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr GE a_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_GE)
					bool.Op=">="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr AND a_expr   
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_AND)
					bool.Op="AND"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|a_expr OR a_expr  
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_OR)
					bool.Op="OR"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|NOT a_expr     
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_NOT)
					bool.Op="NOT"
					bool.Left = nil
					bool.Right= $2
					$$ = bool
				}   

			|a_expr LIKE a_expr 
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_LIKE)
					bool.Op="LIKE"
					bool.Left = $1
					bool.Right= $3
					$$ = bool
				}   
			|a_expr NOT LIKE a_expr 
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_NOT_LIKE)
					bool.Op="NOT LIKE"
					bool.Left = $1
					bool.Right= $4
					$$ = bool
				}   
			|a_expr IS NULL 
				{
					Is:=AST.NewIsExpr(AST.EXPR_NULL)
					Is.Not=false
					Is.Left = $1
					$$ = Is
				}   
			|a_expr IS NOT NULL 
				{
					Is:=AST.NewIsExpr(AST.EXPR_NULL)
					Is.Not=true
					Is.Left = $1
					$$ = Is
				}   
			|a_expr IS TRUE 
				{
					Is:=AST.NewIsExpr(AST.EXPR_TRUE)
					Is.Not=false
					Is.Left = $1
					$$ = Is
				}   
			|a_expr NOT IS TRUE 
				{
					Is:=AST.NewIsExpr(AST.EXPR_TRUE)
					Is.Not=true
					Is.Left = $1
					$$ = Is
				}   
			|a_expr IS FALSE
				{
					Is:=AST.NewIsExpr(AST.EXPR_FALSE)
					Is.Not=false
					Is.Left = $1
					$$ = Is
				}   
			|a_expr IS NOT FALSE 
				{
					Is:=AST.NewIsExpr(AST.EXPR_FALSE)
					Is.Not=true
					Is.Left = $1
					$$ = Is
				}   
			|a_expr BETWEEN b_expr AND a_expr %prec BETWEEN
				{
					Between:=AST.NewRangeExpr()
					Between.Expr = $1
					Between.Left = $3
					Between.Right = $5
					$$ = Between
				}
			|a_expr IN in_expr 
				{
					In:=AST.NewInExpr(AST.EXPR_IN_SIMPLE)
					In.Not = false
					In.Left = $1
					In.Right = $3
					//判断InExpr的类型
					//for expr := range In.Right{
					//	switch expr.(type){
					//		case *AST.SubTableExpr:
					//			In.Ntype=AST.EXPR_IN_SUBLINK
					//	}
					//}
					$$ = In
				}
			|a_expr NOT IN in_expr 
				{
					In:=AST.NewInExpr(AST.EXPR_IN_SIMPLE)
					In.Not = true
					In.Left = $1
					In.Right = $4
					//判断InExpr的类型
					//for expr := range In.Right{
					//	switch expr.(type){
					//		case *AST.SubTableExpr:
					//			In.Ntype=AST.EXPR_IN_SUBLINK
					//	}
					//}
					$$ = In
				}
			|a_expr subquery_op subquery_type select_with_parens %prec OP 
				{
					Sub:=AST.NewSubQueryExpr(AST.EXPR_SUBQUERY)
					Sub.Op = $2
					Sub.Subtype = $3
					Sub.Left = $1
					Sub.Right = $4

					switch Sub.Subtype{
						case "SOME":
							Sub.Ntype=AST.EXPR_SUBQUERY_SOME
						case "ANY":
							Sub.Ntype=AST.EXPR_SUBQUERY_ANY
						case "ALL":
							Sub.Ntype=AST.EXPR_SUBQUERY_ALL
					}

					$$ = Sub
				}
			|a_expr subquery_op subquery_type '(' a_expr ')'	%prec OP 
				{
					Sub:=AST.NewSubQueryExpr(AST.EXPR_SUBQUERY)
					Sub.Op = $2
					Sub.Subtype = $3
					Sub.Left = $1
					Sub.Right = $5

					switch Sub.Subtype{
						case "SOME":
							Sub.Ntype=AST.EXPR_SUBQUERY_SOME
						case "ANY":
							Sub.Ntype=AST.EXPR_SUBQUERY_ANY
						case "ALL":
							Sub.Ntype=AST.EXPR_SUBQUERY_ALL
					}

					$$ = Sub
				}

b_expr:
			c_expr
				{ 
					$$ = $1 
				}
			|'+' b_expr %prec UMINUS
				{
					$$ = $2
				}
			|'-' b_expr %prec UMINUS
				{
					$$ = AST.DoValueExpr($2)
				}
			|b_expr '+' b_expr 		
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_PLUS)
					bool.Op="+"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}					
			|b_expr '-' b_expr		
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_MINUS)
					bool.Op="-"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}						
			|b_expr '*' b_expr    	
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_MULTIPLY)
					bool.Op="*"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   				  
			|b_expr '/' b_expr  
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_DIVIDE)
					bool.Op="/"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}     					  
			|b_expr '%' b_expr  
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_REMAINDER)
					bool.Op="%"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}     					    					  
			|b_expr '<' b_expr   
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_L)
					bool.Op="<"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}    					  
			|b_expr '>' b_expr   
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_G)
					bool.Op=">"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}    					  
			|b_expr '=' b_expr 
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_E)
					bool.Op="="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				} 
			|b_expr AA b_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_AA)
					bool.Op="&&"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}  
			|b_expr OO b_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_OO)
					bool.Op="||"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}  
			|b_expr SL b_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_SL)
					bool.Op="<<"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}  
			|b_expr SR b_expr      
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_SR)
					bool.Op=">>"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   
			|b_expr LEG b_expr      
				{ 
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_LEG)
					bool.Op="<=>"
					bool.Left = $1
					bool.Right= $3
					$$ = bool
				}   					  
			|b_expr LNE b_expr   
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_NE)
					bool.Op="!="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   					  
			|b_expr NE b_expr     	
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_NE)
					bool.Op="<>"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}   				 
			|b_expr IE b_expr 
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_IE)
					bool.Op="=="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}     					 
			|b_expr LE b_expr
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_LE)
					bool.Op="<="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}      

			|b_expr GE b_expr  
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_GE)
					bool.Op=">="
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}    					  
			|b_expr OR b_expr  	
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_OR)
					bool.Op="OR"
					bool.Left = $1
					bool.Right= $3
					$$ = bool

				}  						
			|NOT b_expr    		
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_NOT)
					bool.Op="NOT"
					bool.Left = nil
					bool.Right= $2
					$$ = bool
				}   					
			|b_expr LIKE b_expr
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_LIKE)
					bool.Op="LIKE"
					bool.Left = $1
					bool.Right= $3
					$$ = bool
				}    
			|b_expr NOT LIKE b_expr 
				{
					bool:=AST.NewBoolExpr(AST.EXPR_BOOL_NOT_LIKE)
					bool.Op="NOT LIKE"
					bool.Left = $1
					bool.Right= $4
					$$ = bool
				}   
			|b_expr IS NULL 
				{
					Is:=AST.NewIsExpr(AST.EXPR_NULL)
					Is.Not=false
					Is.Left = $1
					$$ = Is
				}   
			|b_expr IS NOT NULL 
				{
					Is:=AST.NewIsExpr(AST.EXPR_NULL)
					Is.Not=true
					Is.Left = $1
					$$ = Is
				}   
			|b_expr IS TRUE 
				{
					Is:=AST.NewIsExpr(AST.EXPR_TRUE)
					Is.Not=false
					Is.Left = $1
					$$ = Is
				}   
			|b_expr NOT IS TRUE 
				{
					Is:=AST.NewIsExpr(AST.EXPR_TRUE)
					Is.Not=true
					Is.Left = $1
					$$ = Is
				}   
			|b_expr IS FALSE 
				{
					Is:=AST.NewIsExpr(AST.EXPR_FALSE)
					Is.Not=false
					Is.Left = $1
					$$ = Is
				}   
			|b_expr NOT IS FALSE 
				{
					Is:=AST.NewIsExpr(AST.EXPR_FALSE)
					Is.Not=true
					Is.Left = $1
					$$ = Is
				}   

c_expr:
			NUMBER		
				{
					s:=AST.NewValueExpr(AST.VALUE_NUMBER)
					s.Val = $1
					$$ = s
				}				
			|'(' a_expr ')' 	
				{
					$$ = $2
				}			
			|EXISTS select_with_parens 
				{
					Sub:=AST.NewSubQueryExpr(AST.EXPR_SUBQUERY_EXISTS)
					Sub.Subtype = "EXISTS"
					Sub.Left = nil
					Sub.Right = $2
					$$ = Sub
				}
			|func_expr opt_collate_set
				{
					$$.(*AST.FuncExpr).Collate = $2
					
				}
			|columnref
				{
					$$ = $1
				}				
			|time_expr	
				{
					s:=AST.NewValueExpr(AST.VALUE_STRING)
					s.Val = $1
					$$ = s
				}
			|case_expr
				{
					$$ = $1
				}	
				

			



/*****************************************************************************
 *
 *	sublink and subquery for SELECT
 *
 *****************************************************************************/

in_expr:
			select_with_parens 
				{
					$$ = AST.NewList($1)
				}
			| '(' expr_list ')'						
				{ 
					$$ = $2
				}


expr_list:	
			a_expr 
				{ 
					$$ = AST.NewList($1)
				}	
			| expr_list ',' a_expr 
				{ 
					$$ = append($1,$3) 
				}	


subquery_op:
			'<'						
				{ 
					$$ = "<" 
				}	
			|'>'					
				{ 
					$$ = ">" 
				}
			|'='					 
				{ 
					$$ = "=" 
				}
			|IE 					
				{ 
					$$ = "==" 
				}
			|GE 					
				{ 
					$$ = ">=" 
				}
			|LE 					
				{
					$$="<="
				}
			|NE 					
				{ 
					$$ = "<>" 
				}


subquery_type:
			SOME 					
				{ 
					$$ = "SOME" 
				}	
			|ANY 					
				{ 
					$$ = "ANY" 
				}	
			|ALL 					
				{ 
					$$ = "ALL" 
				}	


/*****************************************************************************
 *
 *	function for SELECT
 *
 *****************************************************************************/
func_expr:
			func_name '(' ')' 
				{ 
					Func:=AST.NewFuncExpr(AST.EXPR_FUNC)
					Func.Funcname = $1
					Func.Expr = nil

					$$ = Func 
				}
			|func_name '(' func_arg_list ')'  
				{ 
					Func:=AST.NewFuncExpr(AST.EXPR_FUNC)
					Func.Funcname = $1
					Func.Expr = $3

					$$ = Func 
				}

func_name:
			IDENT 
				{
					$$=$1
				}


func_arg_list:  
			func_arg_expr 
				{ 
					$$ = AST.NewList($1)
				}
			| func_arg_list ',' func_arg_expr
				{
					$$ = append($1,$3)
				}


func_arg_expr:  
			opt_distinct a_expr 
				{
					arg:=AST.NewFuncArgExpr()
					arg.Distinct = $1
					arg.Expr= $2
					$$ = arg
				}
			|opt_distinct a_expr AS func_expr opt_character_set
				{
					arg:=AST.NewFuncArgExpr()
					arg.Distinct = $1
					arg.Expr= $2
					arg.TypeAlias=$4
					arg.Character=$5
					$$ = arg
				}
			|opt_distinct a_expr opt_col_alias opt_character_set
				{ 
					arg:=AST.NewFuncArgExpr()
					arg.Distinct = $1
					arg.Expr= $2
					arg.Alias=$3
					arg.Character=$4
					$$ = arg
				}
opt_distinct:
			DISTINCT
				{
					$$ = " DISTINCT "
				}
			|DISTINCTROW
				{
					$$ = " DISTINCTROW "
				}
			|/*EMPTY*/
				{
				    $$ = ""
				}

opt_character_set:
			CHARACTER SET IDENT 
				{	
					$$=$3
				}


opt_collate_set:
			COLLATE IDENT 
				{
					$$=$2
				}
			|/* EMPTY*/ 
				{
					$$=""
				}



/*****************************************************************************
 *
 *	time_expr for SELECT
 *
 *****************************************************************************/

time_expr:
			INTERVAL NUMBER time_type 
				{ 
					$$ = " INTERVAL "+ $2 + $3
				}
time_type:
			DAY                 			{ $$ = "DAY"}
			|DAY_HOUR           			{ $$ = "DAY_HOUR"}
			|DAY_MICROSECOND    			{ $$ = "DAY_MICROSECOND"}
			|DAY_MINUTE         			{ $$ = "DAY_MINUTE"}
			|DAY_SECOND         			{ $$ = "DAY_SECOND"}
			|HOUR               			{ $$ = "HOUR"}
			|HOUR_MICROSECOND   			{ $$ = "HOUR_MICROSECOND"}
			|HOUR_MINUTE        			{ $$ = "HOUR_MINUTE"}
			|HOUR_SECOND        			{ $$ = "HOUR_SECOND"}
			|MICROSECOND        			{ $$ = "MICROSECOND"}
			|MINUTE             			{ $$ = "MINUTE"}
			|MINUTE_MICROSECOND 			{ $$ = "MINUTE_MICROSECOND"}
			|MINUTE_SECOND      			{ $$ = "MINUTE_SECOND"}
			|MONTH              			{ $$ = "MONTH"}
			|QUARTER            			{ $$ = "QUARTER"}
			|SECOND             			{ $$ = "SECOND"}
			|SECOND_MICROSECOND 			{ $$ = "SECOND_MICROSECOND"}
			|WEEK               			{ $$ = "WEEK"}
			|YEAR               			{ $$ = "YEAR"}
			|YEAR_MONTH         			{ $$ = "YEAR_MONTH"}


/*****************************************************************************
 *
 *	case_expr for SELECT
 *
 *****************************************************************************/

case_expr:	
			CASE case_arg when_clause_list case_default END
				{
					ca:=AST.NewCaseExpr(AST.EXPR_CASE)
					ca.Expr = $2
					ca.When = $3
					ca.Else = $4

					$$ = ca
				}
		

when_clause_list:
			when_clause								
				{ 
					$$ = AST.NewList($1)
				}
			| when_clause_list when_clause			
				{ 
					$$ = append($1, $2) 
				}
		

when_clause:
			WHEN a_expr THEN a_expr
				{
					when:=AST.NewWhenExpr()
					when.Left = $2
					when.Right = $4

					$$ = when
				}
		

case_default:
			ELSE a_expr								
				{ 
					el:=AST.NewElseExpr()
					el.Expr = $2
					$$ = el
				}
			| /*EMPTY*/								
				{ 
					$$ = nil 
				}
		

case_arg:	a_expr									
				{ 
					$$ = $1
				}
			| /*EMPTY*/								
				{ 
					$$ = nil 
				}
		