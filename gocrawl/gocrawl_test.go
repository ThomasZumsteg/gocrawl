package gocrawl

import (
    "flag"
    "fmt"
    "os"
	"testing"
)

const testVersion = 1

func TestTestVersion(t *testing.T) {
	if TestVersion != testVersion {
		t.Errorf("TestVerions done match: Test version %v, module version %v", testVersion, TestVersion)
	}
}

func TestCreateNewConnecton(t *testing.T) {
	for _, test := range []struct {
		passes       bool
		name         string
		ip           string
		expectedName string
		expectedIP   string
	}{
		{true, "device1", "1.1.1.1", "device1", "1.1.1.1"},
		{true, "device2", "2.2.2.2", "device2", "2.2.2.2"},
		{true, "device3", "3.3.3.3", "device3", "3.3.3.3"},
		{false, "device4", "4.4.4.4", "device5", "5.5.5.5"},
	} {
		dev := NewDevice(test.name, test.ip)
		if test.passes == (dev.name != test.expectedName) {
			t.Errorf("Expected %v, got %v", test.expectedName, dev.name)
		} else if test.passes == (dev.ip != test.expectedIP) {
			t.Errorf("Expected %v, got %v", test.expectedIP, dev.ip)
		}
	}
}

func TestConnect(t *testing.T) {
    fmt.Println(*hostNamePtr)
    dev := NewDevice(*hostNamePtr, *hostIpPtr)
    dev.Connect(*userPtr, *passPtr, "show ver")
}

var userPtr *string
var passPtr *string
var hostNamePtr *string
var hostIpPtr *string

func TestMain(m *testing.M) {
    userPtr = flag.String("user", "", "The user to test ssh connections")
    passPtr = flag.String("pass", "", "The password to test ssh connection")
    hostNamePtr = flag.String("name", "", "The hostname for the device")
    hostIpPtr = flag.String("ip", "", "The IP of the device")
    flag.Parse()
    os.Exit(m.Run())
}
