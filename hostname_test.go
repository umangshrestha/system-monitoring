package main

import (
	"testing"
)

func TestGetHostName(t *testing.T) {
	//$ hostname
	//shrestha
	const Command = "hostname"
	expected := RunSubprocess(Command)
	output := GetHostName()
	if expected != output {
		t.Errorf("expected: %s, recieved: %s", expected, output)
	}
}
