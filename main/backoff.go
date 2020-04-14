package main

import (
    "fmt"
    "math/rand"
    "time"

    //"time"
)

const MAXSLEEP = 128

func main() {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    for numsec := 10; numsec <= MAXSLEEP; numsec <<= 1 {
        // TODO

        //if numsec <= MAXSLEEP/2 {
           // time.Sleep(time.Second * time.Duration(numsec))
           // numsec = r.Intn(numsec)+1
            fmt.Println("slepp time(s):", r.Intn(numsec)+1)
        //}
    }
}