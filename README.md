package linux

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


func GetCurrentCpu(r chan float64);
func GetProcs()  int;
func GetHz()  uint;
func GetHostName() string;
func GetInterfaceNames() []string; 
func GetUptime() float64;
func GetBytes(interfaceName string) Bytes;
func GetMem() MemKB

The output csv of example.go is
//epoch, cpu, mUsed, sUsed, cache, uptimeInMin
