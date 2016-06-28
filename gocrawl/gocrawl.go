package gocrawl

import (
    "io"
    "golang.org/x/crypto/ssh"
)

// TestVersion is the version that the unit tests are run against
const Version = 2

//remoteSession is an individual communication channel with a RemoteClient 
//it wraps ssh.Session for dependency injection to implement unit testing
type remoteSession interface {
    StderrPipe() (io.Reader, error)
    StdinPipe() (io.WriteCloser, error)
    RequestPty(term string, h, w int, termmodes ssh.TerminalModes) error
    Shell() error
}

//Device 
type Device struct {
    Hostname string
    Stdin chan string
    Stdout chan string
    session remoteSession
}

func NewDevice(hostname string) Device {
    return Device{
        Hostname : hostname,
        Stdin : nil,
        Stdout : nil,
        session: getSession(),
    }
}

func getSession() remoteSession {
    return nil
}


func (dev Device) Connect(user, pass string) error {
    return nil
}

func (dev Device) Send(command string) (string, error) {
    return "response", nil
}
