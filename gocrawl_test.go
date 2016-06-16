package gocrawl

import (
	"testing"
)

const testVersion = 1

func TestCreateNewConnecton(t *testing.T) {
	if TestVersion != testVersion {
		t.Errorf("TestVerions done match: Test version %v, module version %v", testVersion, TestVersion)
	}
	for _, test := range []struct {
		passes       bool
		name         string
		ip           string
		expectedName string
		expectedIp   string
	}{
		{true, "device1", "1.1.1.1", "device1", "1.1.1.1"},
		{true, "device2", "2.2.2.2", "device2", "2.2.2.2"},
		{true, "device3", "3.3.3.3", "device3", "3.3.3.3"},
		{false, "device4", "4.4.4.4", "device5", "5.5.5.5"},
	} {
		dev := newConnection(test.name, test.ip)
		if test.passes == (dev.name != test.expectedName) {
			t.Errorf("Expected %v, got %v", test.expectedName, dev.name)
		} else if test.passes == (dev.ip != test.expectedIp) {
			t.Errorf("Expected %v, got %v", test.expectedIp, dev.ip)
		}
	}
}
