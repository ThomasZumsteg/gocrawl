package gocrawl

import (
    // "fmt"
    // "golang.org/x/crypto/ssh"
)

func NewStubDevice(hostname string, responses map[string]string) Device {
    return Device{
        Hostname : hostname,
        Stdin : nil,
        Stdout : nil,
        session : stubSession(),
    }
}


func stubSession() remoteSession {
    return nil 
}
