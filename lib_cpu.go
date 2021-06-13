package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Cpu struct {
	idleTime uint64
	totTime  uint64
}

type PidStat struct {
	cpu    float64
	uptime float64
}

func GetProcs() int {
	//$ cat /proc/cpuinfo |grep "cpu cores"|uniq
	//cpu cores       : 4
	const FileName = "/proc/cpuinfo"
	var ncores int = int(0)
	f, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for {
		scanner.Scan()
		if line := scanner.Text(); strings.Contains(line, "cpu cores") {
			arr := strings.Split(line, " ")
			d, _ := strconv.ParseInt(arr[2], 10, 8)
			// procs is twice the cores
			ncores = int(d) * 2
			break
		}
	}
	return ncores
}

func GetCurrentCpu(r chan float64) {
	go func() {
		defer close(r)
		r1 := readCpuFile()
		time.Sleep(time.Second * 1)
		r2 := readCpuFile()
		r <- calulateCpu(r1, r2)
	}()
}

//grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$4+$5)} END {print usage "%"}'
//TODO

/*
/proc/stat contains following data:
        user    nice   system  idle      iowait irq   softirq  steal  guest  guest_nice

1.	read the first line of   /proc/stat
2.	discard the first word of that first line   (it's always cpu)
3.	sum all of the times found on that first line to get the total time
4.	divide the fourth column ("idle") by the total time, to get the fraction of time spent being idle
5.	subtract the previous fraction from 1.0 to get the time spent being   not   idle
6	multiple by   100   to get a percentage
*/
func readCpuFile() Cpu {
	//$ cat /proc/stat
	//cpu  74608   2520   24433   1117073   6176   4054  0        0      0      0
	//...
	const FileName = "/proc/stat"
	f, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	//reading the first line of   /proc/stat
	line := scanner.Text()[5:] //to remove "cpu  " the convert to [] string
	lineArr := strings.Fields(line)
	defer f.Close()
	idleTime, _ := strconv.ParseUint(lineArr[3], 10, 64)
	var totTime uint64
	//sum all of the times found on that first line to get the total time
	for _, elem := range lineArr {
		u, _ := strconv.ParseUint(elem, 10, 64)
		totTime += u
	}
	return Cpu{idleTime: idleTime, totTime: totTime}

}

func calulateCpu(last, current Cpu) float64 {
	//cpu % = ( 1 - (idle_time2 -idle_time1)/(total_time2 - total_time1) )*100
	delIdle := current.idleTime - last.idleTime
	delTot := current.totTime - last.totTime
	cpuUsage := (1.0 - float64(delIdle)/float64(delTot)) * 100.0
	return cpuUsage
}

/*

 $cat /proc/225/stat
 225 (scsi_tmf_4) I 2 0 0 0 -1 69238880 0 0 0 0 0 0 0 0 0 -20 1 0 119 0 0 18446744073709551615 0 0 0 0 0 0 0 2147483647 0 0 0 0 17 1 0 0 0 0 0 0 0 0 0 0 0 0 0

#14 utime - CPU time spent in user code, measured in clock ticks

  #15 stime - CPU time spent in kernel code, measured in clock ticks

  #16 cutime - Waited-for children's CPU time spent in user code (in clock ticks)

  #17 cstime - Waited-for children's CPU time spent in kernel code (in clock ticks)

  #22 starttime - Time when the process started, measured in clock ticks

+    Hertz (number of clock ticks per second) of your system.

  In most cases, getconf CLK_TCK can be used to return the number of clock ticks.

  The sysconf(_SC_CLK_TCK) C function call may also be used to return the hertz value.

Calculation
	First we determine the total time spent for the process:
	total_time = utime + stime
	We also have to decide whether we want to include the time from children processes. If we do, then we add those values to total_time:
	total_time = total_time + cutime + cstime
	Next we get the total elapsed time in seconds since the process started:
	seconds = uptime - (starttime / Hertz)
	Finally we calculate the CPU usage percentage:
	cpu_usage = 100 * ((total_time / Hertz) / seconds)
*/

func GetPidCpuAndUptime(pid string) (PidStat, error) {
	fileName := filepath.Join("/proc", pid, "stat")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// sometimes the process gets terminated while scripts are running
		// in those cases the folders get removed
		// in such cases return error
		return PidStat{cpu: 0.0, uptime: 0.0}, errors.New("process got terminated")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)
	uTime, _ := strconv.ParseUint(fields[14], 10, 64)
	sTime, _ := strconv.ParseUint(fields[15], 10, 64)
	cuTime, _ := strconv.ParseUint(fields[16], 10, 64)
	csTime, _ := strconv.ParseUint(fields[17], 10, 64)
	startTime, _ := strconv.ParseUint(fields[14], 10, 64)
	totalTime := uTime + sTime + cuTime + csTime
	hz := GetHz()
	upTime := GetUptime()
	uptime := upTime - float64(startTime/uint64(hz))
	cpu := 100.0 * float64((totalTime / uint64(hz))) / uptime
	return PidStat{cpu: cpu, uptime: uptime}, nil
}
