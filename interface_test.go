package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestGetInterfaceNames(t *testing.T) {
	//$ ifconfig |grep -oE "^[a-z0-9]+"|xargs
	//docker0 enp6s0f1 lo wlp0s20f3
	const Command = "ifconfig |grep -oE \"^[a-z0-9]+\"|xargs"
	data := RunSubprocess(Command)
	expected := strings.Split(data, " ")
	o := GetInterfaceNames()
	// the interface names are not sorted. So for testing we need to sort them first.
	sort.Slice(o, func(i int, j int) bool { return o[i] < o[j] })
	for i, output := range o {
		if expected[i] != output {
			t.Errorf("expected: %s, recieved: %s", expected[i], output)
		}
	}
}

func TestGetBytes(t *testing.T) {
	for _, interfaceName := range GetInterfaceNames() {
		const (
			Rx = 0
			Tx = 1
		)
		Command := fmt.Sprintf("ifconfig %s|awk '/packets/ {print $5}'|xargs", interfaceName)
		//$ ifconfig wlp0s20f3|awk '/packets/ {print $5}'|xargs
		//20299 17045
		// RX    TX
		data := RunSubprocess(Command)
		arr := strings.Split(data, " ")
		var expected []uint64
		for _, data := range arr {
			d, _ := strconv.ParseUint(data, 10, 64)
			expected = append(expected, d)
		}
		output := GetBytes(interfaceName)
		if expected[Rx] != output.Rx {
			t.Errorf("expected: %d, recieved: %d", expected[Rx], output.Rx)
		}

		if expected[Tx] != output.Tx {
			t.Errorf("expected: %d, recieved: %d", expected[Tx], output.Tx)
		}
	}
}
