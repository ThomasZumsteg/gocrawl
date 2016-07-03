// +build unit
package gocrawl

import (
    "fmt"
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
    if err := dev.Connect(USER, PASS); err != nil {
        t.Errorf("Got an error on connect %s", err)
    } else if dev.Stdout == nil {
        t.Error("Stdout is nil")
    } else if dev.Stdin == nil {
        t.Error("Stdin is nil")
    }
}

func TestSendCommand(t *testing.T) {
    dev := NewDevice(HOSTNAME)
    if err := dev.Connect(USER, PASS); err != nil {
        t.Errorf("Got an error on connect %s", err)
    }

    fmt.Println("Reading welcome message")
    if response := <-dev.Stdout; response.err == nil {
        t.Logf("Got: %s", response.text)
    } else {
        t.Errorf("Failed to read welcome message: %v", response.err)
    }

    fmt.Println("Sending command")
    dev.Stdin.Write([]byte("show ver"))

    fmt.Println("Reading response")
    if response := <-dev.Stdout; response.err == nil {
        t.Logf("Got: %s", response.text)
    } else {
        t.Errorf("Failed to read message: %v", response.err)
    }
    close (dev.Stdout)
}

