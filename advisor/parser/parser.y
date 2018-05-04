//头文件
%{
package parser

import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "strings"
%}


//类型定义
%union {
  	ident       	string 	
  	str 			string
  	val  			string	
  	tag 			int
	boolean			bool  	


  	node          	ast.Node
  	list 			[]ast.ExprNode
  	expr 			ast.ExprNode


}

//token定义
%token <ident>		BuiltinCharacterIdent
%token <ident>		BuiltinFucTimeAddIdent
%token <ident>		BuiltinFucTimeSubIdent
%token <ident>		BuiltinTimeUnitIdent
%token <ident>		IDENT

%token <str>		DOUBLE_QUOTA_STRING
%token <str>		PARAM_MARKER
%token <str>		SINGLE_QUOTA_STRING
%token <str>		STRING

%token <val>     	BIT_NUMBER 
%token <val>      	HEX_NUMBER
%token <val>      	INTEGER 
%token <val> 		FLOAT

%token <ident>		ALL
%token <ident>		ANY
%token <ident>		AS
%token <ident>		ASC
%token <ident>		AVG
%token <ident>		BY
%token <ident>		CAST
%token <ident>		CHARACTER
%token <ident>		COMMENT
%token <ident>		CONVERT
%token <ident>		COUNT
%token <ident>		DESC
%token <ident>		DISTINCT
%token <ident>		DISTINCTROW
%token <ident>		FALSE
%token <ident>		FOR
%token <ident>		FROM
%token <ident>		GROUP
%token <ident>		HAVING
%token <ident>		IF
%token <ident>		INTO
%token <ident>		LIMIT
%token <ident>		LOCK
%token <ident>		MAX
%token <ident>		MIN
%token <ident>		MODE
%token <ident>		NULL
%token <ident>		OFFSET
%token <ident>		ORDER
%token <ident>		POSITION
%token <ident>		QUARTER
%token <ident>		SELECT
%token <ident>		SET
%token <ident>		SHARE
%token <ident>		SOME
%token <ident>		SOUNDS
%token <ident>		SUM
%token <ident>		TO
%token <ident>		TRUE
%token <ident>		UPDATE
%token <ident>		WHERE
%token <ident>		EXISTS
%token <ident>		OUTER
 
%type <tag> 		all_or_distinct
%type <tag> 		asc_or_desc
%type <tag> 		collation_name
%type <tag> 		func_name_time_add
%type <tag> 		func_name_time_sub
%type <tag> 		inner_join
%type <tag> 		opt_distinct
%type <tag> 		subquery_type
%type <tag> 		time_unit

%type <boolean>		is_or_not
%type <boolean>		opt_not
%type <boolean>		true_or_false
%type <boolean>		join_type
%type <boolean>		outer_opt

%type <node>		distinct_clause
%type <node>		for_locking_clause
%type <node>		from_clause
%type <node>		group_clause
%type <node>		having_clause
%type <node>		into_clause
%type <node>		opt_for_locking_clause
%type <node>		opt_select_limit
%type <node>		opt_sort_clause
%type <node>		select_clause
%type <node>		select_limit
%type <node>		select_no_parens
%type <node>		select_stmt
%type <node>		select_with_parens
%type <node>		simple_select
%type <node>		sort_clause
%type <node>		where_clause

%type <val>			alias
%type <val>			alias_name
%type <val>			cast_type
%type <val>			comparison_operator
%type <val>			db_name
%type <val>			double_at_ident
%type <val>			func_name
%type <val>			identifier
%type <val>			identifier_or_star
%type <val>			keyword_as_func
%type <val>			opt_alias
%type <val>			single_at_ident
%type <val>			table_name
     	 
%type <expr> 		agg_expr
%type <expr> 		bit_expr
%type <expr> 		boolean_primary
%type <expr> 		case_arg
%type <expr> 		case_expr
%type <expr> 		cast_expr
%type <expr> 		column_ref
%type <expr> 		else_arg
%type <expr> 		expr
%type <expr> 		func_expr
%type <expr> 		group_by_item
%type <expr> 		join_qual
%type <expr> 		joined_table
%type <expr> 		literal
%type <expr> 		marker
%type <expr> 		offset_expr
%type <expr> 		predicate
%type <expr> 		relation
%type <expr> 		rows_offset
%type <expr> 		simple_expr
%type <expr> 		sortby
%type <expr> 		str_expr
%type <expr> 		table_factor
%type <expr> 		table_ref
%type <expr> 		table_refs
%type <expr> 		target_el
%type <expr> 		time_expr
%type <expr> 		variable
%type <expr> 		when_arg

%type <list>		expr_list
%type <list>		from_list
%type <list>		group_by_list
%type <list>		limit_list
%type <list>		opt_target_list
%type <list>		rows
%type <list>		sortby_list
%type <list>		target_list
%type <list>		when_list
			

//优先级定义
%nonassoc LOW
%right <val>  UNION EXCEPT 
%left <val>  INTERSECT
%left <val>  ','
%left <val>  JOIN STRAIGHT_JOIN NATURAL LEFT RIGHT INNER CROSS USE
%left <val> TABLE_REF_PRIORITY
%left <val>  ON USING
%left <val>  OR OO 
%left <val> XOR
%left <val> AND AA
%left <val>  BETWEEN CASE WHEN THEN ELSE
%left <val>  '=' '<' '>' LIKE REGEXP IN GE LE NE LG IE LNE LEG IS
%left <val>  '|'
%left <val>  '&' 
%left <val>  SL SR
%left <val>  '+' '-'
%left <val>  '*' '/' '%' DIV MOD
%left <val>  '^'
%left <val>  '~'
%right <val> '!' NOT
%nonassoc <val>  OP
%right <val> COLLATE
%nonassoc INTERVAL
//select_with_parens用于消除所有的(select)形式的语句，需优先将将'(' select_with_parens ')'规约为 select_with_parens,而不是reduce为其他符号
//如果出现shift-reduce冲突，强制给予规约低于')'的优先级,保证规则select_with_parens <---'(' select_with_parens ')' 为第一优先级
%nonassoc <val>  UMINUS
%nonassoc <val>  '.'
%nonassoc <val> ')'
%left <val>  END


%start start
%%


start:
			select_stmt
			{
				 yylex.(*Tokener).ParseStmt  = $1
			}

/************************************************************************************
 *
 *  Select Statements
 *
 **********************************************************************************/



select_stmt:
			select_no_parens			%prec UMINUS
			| select_with_parens		%prec UMINUS


select_with_parens:
			'(' select_no_parens ')'				
				{ 
					$$ = $2 
				}
			| '(' select_with_parens ')'
				{ 
					$$ = $2 
				}	

select_no_parens:
			simple_select							
				{ 
					$$ = $1 
				}
			| select_clause sort_clause        		
				{ 
					switch node:=$1.(type){
						case *ast.SimpleSelectStmt:
							node.Sort=$2.(*ast.SortClause)
						case *ast.UnionStmt:
							node.Sort=$2.(*ast.SortClause)
					}

					$$ = $1
				}
			| select_clause opt_sort_clause for_locking_clause opt_select_limit 
				{	
					switch node:=$1.(type){
						case *ast.SimpleSelectStmt:
							node.Sort=$2.(*ast.SortClause)
							node.Lock=$3.(*ast.LockClause)
							node.Limit=$4.(*ast.LimitClause)
						case *ast.UnionStmt:
							node.Sort=$2.(*ast.SortClause)
							node.Lock=$3.(*ast.LockClause)
							node.Limit=$4.(*ast.LimitClause)
					}

					$$ = $1
				}
			| select_clause opt_sort_clause select_limit opt_for_locking_clause  
				{
					switch node:=$1.(type){
						case *ast.SimpleSelectStmt:
							node.Sort=$2.(*ast.SortClause)
							node.Lock=$4.(*ast.LockClause)
							node.Limit=$3.(*ast.LimitClause)
						case *ast.UnionStmt:
							node.Sort=$2.(*ast.SortClause)
							node.Lock=$4.(*ast.LockClause)
							node.Limit=$3.(*ast.LimitClause)
					}

					$$ = $1

					$$ = $1
				}

select_clause:
			simple_select				{ $$ = $1 }
			| select_with_parens		{ $$ = $1 }


simple_select:
			SELECT opt_target_list into_clause from_clause where_clause group_clause having_clause 
				{	
 					$$ = &ast.SimpleSelectStmt{
 						Target:&ast.TargetClause{Target_ref:$2},
 						Into:$3.(*ast.IntoClause),
 						From:$4.(*ast.FromClause),
 						Where:$5.(*ast.WhereClause),
 						Group:$6.(*ast.GroupClause),
 						Having:$7.(*ast.HavingClause),
 					}
				}
			|SELECT distinct_clause target_list into_clause from_clause where_clause group_clause having_clause 
				{	
					$$ = &ast.SimpleSelectStmt{
						Distinct:$2.(*ast.DistinctClause),
						Target:&ast.TargetClause{Target_ref:$3},	
 						Into:$4.(*ast.IntoClause),
 						From:$5.(*ast.FromClause),
 						Where:$6.(*ast.WhereClause),
 						Group:$7.(*ast.GroupClause),
 						Having:$8.(*ast.HavingClause),
					}
				}
			| select_clause UNION all_or_distinct select_clause
				{
					stmt:=&ast.UnionStmt{
						DistinctType:$3,
						Left:$1,
						Right:$4,
					}
					stmt.SetTag(ast.Union)
					$$ = stmt
				}
			| select_clause INTERSECT all_or_distinct select_clause
				{
					stmt:=&ast.UnionStmt{
						DistinctType:$3,
						Left:$1,
						Right:$4,
					}
					stmt.SetTag(ast.Except)
					$$ = stmt
				}
			| select_clause EXCEPT all_or_distinct select_clause
				{
          			stmt:=&ast.UnionStmt{
          				DistinctType:$3,
          				Left:$1,
          				Right:$4,
          			}
					stmt.SetTag(ast.Intersect)
					$$ = stmt
				}



/*****************************************************************************
 *
 *	expr for SELECT
 *
 *****************************************************************************/

expr:
  			expr OR expr %prec OR				
  				{ 
            		expr:=&ast.And0OrExpr{Operator:"OR",Expr:[]ast.ExprNode{}}
            		expr.SetTag(ast.Expr_Or)
            		expr.Combine($1,$3)
            		$$ = expr

        		}
        	| expr OO expr        
          		{ 
          		    expr:=&ast.And0OrExpr{Operator:"OR",Expr:[]ast.ExprNode{}}
          		    expr.SetTag(ast.Expr_Or)
            		expr.Combine($1,$3)
            		$$ = expr
          		}
  			| expr XOR expr				{ $$ = &ast.Expr{Operator:"XOR",Left:$1,Right:$3}}
  			| expr AND expr				
  			  	{
        			expr:=&ast.And0OrExpr{Operator:"AND",Expr:[]ast.ExprNode{}}
        			expr.SetTag(ast.Expr_And)
            		expr.Combine($1,$3)
            		$$ = expr
  				}
  			| expr AA expr 				
  				{
        			expr:=&ast.And0OrExpr{Operator:"AND",Expr:[]ast.ExprNode{}}
        			expr.SetTag(ast.Expr_And)
            		expr.Combine($1,$3)
            		$$ = expr
  				}
  			| NOT expr 					{ $$ = &ast.NotExpr{Expr:$2} }
  			| boolean_primary			{  $$ = $1  }
  			| boolean_primary is_or_not true_or_false 
  				{
  					$$ = &ast.IsTrueExpr{HasNot:$2,IsTrue:$3,Expr:$1}
  				}

boolean_primary:
  			boolean_primary is_or_not NULL
  			  	{ 
  			  		$$ = &ast.IsNullExpr{HasNot:$2,Left:$1}
  				}
  			| boolean_primary LEG predicate
  				{
  					$$ = &ast.Expr{Operator:"=",Left:$1,Right:$3}
  				}
  			| boolean_primary comparison_operator predicate
  				{
  					$$ = &ast.Expr{Operator:$2,Left:$1,Right:$3}

  			  	}
  			| boolean_primary comparison_operator subquery_type select_with_parens
  				{
  					expr := &ast.SubqueryExpr{Operator:$2,Subtype:$3,Left:$1,Select:$4}
  					expr.SetTag(ast.Subquery_Operator)
  					$$ = expr
  				}
  			| predicate
  				{
  					$$ = $1
  				}

 predicate:
  			bit_expr opt_not IN select_with_parens
  			  	{
  			  		expr:= &ast.SubqueryExpr{HasNot:$2,Left:$1,Select:$4}
  			  		expr.SetTag(ast.Subquery_In)
  			  		$$= expr
  			  	}
  			| bit_expr opt_not IN '(' expr_list ')'
  				{	
  					$$ = &ast.InExpr{HasNot:$2,Left:$1,Right:$5}
  				}
  			| bit_expr opt_not BETWEEN bit_expr AND predicate
  				{
  					$$ =&ast.BetweenExpr{HasNot:$2,Expr:$1,Left:$4,Right:$6}
  				}
  			| bit_expr SOUNDS LIKE bit_expr
  				{
  					$$ = &ast.LikeExpr{HasNot:false,Left:$1,Right:$4}
  				}
  			| bit_expr opt_not LIKE simple_expr
  				{
  					$$ = &ast.LikeExpr{HasNot:$2,Left:$1,Right:$4}
  				}
  			| bit_expr opt_not REGEXP bit_expr
  				{
  					$$ = &ast.RegexpExpr{HasNot:$2,Left:$1,Right:$4}
  				}
  			| bit_expr
  				{
  					$$ = $1
  				}


bit_expr:
  			bit_expr '|' bit_expr 		{ $$ = &ast.Expr{Operator:"|",Left:$1,Right:$3}	}
  			| bit_expr '&' bit_expr 	{ $$ = &ast.Expr{Operator:"&",Left:$1,Right:$3}	}
  			| bit_expr SL bit_expr 		{ $$ = &ast.Expr{Operator:"<<",Left:$1,Right:$3}}
  			| bit_expr SR bit_expr 		{ $$ = &ast.Expr{Operator:">>",Left:$1,Right:$3}}
  			| bit_expr '+' bit_expr 	{ $$ = &ast.Expr{Operator:"+",Left:$1,Right:$3}	}
  			| bit_expr '-' bit_expr 	{ $$ = &ast.Expr{Operator:"-",Left:$1,Right:$3}	}
  			| bit_expr '*' bit_expr 	{ $$ = &ast.Expr{Operator:"*",Left:$1,Right:$3}	}
  			| bit_expr '/' bit_expr 	{ $$ = &ast.Expr{Operator:"/",Left:$1,Right:$3}	}
  			| bit_expr DIV bit_expr 	{ $$ = &ast.Expr{Operator:"DIV",Left:$1,Right:$3} }
  			| bit_expr MOD bit_expr 	{ $$ = &ast.Expr{Operator:"MOD",Left:$1,Right:$3} }
  			| bit_expr '%' bit_expr 	{ $$ = &ast.Expr{Operator:"%",Left:$1,Right:$3}	}
  			| bit_expr '^' bit_expr 	{ $$ = &ast.Expr{Operator:"^",Left:$1,Right:$3}	}
  			| bit_expr '+' INTERVAL expr time_unit %prec '+'
  				{
  					$$ = &ast.IntervalExpr{Operator:ast.Op_Add,Left:$1,Right:$4,TimeUnit:$5}
  				}
  			| bit_expr '-' INTERVAL expr time_unit %prec '-'
  				{
  			  		$$ = &ast.IntervalExpr{Operator:ast.Op_Minus,Left:$1,Right:$4,TimeUnit:$5}
  			  	}
  			| simple_expr %prec LOW
  				{
  					$$ = $1
  				}

simple_expr:
  			 literal  					{ $$ = $1 }
  			| column_ref 				{ $$ = $1 }
  			| variable 					{ $$ = $1 }
  			| func_expr 				{ $$ = $1 }
  			| marker 					{ $$ = $1 }
  			| simple_expr COLLATE collation_name %prec OP  
  				{ 
  					$$ = &ast.CollateExpr{Expr:$1,Collate:$3}
  				}
  			| '+' simple_expr	%prec OP
  				{
  					$$ = &ast.UnaryExpr{Operator:"+",Expr:$2}
  				}
  			| '-' simple_expr	%prec OP
  				{
  					$$ = &ast.UnaryExpr{Operator:"-",Expr:$2}	
  				}
  			| '~' simple_expr	%prec OP
  				{
  					$$ = &ast.UnaryExpr{Operator:"~",Expr:$2}
  				}
  			| '!' simple_expr 	%prec OP
  			  	{
  					$$ = &ast.UnaryExpr{Operator:"!",Expr:$2}
  				}
  			| '(' expr ')'
  				{
            		switch node:=$2.(type){
              		case *ast.Expr:
                		node.IsEnclosed=true
                	case *ast.And0OrExpr:
                		node.IsEnclosed=true
            		}
  					$$ = $2
  				}
  			//降低规约优先级
  			| select_with_parens  		%prec UMINUS
  				{
  					$$ = $1.(ast.ExprNode)
  				}
  			| EXISTS select_with_parens
  				{
  					expr:= &ast.SubqueryExpr{Select:$2}
            		expr.SetTag(ast.Subquery_Exists)
            		$$ = expr
  				}
  			| case_expr
  				{
  					$$ = $1
  				}

literal:
		    BIT_NUMBER
			   	{
					Expr := &ast.Integer{Operator:[]string{},Val:$1}
					Expr.SetTag(ast.Bit)
					$$ = Expr
			   	}
		    |HEX_NUMBER
			   	{
			   		Expr := &ast.Integer{Operator:[]string{},Val:$1}
					Expr.SetTag(ast.Hex)
					$$ = Expr
			   }
		    |INTEGER
			   	{
			   		$$ = &ast.Integer{Operator:[]string{},Val:$1}
			   	}
			|FLOAT
				{
					val:=strings.Split($1,".")
					$$ = &ast.Float{
            Operator:[]string{},
						Integer:val[0],
						Decimal:val[1],
					}
				}
			|STRING 					{ $$ =&ast.Marker{Str:$1} }
			|SINGLE_QUOTA_STRING 		
        { 
          str:=&ast.String{Operator:[]string{},Str:$1} 
          str.SetTag(ast.String_single)
          $$ = str
        }
			|DOUBLE_QUOTA_STRING 		
        { 
          str:=&ast.String{Operator:[]string{},Str:$1} 
          str.SetTag(ast.String_double)
          $$ = str
        }
			|true_or_false 				{ $$ =&ast.True{IsTrue:$1} }
      		|NULL 						{ $$ =&ast.Null{} }


identifier:
        	IDENT 		{ $$ = $1 }

identifier_or_star:
          	identifier
          	    {
          	        $$ = $1
          	    }
          	|'*' 
              	{
              	    $$ = "*"
              	}
      
variable:
    		single_at_ident 			{ $$ =&ast.UserVariable{Variable:$1} }
    		|double_at_ident			{ $$ =&ast.SystemVariable{Schema:$1} }
      		|double_at_ident '.' identifier
      			{
      				$$ =&ast.SystemVariable{Schema:$1,Parameter:$3} 
      			}

column_ref:
      		identifier_or_star
      			{
       	 			if $1=="*"{
       	 				$$ = &ast.Star{Column:[]*ast.Column{}}
       	 			}else{
       	 				$$ = &ast.Column{Column:$1}
       	 			}
      			}
      		|identifier '.' identifier_or_star
      			{
      				if $3=="*"{
       	 				$$ = &ast.Star{Table:$1,Column:[]*ast.Column{}}
       	 			}else{
       	 				$$ = &ast.Column{Table:$1,Column:$3}
       	 			}
      			}
      		|identifier '.' identifier '.' identifier_or_star
      			{
      				if $5=="*"{
       	 				$$ = &ast.Star{DB:$1,Table:$3,Column:[]*ast.Column{}}
       	 			}else{
       	 				$$ = &ast.Column{DB:$1,Table:$3,Column:$5}
       	 			}
      			}


func_expr:
			agg_expr
				{
					$$ = $1
				}
      		|time_expr
      		 	{
      		 	  	$$ = $1
      		 	}
          	|cast_expr
            	{
                	$$ = $1
            	}
          	|str_expr
            	{
               		$$ = $1
            	}
      		|func_name '(' ')' 	
      			{ 
      				$$ = &ast.FuncCall{Name:$1} 
      			}
			|func_name '(' expr_list ')'
				{
					$$ = &ast.FuncCall{Name:$1,Arg:$3} 
				}

agg_expr:
			AVG '(' all_or_distinct expr ')'
				{
					expr := &ast.AggregationFuncCall{DistinctType:$3,Arg:[]ast.ExprNode{$4}}
					expr.SetTag(ast.Aggregation_Avg)
  			  		$$ = expr
				}
			|COUNT '(' opt_distinct expr_list ')'
				{
					expr := &ast.AggregationFuncCall{DistinctType:$3,Arg:$4}
					expr.SetTag(ast.Aggregation_Count)
  			  		$$ = expr
				}
			|COUNT '(' ALL expr ')'
				{
					expr := &ast.AggregationFuncCall{DistinctType:ast.AST_ALL,Arg:[]ast.ExprNode{$4}}
					expr.SetTag(ast.Aggregation_Count)
  			  		$$ = expr
				}
			|COUNT '(' expr ')'	
				{
					expr := &ast.AggregationFuncCall{DistinctType:ast.AST_EMPTY,Arg:[]ast.ExprNode{$3}}
					expr.SetTag(ast.Aggregation_Count)
  			  		$$ = expr
				}
			|COUNT '(' '*' ')'
				{
					expr := &ast.AggregationFuncCall{DistinctType:ast.AST_EMPTY,Arg:[]ast.ExprNode{&ast.Star{Column:[]*ast.Column{}}}}
					expr.SetTag(ast.Aggregation_Count)
  			  		$$ = expr
				}
			|MAX '(' all_or_distinct expr ')'
				{
					expr := &ast.AggregationFuncCall{DistinctType:$3,Arg:[]ast.ExprNode{$4}}
					expr.SetTag(ast.Aggregation_Max)
  			  		$$ = expr
				}
			|MIN '(' all_or_distinct expr ')'
				{
					expr := &ast.AggregationFuncCall{DistinctType:$3,Arg:[]ast.ExprNode{$4}}
					expr.SetTag(ast.Aggregation_Min)
  			  		$$ = expr
				}
			|SUM '(' all_or_distinct expr ')'
				{
					expr := &ast.AggregationFuncCall{DistinctType:$3,Arg:[]ast.ExprNode{$4}}
					expr.SetTag(ast.Aggregation_Sum)
  			  		$$ = expr
				}

time_expr:
        	func_name_time_add '(' expr ',' INTERVAL expr time_unit ')'
        	  	{
					Expr:=&ast.CalTimeFuncCall{Interval:&ast.IntervalExpr{Operator:ast.Op_Add,Left:$3,Right:$6,TimeUnit:$7}}
					Expr.SetTag($1)
					$$ = Expr
        	  	}
        	|func_name_time_sub '(' expr ',' INTERVAL expr time_unit ')'
        	  	{
					Expr:=&ast.CalTimeFuncCall{Interval:&ast.IntervalExpr{Operator:ast.Op_Minus,Left:$3,Right:$6,TimeUnit:$7}}
					Expr.SetTag($1)
					$$ = Expr
        	  	}

cast_expr:
          CAST '(' expr AS cast_type ')'
            	{
              	$$ = &ast.CastFuncCall{Expr:$3,CastType:$5}
            	}
          |CONVERT '(' expr ',' cast_type ')'
            	{
               	 	$$ = &ast.CastFuncCall{Expr:$3,CastType:$5}
            	}
          |CONVERT '(' expr USING cast_type ')'
            	{
                	$$ = &ast.CastFuncCall{Expr:$3,CastType:$5}
            	}

str_expr:
          POSITION '(' bit_expr IN expr ')'
            	{
                	$$ = &ast.StringFuncCall{Substr:$3,Str:$5}
            	}

expr_list:
			expr 						{ 	$$ = []ast.ExprNode{$1} }
			|expr_list ',' expr 		{   $$ = append ($1,$3) 	}

marker:
			PARAM_MARKER  				{ $$ = &ast.Marker{Str:$1} }

case_expr:	
			CASE case_arg when_list else_arg END
				{
					$$ = &ast.CaseExpr{Case:$2,When:$3,Else:$4}
				}

case_arg:	
			expr						{ $$ = $1  }
			| /*EMPTY*/					{ $$ = nil }

when_list:
			when_arg					{ $$ = []ast.ExprNode{$1} }
			|when_list when_arg			{ $$ = append($1, $2) }
		
when_arg:
			WHEN expr THEN expr 		{ $$ = &ast.Expr{Left:$2,Right:$4}}
		
else_arg:
			ELSE expr					{ $$ = &ast.Expr{Left:$2}}
			| /*EMPTY*/					{ $$ = nil }
				
opt_target_list:
			target_list					{ $$ = $1  }
			| /* EMPTY */				{ $$ = nil }

target_list:
			target_el        			{ $$ = []ast.ExprNode{$1} }			
			|target_list ',' target_el  { $$ = append($1, $3) }				

target_el:
            expr opt_alias 				
            	{    
              		tb:= &ast.Tuple{Ref:$1,Alias:$2} 
              		switch $1.(type){
              		case *ast.AggregationFuncCall:
              		  tb.SetTag(ast.Tuple_agg)
              		case *ast.FuncCall,*ast.CastFuncCall,*ast.CalTimeFuncCall,*ast.StringFuncCall:
              		  tb.SetTag(ast.Tuple_funcall)
              		case *ast.Column:
              		  tb.SetTag(ast.Tuple_column)
                  case *ast.Star:
                    tb.SetTag(ast.Tuple_star)
                  case *ast.SimpleSelectStmt:
                    tb.SetTag(ast.Tuple_SimpleSubquery)
                  case *ast.UnionStmt:
                    tb.SetTag(ast.Tuple_UnionSubquery)
              		default:
              		  tb.SetTag(ast.Tuple_expr)
              		}
              		$$ = tb
            	} 			

single_at_ident:
            '@' identifier 				{ $$ = $2 }

double_at_ident:
            '@' '@' identifier 			{ $$ = $3 }

func_name:
      		identifier 					{ $$ = $1 }
      		|keyword_as_func   			{ $$ = $1 }

opt_alias:
			alias 						{ $$ = $1 }
			|/*EMPTY*/ 					{ $$ = "" } 

alias:
			AS alias_name 				{ $$ = $2 } 
			|alias_name 				{ $$ = $1 }

alias_name:
			identifier 					{ $$ = $1 }
			|STRING 					{ $$ = $1 }
			|DOUBLE_QUOTA_STRING		{ $$ = $1 }
			|SINGLE_QUOTA_STRING 		{ $$ = $1 }


true_or_false:
			TRUE 						{ $$ = true }
			|FALSE 						{ $$ = false}



subquery_type:
			SOME 						{ $$ = ast.Operator_Some }	
			|ANY 						{ $$ = ast.Operator_All }	
			|ALL 						{ $$ = ast.Operator_Any }	

comparison_operator:
 			'='  						{$$ = "="}
 			| GE 						{$$ = ">="}
 			| '>'						{$$ =">"}
 			| LE 						{$$ = "<="}
 			| '<'						{$$ = "<"}
 			| LG						{$$ = "<>"}
 			| LNE						{$$ = "!="}
				
is_or_not:
			IS 							{ $$ = false }
			|IS NOT 					{ $$ = true }
				

opt_not:
			NOT 						{ $$ = true }
			|/*EMPTY*/ 					{ $$ = false }

time_unit:
			BuiltinTimeUnitIdent 		{ $$ = ast.TimeUnit[$1] }
						

collation_name:
      		BuiltinCharacterIdent   	{ $$ = ast.Character[$1] }
			

keyword_as_func:
     		LEFT                    	{ $$ = "LEFT" }
     		|IN                     	{ $$ = "IN" }
     		|IS                     	{ $$ = "IS" }

cast_type:
      		identifier               	{ $$ = $1 }
	
func_name_time_add:	
      		BuiltinFucTimeAddIdent   	{ $$ = ast.FuncTime[$1] }   
	
func_name_time_sub:                  	     
     		BuiltinFucTimeSubIdent   	{ $$ = ast.FuncTime[$1] }        


/*****************************************************************************
 *
 *	distinct clause for SELECT
 *
 *****************************************************************************/

distinct_clause:
			opt_distinct
				{
					clause:=&ast.DistinctClause{}
					clause.SetTag($1)
					$$ = clause
				}

all_or_distinct:
			ALL							{ $$ = ast.AST_ALL}
			|opt_distinct 				{ $$ = $1}
			| /*EMPTY*/					{ $$ = ast.AST_EMPTY }			

opt_distinct:
			DISTINCT 					{ $$ = ast.AST_DISTINCT }
			| DISTINCTROW 				{ $$ = ast.AST_DISTINCTROW }


/*****************************************************************************
 *
 *	into clause for SELECT
 *
 *****************************************************************************/

 into_clause:
 			INTO table_name 			{ $$ = &ast.IntoClause{Relation:&ast.Relation{Table:$2}} }			
 			|/*EMPTY*/					
 				{ 
 					clause:=&ast.IntoClause{}
 					clause.SetTag(ast.AST_EMPTY)
 					$$ = clause
 				}		

table_name:
		  	identifier 					{ $$ = $1 }
			|STRING 					{ $$ = $1 }
			|DOUBLE_QUOTA_STRING   		{ $$ = $1 }
				
		  	
/*****************************************************************************
 *
 *	from_clause for SELECT
 *
 *****************************************************************************/

from_clause:
			FROM from_list				{ $$ = &ast.FromClause{Table_ref:$2}}		
			| /*EMPTY*/					
			 	{ 
 					clause:=&ast.FromClause{}
 					clause.SetTag(ast.AST_EMPTY)
 					$$ = clause
 				}				

		
from_list: 
			table_ref					{ $$ = []ast.ExprNode{$1}}				
			| from_list ',' table_ref	{ $$ = append($1,$3) }		

table_ref:
			table_factor 				{ $$ = $1}
			|joined_table 				{ $$ = $1}

table_factor:
			relation opt_alias       
				{
          			$1.(*ast.Relation).Alias=$2
          			$$ = $1
				}
			|select_with_parens alias        
				{
				  	tb:= &ast.Relation{Select:$1,Alias:$2}
  					tb.SetTag(ast.Relation_Subquery)
  					$$ = tb
				}
			|'(' table_refs ')'
				{
          switch node:= $2.(type){
            case *ast.Join:
              node.IsEnclosed=true
          }
					$$ = $2
				}

relation:
			table_name    				
        		{ 
        		    tb:= &ast.Relation{Table:$1}
        		    tb.SetTag(ast.Relation_Table)
        		    $$ = tb
        		}
			|db_name '.' table_name 	
        		{
        		  tb:=&ast.Relation{DB:$1,Table:$3}
        		  tb.SetTag(ast.Relation_Table)
        		  $$ = tb
        		}

db_name:
			identifier 					{ $$ = $1}
			|STRING 					{ $$ = $1}

table_refs:
			table_ref 					{ $$ = $1} 
			|table_refs ',' table_ref 	
				{
					rel := &ast.Join{Left:$1,Right:$3}
					rel.SetTag(ast.Cross_join)
					$$ = rel
				}



joined_table:
			//为了保证出现join_qual则一定使用第二条规则进行规约，需要给予第一条规则最低优先级，消除规约冲突
			table_ref inner_join table_ref %prec TABLE_REF_PRIORITY
				{
					rel := &ast.Join{Left:$1,Right:$3}
					rel.SetTag(ast.Cross_join)
					$$ = rel
				}
			|table_ref inner_join table_ref join_qual
				{
					rel := &ast.Join{Left:$1,Right:$3,Condition:$4}
					rel.SetTag(ast.Inner_join)
					$$ = rel
				}
			|table_ref STRAIGHT_JOIN table_factor
				{
					rel := &ast.Join{Left:$1,Right:$3}
					rel.SetTag(ast.Straight_join)
					$$ = rel
				}
			|table_ref STRAIGHT_JOIN table_factor ON expr
				{
					rel := &ast.Join{Left:$1,Right:$3,Condition:$5}
					rel.SetTag(ast.Straight_join)
					$$ = rel
				}
			|table_ref join_type outer_opt JOIN table_ref join_qual
				{
					rel := &ast.Join{HasOuter:$3,Left:$1,Right:$5,Condition:$6}
					if $2{
						rel.SetTag(ast.Left_join)
					}else{
						rel.SetTag(ast.Right_join)
					}
					$$ = rel
				}
			|table_ref NATURAL join_type outer_opt JOIN table_factor
				{
					rel := &ast.Join{HasNatural:true,HasOuter:$4,Left:$1,Right:$6}
					if $3{
						rel.SetTag(ast.Left_join)
					}else{
						rel.SetTag(ast.Right_join)
					}
					$$ = rel
				}

inner_join:
			JOIN 						    { $$ = 1 }
			|INNER JOIN 				{ $$ = 2}
			|CROSS JOIN 				{ $$ = 3 }


join_type:
			LEFT 						{ $$ = true }
			|RIGHT  					{ $$ = false }

outer_opt:
			OUTER  						{ $$ = true }
			|/*EMPTY*/ 					{ $$ = false }

join_qual:
			ON expr  					{ $$ = $2 }
			|USING '(' expr_list ')'	{ $$ = &ast.Using{Column:$3} }		
				

/*****************************************************************************
 *
 *	where_clause for SELECT
 *
 *****************************************************************************/

where_clause:
			WHERE expr					
				{ 
					clause:=&ast.WhereClause{} 
					switch $2.(type){
					case *ast.And0OrExpr:
						clause.Where=$2
					default:
						clause.Where=&ast.And0OrExpr{Operator:"And",Expr:[]ast.ExprNode{$2}}
						clause.Where.SetTag(ast.Expr_And)
					}
					$$ = clause
					
				}
			| /*EMPTY*/								 	
				{ 
 					clause:=&ast.WhereClause{}
 					clause.SetTag(ast.AST_EMPTY)
 					$$ = clause
 				}					

/*****************************************************************************
 *
 *	group_clause for SELECT
 *
 *****************************************************************************/

group_clause:
			GROUP BY group_by_list		{ $$ = &ast.GroupClause{Target_ref:$3} }
			| /*EMPTY*/					
				{ 
 					clause:=&ast.GroupClause{}
 					clause.SetTag(ast.AST_EMPTY)
 					$$ = clause
 				}			

group_by_list:
			group_by_item							
				{ 
					$$ = []ast.ExprNode{$1}
				}
			|group_by_list ',' group_by_item		
				{ 
					$$ = append($1,$3)
				}
		

group_by_item:
			expr						{ $$ = $1 }				

/*****************************************************************************
 *
 *	having_clause for SELECT
 *
 *****************************************************************************/

having_clause:
			HAVING expr					{ $$ = &ast.HavingClause{Having:$2} }
			| /*EMPTY*/					
				{ 
 					clause:=&ast.HavingClause{}
 					clause.SetTag(ast.AST_EMPTY)
 					$$ = clause
 				}				

/*****************************************************************************
 *
 *	sort_clause for SELECT
 *
 *****************************************************************************/


opt_sort_clause:
			sort_clause 				{ $$ = $1 }				
			| /*EMPTY*/					
				{ 
 					clause:=&ast.SortClause{}
 					clause.SetTag(ast.AST_EMPTY)
 					$$ = clause
 				}									

sort_clause:
			ORDER BY sortby_list		{ $$ = &ast.SortClause{Target_ref:$3} }			

sortby_list:
			sortby						{ $$ = []ast.ExprNode{$1} }		
			|sortby_list ',' sortby		{ $$ = append($1,$3) }

sortby:
			expr asc_or_desc			{ $$ = &ast.Sortby{Expr:$1,AscType:$2} }			

asc_or_desc:
			ASC							{ $$ = ast.Sort_Asc }
			|DESC						{ $$ = ast.Sort_Desc }		
			|/*EMPTY*/					{ $$ = ast.Sort_Empty }		

/*****************************************************************************
 *
 *	lock_clause for SELECT
 *
 *****************************************************************************/

opt_for_locking_clause:
			for_locking_clause			{ $$ = $1 }
			| /* EMPTY */				
				{ 
 					clause:=&ast.LockClause{}
 					clause.SetTag(ast.AST_EMPTY)
 					$$ = clause
 				}	

for_locking_clause:
			FOR UPDATE									
				{ 	
					lock:= &ast.LockClause{}
					lock.SetTag(ast.Lock_update)
					$$ = lock
				}
			|LOCK IN SHARE MODE	%prec UMINUS
				{ 	
					lock:= &ast.LockClause{}
					lock.SetTag(ast.Lock_share)
					$$ = lock
				}


/*****************************************************************************
 *
 *	Limit for SELECT
 *
 *****************************************************************************/

 opt_select_limit:
			select_limit				{ $$ = $1 }
			| /* EMPTY */				
				{ 
 					clause:=&ast.LimitClause{}
 					clause.SetTag(ast.AST_EMPTY)
 					$$ = clause
 				}		

select_limit:
			limit_list offset_expr		{ $$ = &ast.LimitClause{Limit:$1,Offset:$2} }
			| offset_expr limit_list	{ $$ = &ast.LimitClause{Limit:$2,Offset:$1} }
			| limit_list				{ $$ = &ast.LimitClause{Limit:$1} }		
			| offset_expr				{ $$ = &ast.LimitClause{Offset:$1} }	



limit_list:
			LIMIT rows 					
				{ 
					$$ = $2 
				}
			|LIMIT rows ',' rows_offset 
				{
					$$ = append($2,$4)
				}

rows:
			expr						{ $$ = []ast.ExprNode{$1} }						

offset_expr:
			OFFSET rows_offset 			{ $$ = $2 }			

rows_offset:
			expr						{ $$ = $1 }	