package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestGetMem(t *testing.T) {
	const (
		MemTotal  = 0
		MemUsed   = 1
		MemFree   = 2
		Shared    = 3
		Cache     = 4
		SwapTotal = 5
		SwapUsed  = 6
		SwapFree  = 7
	)
	Command := "free -k"
	//$ ifconfig wlp0s20f3|awk '/packets/ {print $5}'|xargs
	//20299 17045
	// RX    TX
	data := RunSubprocess(Command)
	arr := strings.Split(data, " ")
	var expected []uint64
	for _, d := range arr {
		if data, err := strconv.ParseUint(d, 10, 64); err == nil {
			expected = append(expected, data)
		}
	}
	t.Error(expected)
	output := GetMem()
	if expected[MemUsed] != output.MemUsed {
		t.Errorf("MemUsed: expected: %d, recieved: %d", expected[MemUsed], output.MemUsed)
	}

	if expected[MemTotal] != output.MemTotal {
		t.Errorf("MemTotal: expected: %d, recieved: %d", expected[MemTotal], output.MemTotal)
	}

	if expected[MemFree] != output.MemFree {
		t.Errorf("MemFree: expected: %d, recieved: %d", expected[MemFree], output.MemFree)
	}

	if expected[SwapTotal] != output.SwapTotal {
		t.Errorf("SwapTotal: expected: %d, recieved: %d", expected[SwapTotal], output.SwapTotal)
	}

	if expected[SwapUsed] != output.SwapUsed {
		t.Errorf("SwapUsed: expected: %d, recieved: %d", expected[SwapUsed], output.SwapUsed)
	}

	if expected[SwapFree] != output.SwapFree {
		t.Errorf("SwapFree: expected: %d, recieved: %d", expected[SwapFree], output.SwapFree)
	}

	if expected[Shared] != output.Shared {
		t.Errorf("Shared: expected: %d, recieved: %d", expected[Shared], output.Shared)
	}

	if expected[Cache] != output.Cache {
		t.Errorf("Cache: expected: %d, recieved: %d", expected[Cache], output.Cache)
	}
}
