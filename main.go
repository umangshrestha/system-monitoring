package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const toMb = 1024.0
const toMin = 60.0

func monitorSystem(fileName string) {
	c := make(chan float64)
	GetCurrentCpu(c)
	mem := GetMem()
	mUsed := fmt.Sprintf("%.2f", float64(mem.MemUsed)/toMb)
	sUsed := fmt.Sprintf("%.2f", float64(mem.SwapUsed)/toMb)
	cache := fmt.Sprintf("%.2f", float64(mem.Cache)/toMb)
	uptime := GetUptime()
	uptimeInMin := fmt.Sprintf("%.2f", float64(uptime)/toMin)
	epoch := strconv.FormatInt(time.Now().Unix(), 10)
	cpu := fmt.Sprintf("%.2f", <-c)
	data := [][]string{{epoch, cpu, mUsed, sUsed, cache, uptimeInMin}}
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	err = w.WriteAll(data)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Flush()
	defer f.Close()
}

func main() {
	GetProcs()
	//Create folder if not exists
	getCurrentDate := func() string {
		dt := time.Now()
		return dt.Format("2006-01-02")
	}
	dirname := "./logs/system-monitoring"
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		os.Mkdir(dirname, os.ModePerm)
	}

	dateExt := getCurrentDate() + ".csv"
	systemFile := filepath.Join(dirname, "system_"+dateExt)
	for {
		monitorSystem(systemFile)
	}
}
