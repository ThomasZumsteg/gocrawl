package gocrawl

import (
    // "bytes"
	"fmt"
    "golang.org/x/crypto/ssh"
	// "log"
    "time"
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
    // Check here http://stackoverflow.com/questions/21126195/talking-to-cisco-equipment-using-go-ssh-library
    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{ssh.Password(password)},
    }

    conn, _ := ssh.Dial("tcp", dev.name+":22", config)
    session, _ := conn.NewSession()
    defer session.Close()

	stdout, _ := session.StdoutPipe()
	stdin, _ := session.StdinPipe()

    modes := ssh.TerminalModes{
        ssh.ECHO:          0,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }

    session.RequestPty("xterm", 0, 200, modes)
    session.Shell()
    time.Sleep(10 * time.Second)

    stdin.Write([]byte("show ver\r"))
    time.Sleep(10 * time.Second)

    buff := make([]byte, 2048)
    n, _ := stdout.Read(buff)

    return string(buff[:n])
}
