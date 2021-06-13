package main

/*
#include <unistd.h>

long sysconf(int name);

static unsigned int gethz(void)
{
        return sysconf(_SC_CLK_TCK);
}
*/
import "C"
import (
	"fmt"
	"log"
)

func GetHz() uint {
	// There is no easy way to get
	// kernels clock ticks per second
	// so we are using c defination for access
	hz := sysconf(_SC_CLK_TCK)
	return uint(hz)
}

func main() {
	a, err := GetPidName("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}
