package swap

import (
	"github.com/crypto-smoke/arbiter/swap/driver"
	"sort"
	"sync"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]driver.Driver)
)

func Register(name string, driver driver.Driver) {

	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("swap: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("swap: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

// Drivers returns a sorted list of the names of the registered drivers.
func Drivers() []string {
	driversMu.RLock()
	defer driversMu.RUnlock()
	list := make([]string, 0, len(drivers))
	for name := range drivers {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}
