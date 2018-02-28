package session

import "testing"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

func Test_connect(t *testing.T) {
	user := ""
	password := ""
	host := ""
	port := ""
	database := ""

	Spid := NewSession(user, password, host, port, database)
	db, err := sql.Open(DRIVER, Spid.String())
	if err != nil {
		log.Fatal(err)
	}

	var a int
	var b int

	rows, err := db.Query("select a,b from test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&a, &b)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(a, b)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
