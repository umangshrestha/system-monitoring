package main

import (
	"strconv"
	"testing"
)

func TestGetUptime(t *testing.T) {
	//$ tuptime -sc | grep 'Current uptime' | awk -F\" '{print $4}'
	//1619938259
	const Command = "tuptime -sc | grep 'Current uptime' | awk -F\\\" '{print $4}'"
	data := RunSubprocess(Command)
	expected, _ := strconv.ParseUint(data, 10, 32)
	output := uint64(GetUptime())
	if expected != output {
		t.Errorf("expected: %d, recieved: %d", expected, output)
	}
}
