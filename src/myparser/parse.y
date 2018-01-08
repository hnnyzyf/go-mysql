//头文件
%{
package myparser

func SetParseTree(yylex interface{}, stmt string) {
  yylex.(*Tokener).AST = stmt
}
%}



//类型定义
%union {
  keyword       string 		
  val		    string
  node          string
  list          string
}


//token定义
%token <val> IDENT STRING NUMBER 
%token <keyword> AS ASC ALL ANY AVG
	BY 
	COMMENT COUNT
    DISTINCT DESC
	EXISTS
	FROM FOR FALSE
	GROUP
	HAVING
	INTO IF 
	LIMIT LOCK 
	NULL
	ORDER OFFSET 
	SELECT SOME SHARE SUM
	TO TRUE
	UPDATE USING
	WHERE
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
%left <keyword> '=' '<' '>' IS LIKE IN GE LE NE IE LNE
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

%type<node> SelectStmt select_no_parens select_with_parens simple_select select_clause

%type<node> opt_target_list target_list target_el

%type<node> opt_alias_clause alias_clause ColLabel

%type<node> a_expr b_expr c_expr in_expr expr_list 

%type<node> func_expr func_name func_arg_list func_arg_expr 

%type<node> unreserved_keyword

%type<node> columnref indirection indirection_el

%type<node> subquery_op subquery_type

%type<node> distinct_clause all_or_distinct

%type<node> into_clause

%type<node> from_clause from_list table_ref relation_expr

%type<node> joined_table join_type join_qual name_list name

%type<node> where_clause

%type<node> group_clause group_by_list group_by_item

%type<node> having_clause

%type<node> opt_sort_clause sort_clause  sortby_list sortby  opt_asc_desc

%type<node> opt_for_locking_clause for_locking_clause

%type<node> opt_select_limit select_limit limit_clause offset_clause select_limit_value select_offset_value
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
			| select_clause sort_clause        		{ $$ = $1 +$2 }
			| select_clause opt_sort_clause for_locking_clause opt_select_limit 
				{
					$$ = $1 +$2+$3+$4
				}
			| select_clause opt_sort_clause select_limit opt_for_locking_clause  
				{
					$$ = $1 +$2+$3+$4
				}

select_clause:
			simple_select							{ $$ = $1 }
			| select_with_parens					{ $$ = $1 }


simple_select:
			SELECT opt_target_list into_clause from_clause where_clause group_clause having_clause 
				{
					$$ = "SELECT "+$2+$3+$4+$5+$6
				}
			|SELECT distinct_clause target_list into_clause from_clause where_clause group_clause having_clause 
				{
					$$ ="SELECT " +$2+$3+$4+$5+$6+$7
				}
			| select_clause UNION all_or_distinct select_clause
				{
					$$ = $1 +" UNION " +$3+$4
				}
			| select_clause INTERSECT all_or_distinct select_clause
				{
					$$ = $1 +" INTERSECT " +$3+$4
				}
			| select_clause EXCEPT all_or_distinct select_clause
				{
					$$ = $1 +" EXCEPT " +$3+$4
				}

/*****************************************************************************
 *
 *	distinct clause for SELECT
 *
 *****************************************************************************/

distinct_clause:
			DISTINCT 			{$$=" DISTINCT "}


all_or_distinct:
			ALL										{ $$ = " ALL " }
			| DISTINCT								{ $$ = " DISTINCT " }
			| /*EMPTY*/								{ $$ = "" }
		


/*****************************************************************************
 *
 *	target list for SELECT
 *
 *****************************************************************************/

opt_target_list: 
			target_list								{ $$ = $1 }
			| /* EMPTY */							{ $$ = "" }

target_list:
			target_el        						{ $$ = $1 }
			|target_list ',' target_el 				{ $$ = $$+","+$3 }


target_el:
            a_expr opt_alias_clause      			{ $$ = $1+" "+$2}
            |'*' 									{ $$ = "*" }



/*****************************************************************************
 *
 *	alias clause for SELECT
 *
 *****************************************************************************/

opt_alias_clause:
			alias_clause 							{ $$ = $1}
			|/*EMPTY*/    							{ $$ = "" }

alias_clause:
			AS ColLabel 							{$$ = " AS "+$2}
			|ColLabel   							{$$ = " "+$1}

ColLabel:
			IDENT 									{ $$ = $1 }


/*****************************************************************************
 *
 *	into clause for SELECT
 *
 *****************************************************************************/

 into_clause:
 			INTO ColLabel 							{ $$ = " INTO "+$2}
 			|/*EMPTY*/								{ $$ = ""}

/*****************************************************************************
 *
 *	from_clause for SELECT
 *
 *****************************************************************************/

from_clause:
			FROM from_list							{ $$ = " FROM "+$2 }
			| /*EMPTY*/								{ $$ = "" }
		
from_list:
			table_ref								{ $$ = $1 }
			| from_list ',' table_ref				{ $$ = $$+","+$3 }
		
table_ref:
			relation_expr opt_alias_clause          {$$ = $1 +" "+$2}
			|select_with_parens alias_clause        {$$ = "("+$1+") " +$2}
			|joined_table

relation_expr:
	columnref    {$$ = $1}

/*****************************************************************************
 *
 *	where_clause for SELECT
 *
 *****************************************************************************/

where_clause:
			WHERE a_expr							{ $$ = " WHERE "+$2}
			| /*EMPTY*/								{ $$ = "" }


/*****************************************************************************
 *
 *	group_clause for SELECT
 *
 *****************************************************************************/

group_clause:
			GROUP BY group_by_list					{ $$ =" GROUP BY "+$3 }
			| /*EMPTY*/								{ $$ = "" }


group_by_list:
			group_by_item							{ $$ = $1 }
			| group_by_list ',' group_by_item		{ $$ = $$+","+$3 }
		

group_by_item:
			a_expr									{ $$ = $1 }

/*****************************************************************************
 *
 *	having_clause for SELECT
 *
 *****************************************************************************/


having_clause:
			HAVING a_expr							{ $$ = "  HAVING "+$2 }
			| /*EMPTY*/								{ $$ = "" }

/*****************************************************************************
 *
 *	sort_clause for SELECT
 *
 *****************************************************************************/


opt_sort_clause:
			sort_clause 							{$$ = $1}
			| /*EMPTY*/								{ $$ = "" }

sort_clause:
			ORDER BY sortby_list					{ $$ = " ORDER BY "+$3 }


sortby_list:
			sortby									{ $$ = $1 }
			| sortby_list ',' sortby				{ $$ = $$ +","+$3 }

sortby:
			a_expr opt_asc_desc						{ $$ = $1 + $2 }

	
opt_asc_desc: ASC							{ $$ = " ASC " }
			| DESC							{ $$ = " DESC " }
			| /*EMPTY*/						{ $$ = "" }
		

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
  				    $$ = $1 +" CROSS JOIN " +$2
				}
			| table_ref join_type JOIN table_ref join_qual
				{
        			 $$ = $1 +" "+$2 +" JOIN " +$4 +" "+$5
				}
			| table_ref JOIN table_ref join_qual
				{
					$$ = $1 +" JOIN " +$3 +" "+$4
				}
			| table_ref NATURAL join_type JOIN table_ref
				{
 					$$ = $1 +" NATURAL "+$3+" JOIN " +$4
				}
			| table_ref NATURAL JOIN table_ref
				{
					$$ = $1 +" NATURAL JOIN " +$4
				}

join_type:
	LEFT { $$ = " LEFT " }
	|RIGHT { $$ = " RIGHT " }
	|INNER { $$ = " INNER " }

join_qual:
	ON a_expr { $$ = " ON "+$2}	
	|USING '(' name_list ')'					{ $$ = " USING("+$3+")" }


name_list:	
	name 	{ $$ = $1 }
	| name_list ',' name { $$ = $$+","+$3 }

name:
	ColLabel { $$ = $1}





/*****************************************************************************
 *
 *	lock_clause for SELECT
 *
 *****************************************************************************/

opt_for_locking_clause:
			for_locking_clause						{ $$ = $1 }
			| /* EMPTY */							{ $$ = "" }

for_locking_clause:
			FOR UPDATE									{ $$ = " FOR UPDATE "  }
			FOR SHARE 									{ $$ = " FOR SHARE "  }


/*****************************************************************************
 *
 *	Limit for SELECT
 *
 *****************************************************************************/

 opt_select_limit:
			select_limit						{ $$ = $1}
			| /* EMPTY */						{ $$ = ""}


select_limit:
			limit_clause offset_clause			{ $$ = $1 + $2 }
			| offset_clause limit_clause		{ $$ = $1 + $2 }
			| limit_clause						{ $$ = $1 }
			| offset_clause						{ $$ = $1 }


limit_clause:
			LIMIT select_limit_value
				{ $$ = " LIMIT "+$2 }
			| LIMIT select_limit_value ',' select_offset_value
				{
					$$ = " LIMIT "+$2+","+$4
				}

select_limit_value:
			a_expr									{ $$ = $1 }

select_offset_value:
			a_expr									{ $$ = $1 }


offset_clause:
			OFFSET select_offset_value
			{ $$ = " OFFSET "+$2 }


/*****************************************************************************
 *
 *	expr for SELECT
 *
 *****************************************************************************/
a_expr:
	c_expr { $$ = $1 }
	|a_expr '+' a_expr 		{ $$ = $1+ " + "+$3 }
	|a_expr '-' a_expr		{ $$ = $1 +" - "+$3 }
	|a_expr '*' a_expr      { $$ = $1 +" * "+$3 }
	|a_expr '/' a_expr      { $$ = $1 +" / "+$3 }
	|a_expr '%' a_expr      { $$ = $1 +" % "+$3 }
	|a_expr '^' a_expr      { $$ = $1 +" ^ "+$3 }
	|a_expr '<' a_expr      { $$ = $1 +" < "+$3 }
	|a_expr '>' a_expr      { $$ = $1 +" > "+$3 }
	|a_expr '=' a_expr      { $$ = $1 +" = "+$3 }
	|a_expr LNE a_expr      { $$ = $1 +" != "+$3 }
	|a_expr NE a_expr      { $$ = $1 +" <> "+$3 }
	|a_expr IE a_expr      { $$ = $1 +" == "+$3 }
	|a_expr LE a_expr      { $$ = $1 +" <= "+$3 }
	|a_expr GE a_expr      { $$ = $1 +" >= "+$3 }
	|a_expr AND a_expr     { $$ = $1 +" AND "+$3 }
	|a_expr OR a_expr  { $$ = $1 +" OR "+$3 }
	|NOT a_expr     { $$ = " NOT "+$2 }
	|a_expr LIKE a_expr { $$ = $1 +$2+" LIKE "+$3 }
	|a_expr NOT LIKE a_expr { $$ = $1 +$2+" NOT LIKE "+$3 }
	|a_expr IS NULL { $$ = $1 + " IS NULL " }
	|a_expr IS NOT NULL { $$ = $1 + " IS NOT NULL " }
	|a_expr IS TRUE { $$ = $1 +" IS TRUE " }
	|a_expr NOT IS TRUE { $$ = $1 +"  NOT IS TRUE "}
	|a_expr IS FALSE { $$ = $1 + " IS FALSE "}
	|a_expr NOT IS FALSE { $$ = $1 +" NOT IS FALSE " }
	|a_expr BETWEEN b_expr AND a_expr {$$ = $1 +" BETWEEN "+$3+" AND "+$5 } %prec BETWEEN
	|a_expr IN in_expr {$$ = $1 +" IN "+$3 }
	|a_expr NOT IN in_expr {$$ = $1 +" NOT IN "+$3 }
	|a_expr subquery_op subquery_type select_with_parens %prec OP {$$ = $1 +$2+$3+"("+$4+")"}
	|a_expr subquery_op subquery_type '(' a_expr ')'	%prec OP {$$ = $1 +$2+$3+"("+$5+")"}

b_expr:
	c_expr
	|b_expr '+' b_expr 							{ $$ = $1+ " + "+$3 }
	|b_expr '-' b_expr							{ $$ = $1 +" - "+$3 }
	|b_expr '*' b_expr    					  { $$ = $1 +" * "+$3 }
	|b_expr '/' b_expr    					  { $$ = $1 +" / "+$3 }
	|b_expr '%' b_expr    					  { $$ = $1 +" % "+$3 }
	|b_expr '^' b_expr    					  { $$ = $1 +" ^ "+$3 }
	|b_expr '<' b_expr    					  { $$ = $1 +" < "+$3 }
	|b_expr '>' b_expr    					  { $$ = $1 +" > "+$3 }
	|b_expr '=' b_expr    					  { $$ = $1 +" = "+$3 }
	|b_expr LNE b_expr    					  { $$ = $1 +" != "+$3 }
	|b_expr NE b_expr     					 { $$ = $1 +" <> "+$3 }
	|b_expr IE b_expr     					 { $$ = $1 +" == "+$3 }
	|b_expr LE b_expr     					 { $$ = $1 +" <= "+$3 }
	|b_expr GE b_expr    					  { $$ = $1 +" >= "+$3 }
	|b_expr OR b_expr  							{ $$ = $1 +" OR "+$3 }
	|NOT b_expr    							 { $$ = " NOT "+$2 }
	|b_expr LIKE b_expr { $$ = $1 +$2+" LIKE "+$3 }
	|b_expr NOT LIKE b_expr { $$ = $1 +$2+" NOT LIKE "+$3 }
	|b_expr IS NULL { $$ = $1 + " IS NULL " }
	|b_expr IS NOT NULL { $$ = $1 + " IS NOT NULL " }
	|b_expr IS TRUE { $$ = $1 +" IS TRUE " }
	|b_expr NOT IS TRUE { $$ = $1 +"  NOT IS TRUE "}
	|b_expr IS FALSE { $$ = $1 + " IS FALSE "}
	|b_expr NOT IS FALSE { $$ = $1 +" NOT IS FALSE " }

c_expr:
	STRING 					{ $$ = $1 }
	|NUMBER						{ $$ = $1 }	
	|'(' a_expr ')' 			{ $$ = "( "+$2+" )" }	
	|EXISTS select_with_parens {$$ = " EXISTS ("+$2+")"}
	|func_expr {$$=$1}
	|columnref								{ $$ = $1 }



/*****************************************************************************
 *
 *	table_name,column_name for SELECT
 *
 *****************************************************************************/

columnref:	ColLabel
				{
					$$=$1
				}
			| ColLabel indirection
				{
					$$=$1+$2
				}

indirection:
			indirection_el							{ $$ = $1; }
			| indirection indirection_el			{ $$ = $1+$$ }


indirection_el:
			'.' ColLabel
				{
					$$ = "."+$2
				}
			| '.' '*'
				{
					$$ = ".*"
				}
/*****************************************************************************
 *
 *	sublink and subquery for SELECT
 *
 *****************************************************************************/

in_expr:
	select_with_parens {$$=$1}
	| '(' expr_list ')'						{ $$ = "("+$2+")" }


expr_list:	
	a_expr { $$ = $1 }	
	| expr_list ',' a_expr { $$ = $$+","+$3 }	


subquery_op:
		'<'						{ $$ = "<" }	
		|'>'					{ $$ = ">" }
		|'='					 { $$ = "=" }
		|IE 					{ $$ = "==" }
		|GE 					{ $$ = ">=" }
		|LE 					{$$="<="}
		|NE 					{ $$ = "<>" }


subquery_type:
		SOME 					{ $$ = " SOME " }	
		|ANY 					{ $$ = " ANY " }	
		|ALL 					{ $$ = " ALL " }	


/*****************************************************************************
 *
 *	function for SELECT
 *
 *****************************************************************************/
func_expr:
		func_name '(' ')' { $$ = $1+"()" }
		|func_name '(' func_arg_list ')'  { $$ = $1+"("+$3+")" }

func_name:
		IDENT {$$=$1}
		|unreserved_keyword {$$=$1}


func_arg_list:  
			func_arg_expr { $$ = $1}
			| func_arg_list ',' func_arg_expr
				{
					$$ = $1+","+$3
				}

func_arg_expr:  a_expr
				{
					$$ = $1;
				}

unreserved_keyword:
		SUM   { $$ = " SUM "}
		|AVG  { $$ = " AVG "}
		|COUNT  { $$ = " COUNT "}