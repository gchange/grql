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

import (
	"fmt"
	"sync"
)

var (
	driversMu sync.RWMutex
	drivers = make(map[string]Driver)
)

type Driver interface{
	New(Connection) (Driver, error)
	Database() (string, error)
	Columns(string, string) ([][2]string, error)
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

func New(name string, conn Connection) (Driver, error) {
	if conn == nil {
		return nil, fmt.Errorf("grql: cannot create driver %q with nil database connection", name)
	}

	driversMu.RLock()
	driver, ok := drivers[name]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("grql: unknown driver %q (forgotten import?)", name)
	}
	return driver.New(conn)
}
