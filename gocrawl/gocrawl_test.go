package gocrawl

import (
	"testing"
)

const version = 2

var stubCommands = map[string]string{
    "Hello": "Hi there",
    "What's your name?": "Stubbs",
}

func TestVersion(t *testing.T) {
	if Version != version {
		t.Errorf("Verions done match: Test version %v, module version %v", version, Version)
	}
}

func TestCreateDevice(t *testing.T) {
    dev := NewStubDevice( "Hostname", stubCommands)
    if dev.Hostname != "Hostname" {
        t.Error("Got an error on connect")
    }
}

func TestCreateConnection(t *testing.T) {
    dev := NewStubDevice( "Hostname" , stubCommands)
    err := dev.Connect("Username", "Password")
    if err != nil {
        t.Error("Got an error on connect")
    }
}

func TestSend(t *testing.T) {
    dev := NewStubDevice( "Hostname" , stubCommands)
    err := dev.Connect("Username", "Password")
    response, err := dev.Send("command")
    if err != nil {
        t.Errorf("No error expected, got: %v", err)
    } else if response != "response" {
        t.Error("Expected \"response\": got " + response)
    }
}

func TestConnectError(t *testing.T) {
    dev := NewStubDevice( "Hostname" , stubCommands)
    err := dev.Connect("Username", "Password")
    if err != nil {
        t.Error("Error was expected, no error returned")
    }
}
