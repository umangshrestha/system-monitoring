the article for basic introduction here:
https://umangshrestha09.medium.com/create-your-own-system-monitoring-tool-in-linux-b860e480b151</br>


# package main

type Bytes struct {
	Tx uint64
	Rx uint64
}

type MemKB struct {
	MemTotal uint64
	MemFree  uint64
	MemUsed uint64
	SwapTotal uint64
	SwapFree uint64
	SwapUsed uint64
	Shared uint64
	Cache uint64
}

# functions
func GetCurrentCpu(r chan float64); <br />
func GetProcs()  int; <br />
func GetHz()  uint; <br />
func GetHostName() string; <br />
func GetInterfaceNames() []string; <br /> 
func GetUptime() float64; <br />
func GetBytes(interfaceName string) Bytes; <br />
func GetMem() MemKB; <br />
func GetListOfPid() []string; <br />
func GetPidCpuAndUptime(pid string) (PidStat, error) </br>
func GetPidMem(pid string) uint64; </br>
func GetPidCommandLine(pid string) (string, error); </br>
func GetPidName(pid string) (string, error); </br>

The output csv of example.go:
//epoch, cpu, mUsed, sUsed, cache, uptimeInMin
