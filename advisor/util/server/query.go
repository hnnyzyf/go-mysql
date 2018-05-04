package server

import "fmt"
import "github.com/hnnyzyf/mysql-go/advisor/util/meta"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

//查询每个表的元信息
func (m *Mysql) QueryTable(session string, ob string) (*meta.Meta, error) {

	var Field string
	var Tp string
	var Null string
	var Key sql.NullString
	var Default sql.NullString
	var Extra sql.NullString

	//插入元数据
	tablemeta := meta.NewMeta()

	//连接数据库
	db, err := sql.Open("mysql", session)
	if err != nil {
		return tablemeta, err
	}

	//查询表结构

	query := fmt.Sprintf("DESC %s ", ob)
	rows, err := db.Query(query)
	if err != nil {
		return tablemeta, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&Field, &Tp, &Null, &Key, &Default, &Extra)
		if err != nil {
			return tablemeta, err
		}
		tablemeta.Add(Field, Tp, Null)
	}

	//返回错误
	err = rows.Err()
	if err != nil {
		return tablemeta, err
	}

	defer db.Close()

	return tablemeta, nil
}

//查询表达式是否为真
func (m *Mysql) QueryTrue0False(session string, ob string, tag int) (bool, error) {

	var Val int

	//连接数据库
	db, err := sql.Open("mysql", session)
	if err != nil {
		return false, err
	}

	//查询表结构

	var query string
	if tag == 0 {
		query = fmt.Sprintf("SELECT %s = 1", ob)
	} else {
		query = fmt.Sprintf("SELECT %s", ob)
	}
	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	//插入元数据

	res := []int{}
	for rows.Next() {
		err := rows.Scan(&Val)
		if err != nil {
			return false, err
		}
		res = append(res, Val)
	}

	//返回错误
	err = rows.Err()
	if err != nil {
		return false, err
	}

	defer db.Close()

	if res[0] == 1 {
		return true, err
	} else {
		return false, err
	}
}

//查询指定目标的统计信息
func (m *Mysql) QuerySelectivity(session string, query string) (int, error) {

	var Val int

	//连接数据库
	db, err := sql.Open("mysql", session)
	if err != nil {
		return 0, err
	}

	rows, err := db.Query(query)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	//插入元数据

	res := []int{}

	for rows.Next() {
		err := rows.Scan(&Val)
		if err != nil {
			return 0, err
		}
		res = append(res, Val)
	}

	//返回错误
	err = rows.Err()
	if err != nil {
		return 0, err
	}

	defer db.Close()

	return res[0], nil
}

//查询指定表的索引信息

func (m *Mysql) QueryIndex(session string, table string) (*meta.Index, error) {

	var Table string
	var Non_unique int
	var Key_name string
	var Seq_in_index int
	var Column_name string
	var Collation string
	var Cardinality int
	var Sub_part sql.NullString
	var Packed sql.NullString
	var Null sql.NullString
	var Index_type string
	var Comment string
	var Index_comment string

	res := meta.NewIndex()

	//连接数据库
	db, err := sql.Open("mysql", session)
	if err != nil {
		return res, err
	}

	query := fmt.Sprintf("SHOW INDEX FROM %s", table)

	//执行查询
	rows, err := db.Query(query)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	//插入元数据

	for rows.Next() {
		err := rows.Scan(
			&Table, &Non_unique, &Key_name, &Seq_in_index,
			&Column_name, &Collation, &Cardinality, &Sub_part,
			&Packed, &Null, &Index_type, &Comment, &Index_comment)
		if err != nil {
			return res, err
		}
		if Key_name == "PRIMARY" {
			res.HasPrimary = true
		}
		res.AddIndex(Key_name, Column_name, Cardinality)
	}

	//返回错误
	err = rows.Err()
	if err != nil {
		return res, err
	}

	defer db.Close()

	return res, nil
}

//查询每个表的元信息
func (m *Mysql) QueryRow(session string, table string) (int, error) {

	var Name string
	var Engine string
	var Version int
	var Row_format string
	var Rows int
	var Avg_row_length int
	var Data_length int
	var Max_data_length int
	var Index_length int
	var Data_free int
	var Auto_increment sql.NullString
	var Create_time sql.NullString
	var Update_time sql.NullString
	var Check_time sql.NullString
	var Collation sql.NullString
	var Checksum sql.NullString
	var Create_options sql.NullString
	var Comment sql.NullString

	//连接数据库
	db, err := sql.Open("mysql", session)
	if err != nil {
		return 1, err
	}

	//查询表结构

	query := fmt.Sprintf("SHOW TABLE STATUS LIKE '%s'", table)
	rows, err := db.Query(query)
	if err != nil {
		return 1, err
	}
	defer rows.Close()

	row := 1
	for rows.Next() {
		err := rows.Scan(
			&Name, &Engine, &Version, &Row_format,
			&Rows, &Avg_row_length, &Data_length,
			&Max_data_length, &Index_length, &Data_free,
			&Auto_increment, &Create_time, &Update_time,
			&Check_time, &Collation, &Checksum,
			&Create_options, &Comment)
		if err != nil {
			return 1, err
		}

		row = Rows
	}

	//返回错误
	err = rows.Err()
	if err != nil {
		return 1, err
	}

	defer db.Close()

	if row == 0 {
		return 1, nil
	} else {
		return row, nil
	}
}

//查询每个表的元信息
func (m *Mysql) QueryCardinality(session string, query string) (int, error) {

	var cardinality int

	//连接数据库
	db, err := sql.Open("mysql", session)
	if err != nil {
		return 0, err
	}

	//查询表结构
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	res := 0
	for rows.Next() {
		err := rows.Scan(&cardinality)
		if err != nil {
			return 0, err
		}
		res = cardinality
	}

	//返回错误
	err = rows.Err()
	if err != nil {
		return 0, err
	}

	defer db.Close()

	return res, nil
}
