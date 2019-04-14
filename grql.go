/*
   FileName:   grql.go
   Author:     Zhou Gaochang
   @contact:   zhougaochang@corp.netease.com
   @version:
   @Time:       4:24 PM
   Description:
   Changelog:
*/

package grql

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"grql/driver"
)

var (
	grql = New()
)

type GRQL struct {
	engine *gin.Engine
}

func Default() *GRQL {
	return grql
}

func New() *GRQL {
	return &GRQL{
		engine: gin.Default(),
	}
}

func (grql *GRQL) Model(driverName string, conn driver.Connection, database, table string, model interface{},
	prefix, path string) error {
	if database == "" && model == nil {
		return errors.New("grql: model must specify a table name or model")
	}

	dbDriver, err := driver.New(driverName, conn)
	if err != nil {
		return fmt.Errorf("grql: unknown driver %q (forgotten import?)", driverName)
	}
	handler := Handler{
		Driver: dbDriver,
	}
	if model == nil {
		if database == "" {
			database, err = dbDriver.Database()
			if err != nil {
				return err
			}
		}
		_, err := dbDriver.Columns(database, table)
		if err != nil {
			return err
		}
	}

	if path == "" {
		path = "/" + database + "/" + "table"
	}

	grql.engine.HEAD(path, handler.HEADER)
	grql.engine.GET(path, handler.GET)
	grql.engine.POST(path, handler.POST)
	grql.engine.PUT(path, handler.PUT)
	grql.engine.DELETE(path, handler.DELETE)

	onePath := path + "/:id"
	grql.engine.GET()
	return nil
}
