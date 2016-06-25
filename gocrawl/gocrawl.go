package gocrawl

import (
	// "fmt"
    // "golang.org/x/crypto/ssh"
    // "time"
)

// TestVersion is the version that the unit tests are run against
const TestVersion = 2

type Device struct {
    hostname string
}

type Connection interface {
    Connect() (chan string, chan string)
}



// Connect sends a command to the device
// func (dev *Device) Connect(user, password string) (stdout chan string, stdin chan string)  {
//     // Check here http://stackoverflow.com/questions/21126195/talking-to-cisco-equipment-using-go-ssh-library
//     stdout = make(chan string)
//     stdin = make(chan string)

//     config := &ssh.ClientConfig{
//         User: user,
//         Auth: []ssh.AuthMethod{ssh.Password(password)},
//     }

//     conn, _ := ssh.Dial("tcp", dev.name+":22", config)
//     session, _ := conn.NewSession()
//     defer session.Close()

// 	stdoutPipe, _ := session.StdoutPipe()
// 	stdinPipe, _ := session.StdinPipe()

//     modes := ssh.TerminalModes{
//         ssh.ECHO:          0,     // disable echoing
//         ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
//         ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
//     }

//     session.RequestPty("xterm", 0, 200, modes)
//     session.Shell()
//     time.Sleep(1 * time.Second)

//     go func() {
//         select {
//         case command := <-stdin:
//             stdinPipe.Write([]byte(command + "\r"))
//         case <-timeout:
//             close(stdout)
//             return
//         }
//     }()


//     return
// }
