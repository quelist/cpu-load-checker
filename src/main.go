package main

import (
    "os"
    "strconv"
    "time"
    "cpuobserver"
)


func main() {
    //setting default time interval in miliseconds
    var DEFAULT_INTEVERAL int16 =60

    if len(os.Args) > 1 {
        CUSTOM_INTEVERAL,_ := strconv.Atoi(os.Args[1])
        cpuobserver.DoHealthCheckEvery( time.Duration( CUSTOM_INTEVERAL) * time.Millisecond )
    } else {
        cpuobserver.DoHealthCheckEvery( time.Duration( DEFAULT_INTEVERAL ) * time.Millisecond )
    }
    
}