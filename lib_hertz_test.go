package main

import (
	"strconv"
	"testing"
)

func TestGetHz(t *testing.T) {
	//$ getconf CLK_TCK
	//100
	const Command = "getconf CLK_TCK"
	data := RunSubprocess(Command)
	expected, _ := strconv.ParseInt(data, 10, 8)
	output := GetHz()
	if expected != int64(output) {
		t.Errorf("expected: %d, recieved: %d", expected, output)
	}
}
