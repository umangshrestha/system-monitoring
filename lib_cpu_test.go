package main

import (
	"strconv"
	"testing"
)

func TestGetCurrentCpu(t *testing.T) {
	//$ echo $[100-$(vmstat 1 2|tail -1|awk '{print $15}')]
	//4
	const Command = "echo $[100-$(vmstat 1 2|tail -1|awk '{print $15}')]"
	c := make(chan float64)
	GetCurrentCpu(c)
	// kindly note that I am not testing decimal accuaracy
	data := RunSubprocess(Command)
	output := int64(<-c)
	expected, _ := strconv.ParseInt(data, 10, 64)
	if expected != output {
		t.Errorf("expected: %d, recieved: %d", expected, output)
	}
}

func TestGetProcs(t *testing.T) {
	//$ nproc
	//8
	const Command = "nproc"
	data := RunSubprocess(Command)
	output := GetProcs()
	e, _ := strconv.ParseInt(data, 10, 64)
	expected := int(e)
	if expected != output {
		t.Errorf("expected: %d, recieved: %d", expected, output)
	}
}
