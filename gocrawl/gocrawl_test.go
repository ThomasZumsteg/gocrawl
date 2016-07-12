package gocrawl

import (
    // "log"
	"testing"
)

const version = 2

func TestVersion(t *testing.T) {
	if Version != version {
		t.Errorf("Verions done match: Test version %v, module version %v", version, Version)
	}
}

func TestConnect(t *testing.T) {
    dev := NewDevice(HOSTNAME)
    if welcome, err := dev.Connect(USER, PASS); err != nil {
        t.Errorf("Got an error on connect %s", err)
    } else if dev.stdout == nil {
        t.Error("Stdout is nil")
    } else if dev.stdin == nil {
        t.Error("Stdin is nil")
    } else if welcome == "" {
        t.Error("Empty welcome message")
    } else {
        t.Logf("Got welcome message\n----------\n%s\n---------", welcome)
    }
}

func TestSendCommand(t *testing.T) {
    dev := NewDevice(HOSTNAME)
    t.Logf("Connecting")
    if _, err := dev.Connect(USER, PASS); err != nil {
        t.Errorf("Got an error on connect %s", err)
    } else {
        t.Logf("Connected")
    }

    t.Logf("Sending command")

    if response, err := dev.Send("show ver"); err == nil {
        t.Logf("Got: %s", response)
    } else {
        t.Errorf("Failed to read message: %v", err)
    }
}
