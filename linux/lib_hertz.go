package linux

/*
#include <unistd.h>

long sysconf(int name);

static unsigned int gethz(void)
{
        return sysconf(_SC_CLK_TCK);
}
*/
import "C"


func GetHz() uint{
        // There is no easy way to get 
        // kernels clock ticks per second
        // so we are using c defination for access 
        hz := C.gethz()
        return uint(hz)
}

