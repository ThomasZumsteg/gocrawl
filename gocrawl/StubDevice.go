package gocrawl

import (
    "io"
    // "fmt"
    "golang.org/x/crypto/ssh"
)

type stubSession map[string]string

func NewStubDevice(hostname string, responses map[string]string) Device {
    return Device{
        Hostname : hostname,
        Stdin : make(chan string),
        Stdout : make(chan string),
        session : newStubSession(responses),
    }
}


func newStubSession(response map[string]string) remoteSession {
    return nil
}

func (stub *stubSession) StdoutPipe() (io.Reader, error) {
    return nil, nil
}

func (stub *stubSession) StdinPip() (io.WriteCloser, error) {
    return nil, nil
}

func (stub *stubSession) RequestPty(term string, h, w int, termmodes ssh.TerminalModes) error {
    return nil
}

func (stub *stubSession) Shell() error {
    return nil
}
