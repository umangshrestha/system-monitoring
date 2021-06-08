package main

import "testing"

func TestRunSubprocess(t *testing.T) {
	// input arguments for go test
	var flagTest = []struct {
		in  string
		out string
	}{
		{"echo 2", "2"},
		{"echo '1'", "1"},
		{"echo 'apple'", "apple"},
		{"echo 'ball'", "ball"},
	}

	//looping over go test
	for _, data := range flagTest {
		recieved := RunSubprocess(data.in)
		if data.out != recieved {
			t.Errorf("expected: %s len: %d, observed: %s len: %d", data.out, len(data.out), recieved, len(recieved))
		}
	}

}
