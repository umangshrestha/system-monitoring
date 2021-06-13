package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type MemKB struct {
	// ram
	MemTotal uint64
	MemFree  uint64
	MemUsed  uint64
	// swap
	SwapTotal uint64
	SwapFree  uint64
	SwapUsed  uint64
	Shared    uint64
	Cache     uint64
}

func GetMem() MemKB {
	/*
		$ cat /proc/meminfo
		MemTotal:       16230988 kB
		MemFree:         7264300 kB
		...
		Buffers:          342284 kB
		Cached:          2892272 kB
		SwapCached:            0 kB
		...
		SwapTotal:       2097148 kB
		SwapFree:        2097148 kB
		...
		Shmem:            697432 kB
		...
		492464 kB
		SReclaimable:     254032 kB
		...
	*/
	const FileName = "/proc/meminfo"
	f, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	var memTot, memFree, buffer, slab uint64
	var swapTot, swapFree uint64
	var cached, shmem, sReclamaible uint64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		switch arr := strings.Fields(scanner.Text()); arr[0] {
		case "MemTotal:":
			memTot, _ = strconv.ParseUint(arr[1], 10, 64)
		case "MemFree:":
			memFree, _ = strconv.ParseUint(arr[1], 10, 64)
		case "Buffers:":
			buffer, _ = strconv.ParseUint(arr[1], 10, 64)
		case "SwapTotal:":
			swapTot, _ = strconv.ParseUint(arr[1], 10, 64)
		case "SwapFree:":
			swapFree, _ = strconv.ParseUint(arr[1], 10, 64)
		case "Cached:":
			cached, _ = strconv.ParseUint(arr[1], 10, 64)
		case "SReclaimable":
			sReclamaible, _ = strconv.ParseUint(arr[1], 10, 64)
		case "Shmem":
			shmem, _ = strconv.ParseUint(arr[1], 10, 64)
		case "Slab":
			slab, _ = strconv.ParseUint(arr[1], 10, 64)
		}
	}
	/*
		free output 	coresponding /proc/meminfo fields
		Mem: total 	MemTotal
		Mem: used 	MemTotal - MemFree - Buffers - Cached - Slab
		Mem: free 	MemFree
		Mem: shared 	Shmem
		Mem: buff/cache 	Buffers + Cached + Slab
		Mem:available 	MemAvailable
		Swap: total 	SwapTotal
		Swap: used 	SwapTotal - SwapFree
		Swap: free 	SwapFree
	*/
	return MemKB{
		MemTotal: memTot,
		MemFree:  memFree,
		// MemUsed: expected: 3914728, recieved: 4104908
		MemUsed:   memTot - memFree - buffer - slab - cached - sReclamaible,
		SwapFree:  swapFree,
		SwapTotal: swapTot,
		SwapUsed:  swapTot - swapFree,
		// Shared: expected: 783256, recieved: 0
		Shared: shmem + sReclamaible,
		//lib_memory_test.go:64: Cache: expected: 2865608, recieved: 2675428
		Cache: buffer + cached + slab + sReclamaible,
	}
}

/*
int get_memory_usage(pid_t pid) {
  int fd, data, stack;
  char buf[4096], status_child[NAME_MAX];
  char *vm;
  sprintf(status_child, "/proc/%d/status", pid);
  if ((fd = open(status_child, O_RDONLY)) < 0)
  	return -1;
  read(fd, buf, 4095);
  buf[4095] = '\0';
  close(fd);
  data = stack = 0;
  vm = strstr(buf, "VmData:");
  if (vm) {
  	sscanf(vm, "%*s %d", &data);
  }
  vm = strstr(buf, "VmStk:");
  if (vm) {
  	sscanf(vm, "%*s %d", &stack);
  }
  return data + stack;
}
*/

func GetPidMem(pid string) (uint64, error) {
	fileName := filepath.Join("/proc", pid, "status")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// sometimes the process gets terminated while scripts are running
		// in those cases the folders get removed
		// in such cases return error
		return 0, errors.New("process got terminated")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data, stack uint64
	for scanner.Scan() {
		switch arr := strings.Fields(scanner.Text()); arr[0] {
		case "VmData:":
			data, _ = strconv.ParseUint(arr[1], 10, 64)
		case "VmStk:":
			stack, _ = strconv.ParseUint(arr[1], 10, 64)

		}

	}
	return data + stack, nil
}
