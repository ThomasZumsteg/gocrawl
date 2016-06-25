package gocrawl

import (
	"testing"
)

const testVersion = 2

func TestTestVersion(t *testing.T) {
	if TestVersion != testVersion {
		t.Errorf("TestVerions done match: Test version %v, module version %v", testVersion, TestVersion)
	}
}

func TestConnect(t *testing.T) {
    dev := Device{ "Hostname" }
    stdout, stdin, err := OpenConnection("Hostname", "Username", "Password", stubConnection)
    if err != nil {
        t.Error("Got an error on connect")
    }
}

func TestSendCommand(t *testing.T) {
    stdout, stdin, err := OpenConnection("Hostname", "Username", "Password", 

