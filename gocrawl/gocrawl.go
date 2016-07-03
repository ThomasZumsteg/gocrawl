package gocrawl

import (
    "fmt"
    "io"
    "strings"
    // "time"
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

    if sshIn, stdInErr := session.StdinPipe(); stdInErr == nil {
        dev.Stdin = sshIn
    } else {
        return fmt.Errorf("Stdin pipe failed: %v", stdInErr)
    }

    modes := ssh.TerminalModes{
        ssh.ECHO:          1,     // disable echoing
        ssh.TTY_OP_ISPEED: 144000, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 144000, // output speed = 14.4kbaud
    }

    if ptyErr := session.RequestPty("xterm", 80, 0, modes); ptyErr != nil {
        return fmt.Errorf("Request Pty failed: %v", ptyErr)
    }

    if shellErr := session.Shell(); shellErr != nil {
        return fmt.Errorf("Shell failed: %v", shellErr)
    }

    return nil
}

func bufferedRead(source io.Reader) (chan string, chan error) {
    buff := make([]byte, 1000)
    outChan := make(chan string)
    errChan  := make(chan error)
    go func() {
        for {
            if bytes_read, err := source.Read(buff); err != nil {
                errChan <- fmt.Errorf("Error on read: %s", err)
            } else if bytes_read > 0 {
                outChan <- string(buff[:bytes_read])
            }
        }
    }()
    return outChan, errChan
}

func (dev *Device) assignStdout(source io.Reader) {
    dev.Stdout = make(chan response)
    go func() {
        var output string
        fragments, errs := bufferedRead(source)
        for {
            select {
            case fragment := <-fragments:
                output += fragment
                if strings.HasSuffix(output, ">") {
                    dev.Stdout <- response{ output, nil }
                    output = ""
                }
            case <-errs:
                return
            }
        }
    }()
}

