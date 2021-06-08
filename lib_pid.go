package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getPidofProcess() map[string][]string {
	/*
		$ ls /proc/[0-9]+
		 114
		$ cat /proc/[0-9]/comn
		 python3
		[return]
		{
			python: [1,2,3],
			mysql: [5, 6],
		}
	*/
	dict := make(map[string][]string)
	files, err := filepath.Glob("/proc/[0-9]*")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		procNum := strings.Split(file, "/")[2]
		f, err := os.Open(filepath.Join(file, "comm"))
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		a := scanner.Text()

		if val, ok := dict[a]; ok {
			dict[a] = append(val, procNum)
		} else {
			dict[a] = []string{procNum}
		}
	}
	return dict
}
