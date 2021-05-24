package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Bytes struct {
	Tx uint64
	Rx uint64
}

func GetInterfaceNames() []string {
	//$ ls /sys/class/net/
	//docker0  enp6s0f1  lo  wlp0s20f3

	const FolderInterface = "/sys/class/net/"
	file, err := os.Open(FolderInterface)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	names, _ := file.Readdirnames(0)
	return names
}

func readStatistics(interfaceName string, tx string) uint64 {
	//$ cat /sys/class/net/wlp0s20f3/statistics/tx_bytes
	//470798533
	fileName := fmt.Sprintf("/sys/class/net/%s/statistics/%s_bytes", interfaceName, tx)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	bytes, _ := strconv.ParseUint(line, 10, 64)
	return bytes
}

func GetBytes(interfaceName string) Bytes {
	return Bytes{
		Tx: readStatistics(interfaceName, "tx"),
		Rx: readStatistics(interfaceName, "rx"),
	}
}
