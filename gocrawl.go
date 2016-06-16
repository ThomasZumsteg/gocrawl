package gocrawl

// TestVersion is the version that the unit tests are run against
const TestVersion = 1

// Device is a network device that sends and receives commands
type Device struct {
	name string
	ip   string
}

// NewDevice creates a new device
func NewDevice(name, ip string) *Device {
	return &Device{name: name, ip: ip}
}
