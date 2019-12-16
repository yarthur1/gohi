package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
)

const csvPath string = "/data/redis_memory/memory.csv"
const rdbToolPath string = "/usr/local/bin/rdb"
const rdbPath string = "/data/redis/dump.rdb"

func decodeRDB(csvPath string, rdbPath string) {
    start := time.Now()
    cmd := exec.Command(rdbToolPath, "-c", "memory",rdbPath,"-f",csvPath)
    err := cmd.Start()
    if err != nil {
        fmt.Println("Execute Command failed:" + err.Error())
        os.Exit(1)
    }
    cmd.Wait()

    fmt.Println(cmd.Args)
    elapsed := time.Since(start)
    fmt.Println("decode rdb file time: ", elapsed)
}

func main(){
    decodeRDB(csvPath, rdbPath)
}