package gocrawl

import (
    "io"
    "golang.org/x/crypto/ssh"
)

// TestVersion is the version that the unit tests are run against
const Version = 2

type RemoteClient interface {
    NewSession() (RemoteSession, error)
}

type RemoteSession interface {
    StderrPipe() (io.Reader, error)
    StdinPipe() (io.WriteCloser, error)
    RequestPty(term string, h, w int, termmodes ssh.TerminalModes) error
    Shell() error
}

type device struct {
    Hostname string
    stdin chan string
    stdout chan string
    Dial func(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error)
}

func NewDevice(hostname string) device {
    return device{
        Hostname : hostname,
        stdin : nil,
        stdout : nil,
        Dial : ssh.Dial,
    }
}

func NewStubDevice(hostname string, responses map[string]string) device {
    return device{
        Hostname : hostname,
        stdin : nil,
        stdout : nil,
        Dial : StubDial(responses),
    }
}

func StubDial(commandResponse map[string]string) func(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error) {
    stubDial := func(network, add string, config *ssh.ClientConfig) (*ssh.Client, error) {
        return nil, nil
    }
    return stubDial
}

func (dev device) Connect(user, pass string) error {
    return nil
}

func (dev device) Send(command string) (string, error) {
    return "response", nil
}
