package linux

import (
	"bufio"
	"log"
	"os"
)

func GetHostName() string{
	const FileName = "/etc/hostname"
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
	return line
}