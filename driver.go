package mediadevices

import (
	"github.com/kmtekvn/mediadevices/pkg/driver"
)

// RegisterDriverAdapter allows user space level of driver registration
func RegisterDriverAdapter(a driver.Adapter, info driver.Info) error {
	return driver.GetManager().Register(a, info)
}
