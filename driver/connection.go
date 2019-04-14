/*
    FileName:   connection.go
    Author:     Zhou Gaochang
    @contact:   zhougaochang@corp.netease.com
    @version:   
    @Time:       4:52 PM
    Description:
    Changelog:
 */

package driver

import "database/sql"

type Connection interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
