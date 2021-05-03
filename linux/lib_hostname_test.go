package linux

import (
	"system-monitoring/utils"
	"testing"
)

func TestGetHostName(t* testing.T) {
	//$ hostname
	//shrestha
	const Command  = "hostname"
	expected := utils.RunSubprocess(Command)
	output := GetHostName()
	if expected != output {
		t.Errorf("expected: %s, recieved: %s", expected, output)
	}
}