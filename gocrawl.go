package gocrawl

const TestVersion = 1

type device struct {
	name string
	ip   string
}

func newConnection(name, ip string) *device {
	return &device{name: name, ip: ip}
}
