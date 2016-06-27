package gocrawl

import (
)

// TestVersion is the version that the unit tests are run against
const Version = 2

type device struct {
    Hostname string
    stdin chan string
    stdout chan string
}

func NewDevice(hostname string) device {
    return device{
        Hostname : hostname,
        stdin : nil,
        stdout : nil,
    }
}

func (dev device) Connect(user, pass string) error {
    return nil
}

func (dev device) Send(command string) (string, error) {
    return "response", nil
}
