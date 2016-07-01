package gocrawl

import (
    "bytes"
    "fmt"
    "io"
    "time"
    "golang.org/x/crypto/ssh"
)

// TestVersion is the version that the unit tests are run against
const Version = 2

//remoteSession is an individual communication channel with a RemoteClient 
//it wraps ssh.Session for dependency injection to implement unit testing
type remoteSession interface {
    StdoutPipe() (io.Reader, error)
    StdinPipe() (io.WriteCloser, error)
    RequestPty(term string, h, w int, termmodes ssh.TerminalModes) error
    Shell() error
}

type response struct {
    text string
    err error
}

//Device is a remote device that can communicate by sending commands to Stdin
// and recieves responses from Stdout
type Device struct {
    Hostname string
    Stdin io.WriteCloser
    Stdout chan response
}

func NewDevice(hostname string) Device {
    return Device{
        Hostname : hostname,
        Stdin : nil,
        Stdout : nil,
    }
}

func (dev *Device) Connect(user, pass string) error {
    config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}

    client, clientErr := ssh.Dial("tcp", dev.Hostname+":22", config)
    if clientErr != nil {
        return fmt.Errorf("Dialing %s failed: %v", dev.Hostname, clientErr)
    }

    session, sessionErr := client.NewSession()
    if sessionErr != nil {
        return fmt.Errorf("Creating a session failed: %v", sessionErr)
    }

    if sshOut, stdOutErr := session.StdoutPipe(); stdOutErr == nil {
        dev.assignStdout(sshOut)
    } else {
        return fmt.Errorf("Stdout pipe failed: %v", stdOutErr)
    }

    sshIn, stdInErr := session.StdinPipe()
    if stdInErr != nil {
        return fmt.Errorf("Stdin pipe failed: %v", stdInErr)
    }
    dev.Stdin = sshIn

    modes := ssh.TerminalModes{
        ssh.ECHO:          0,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }

    ptyErr := session.RequestPty("vt100", 80, 0, modes)
    if ptyErr != nil {
        return fmt.Errorf("Request Pty failed: %v", ptyErr)
    }

    shellErr := session.Shell()
    if shellErr != nil {
        return fmt.Errorf("Shell failed: %v", shellErr)
    }

    return nil
}

func (dev *Device) assignStdout(source io.Reader) {
    dev.Stdout = make(chan response)
    timeout := time.Millisecond * 100
    terminator := []byte(">")
    go func() {
        var buff []byte
        bytes_read := 0
        err := error(nil)
        var output bytes.Buffer
        for start := time.Now();; time.Sleep(timeout) {
            if time.Now().After(start.Add(timeout * 10)) {
                dev.Stdout <- response{ "", fmt.Errorf("Timed out") }
            }

            bytes_read, err = source.Read(buff)
            if err != nil {
                dev.Stdout <- response{"", fmt.Errorf("Error on read: %s", err)}
            }
            output.Write(buff[:bytes_read])

            if bytes.HasSuffix(buff[:bytes_read], terminator) {
                dev.Stdout <- response{output.String(), nil}
                output.Reset()
            }
        }
    }()
}

