package gocrawl

<<<<<<< HEAD
const TestVersion = 1

type device struct {
=======
// TestVersion is the version that the unit tests are run against
const TestVersion = 1

// Device is a network device that sends and receives commands
type Device struct {
>>>>>>> dev
	name string
	ip   string
}

<<<<<<< HEAD
func newConnection(name, ip string) *device {
	return &device{name: name, ip: ip}
=======
// NewDevice creates a new device
func NewDevice(name, ip string) *Device {
	return &Device{name: name, ip: ip}
>>>>>>> dev
}
