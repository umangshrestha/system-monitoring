package linux

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cpu struct {
	idleTime uint64
	totTime uint64	
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
		scanner.Scan();
		if line := scanner.Text(); strings.Contains(line, "cpu cores"){
			arr := strings.Split(line, " ")
			d, _ := strconv.ParseInt(arr[2],10, 8)
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
		r <-calulateCpu(r1, r2)
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
func readCpuFile() Cpu{
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
		totTime +=u
	}
	return Cpu{idleTime: idleTime, totTime: totTime}

}

func calulateCpu( last, current Cpu) float64{
	//cpu % = ( 1 - (idle_time2 -idle_time1)/(total_time2 - total_time1) )*100
	delIdle := current.idleTime - last.idleTime
	delTot := current.totTime - last.totTime
	cpuUsage := (1.0 - float64(delIdle)/float64(delTot)) * 100.0
	return cpuUsage
}
