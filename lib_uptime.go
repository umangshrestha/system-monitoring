package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetUptime() float64 {
	/*
		$cat /proc/uptime
		>>34903.79  224873.38
		[0] uptime is the duration in seconds that system has run
		[1] idletime is the duraiton in seconds when processor was not being used by any program
		[return] uptime
	*/
	const FileName = "/proc/uptime"
	file, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	var split []string = strings.Split(line, " ")
	upTime, _ := strconv.ParseFloat(split[0], 2)
	return upTime
}
