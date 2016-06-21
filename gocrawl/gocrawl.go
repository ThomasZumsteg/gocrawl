package gocrawl

import (
    "bytes"
	"fmt"
    "golang.org/x/crypto/ssh"
    // "time"
)

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

// GetUserAndPassword get the login details for the user
func (dev *Device) GetUserAndPassword() {
	fmt.Printf("Username: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}

// Connect sends a command to the device
func (dev *Device) Connect(user, password, command string) string {
    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{ssh.Password(password)},
    }

    conn, _ := ssh.Dial("tcp", dev.name+":22", config)
    session, _ := conn.NewSession()
    defer session.Close()

    var buffer bytes.Buffer
    session.Stdout = &buffer
    session.Run(command)

    return buffer.String()
}
