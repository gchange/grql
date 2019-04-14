/*
    FileName:   grql_test.go
    Author:     Zhou Gaochang
    @contact:   zhougaochang@corp.netease.com
    @version:   
    @Time:       4:48 PM
    Description:
    Changelog:
 */

package grql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"grql/driver"
	_ "grql/driver/mysql"
	"testing"
)

func TestNew(t *testing.T) {
	driverName := "mysql"
	db, err := sqlx.Connect(driverName, "root:123456@tcp(localhost:13306)/mysql")
	if err != nil {
		panic(err)
	}
	d, _ := driver.New(driverName, db)
	d.New(db)
	rows, _ := db.Query("select database();")
	defer rows.Close()
	var dbN string
	for rows.Next() {
		rows.Scan(&dbN)
		fmt.Println(dbN)
	}
	query := "select ? from information_schema.columns"
	rows, err = db.Query(query, "table_name")
	fmt.Println(rows, err)
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println(name)
	}
}
