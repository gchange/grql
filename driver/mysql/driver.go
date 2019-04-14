/*
    FileName:   driver.go
    Author:     Zhou Gaochang
    @contact:   zhougaochang@corp.netease.com
    @version:   
    @Time:       4:08 PM
    Description:
    Changelog:
 */

package mysql

import (
	"errors"
	"grql/driver"
)

type MySQLDriver struct {
	conn driver.Connection
}

func (driver *MySQLDriver) New(conn driver.Connection) (driver.Driver, error) {
	return &MySQLDriver{
		conn: conn,
	}, nil
}

func (driver *MySQLDriver) Database() (string, error) {
	rows, err := driver.conn.Query("select database();")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", errors.New("grql: mysql connection have no selected database")
	}
	var database string
	err = rows.Scan(&database)
	if err != nil {
		return "", err
	}
	return database, nil
}

func (driver *MySQLDriver) Columns(database, table string) ([][2]string, error) {
	query := "select column_name, data_type from infomation_schema where table_schema = ? and table_name = ?"
	rows, err := driver.conn.Query(query, database, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columnName, dataType string
	columns := make([][2]string, 0)
	for rows.Next() {
		err = rows.Scan(&columnName, &dataType)
		if err != nil {
			return nil, err
		}
		columns = append(columns, [2]string{columnName, dataType})
	}
	return columns, nil
}

func init() {
	driver.Register("mysql", &MySQLDriver{})
}
