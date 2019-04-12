/*
    FileName:   driver.go
    Author:     Zhou Gaochang
    @contact:   zhougaochang@corp.netease.com
    @version:   
    @Time:       6:29 PM
    Description:
    Changelog:
 */

package driver

import "sync"

var (
	driversMu sync.RWMutex
	drivers = make(map[string]Driver)
)

type Driver interface{
}

func Register(name string, driver Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()

	if driver == nil {
		panic("ggql: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("sql: Register called twice for driver " + name)
	}
	drivers[name] = driver
}
