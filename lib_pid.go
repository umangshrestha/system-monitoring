package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetListOfPid() []string {
	/*
		reading only folders that have integer name
		$ ls /proc/|grep -i "[0–9]"
		10
		1091
		119
		120
		each number represents the pid of running process
	*/
	pidFolder, err := filepath.Glob("/proc/[0-9]*")
	if err != nil {
		log.Fatal(err)
	}
	for i, file := range pidFolder {
		pidFolder[i] = strings.Split(file, "/")[2]
	}
	return pidFolder
}

func GetPidName(pid string) (string, error) {
	// $  cat /proc/11349/comm
	// python3
	fileName := filepath.Join("/proc", pid, "comm")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// sometimes the process gets terminated while scripts are running
		// in those cases the folders get removed
		// in such cases return error
		return "", errors.New("process got terminated")
	}
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	processName := scanner.Text()
	return processName, nil
}

func GetPidCommandLine(pid string) (string, error) {
	// $  cat /proc/11349/cmdLine
	// python3 -i
	fileName := filepath.Join("/proc", pid, "cmdline")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// sometimes the process gets terminated while scripts are running
		// in those cases the folders get removed
		// in such cases return error
		return "", errors.New("process got terminated")
	}
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	processName := scanner.Text()
	return processName, nil
}
