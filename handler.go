/*
    FileName:   handler.go
    Author:     Zhou Gaochang
    @contact:   zhougaochang@corp.netease.com
    @version:   
    @Time:       6:19 PM
    Description:
    Changelog:
 */

package grql

import (
	"github.com/gin-gonic/gin"
	"grql/driver"
)

type Handler struct {
	Driver driver.Driver
	Database string
	Table string
	Model interface{}
	Columns [][2]string
}


func (handler *Handler) HEADER(ctx *gin.Context) {
}

func (handler *Handler) GET(ctx *gin.Context) {
}

func (handler *Handler) POST(ctx *gin.Context) {
}

func (handler *Handler) PUT(ctx *gin.Context) {
}

func (handler *Handler) DELETE(ctx *gin.Context) {
}
