package main

import (
    "bufio"
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

//get请求  https://www.cnblogs.com/-beyond/p/8950222.html
//json 转化  https://segmentfault.com/q/1010000005025933   "10.66.103.196:18015"
// flag获取多个值  https://cloud.tencent.com/developer/article/1174479

// pipeline
// http://vearne.cc/archives/1113
// https://studygolang.com/articles/14605

type Stats struct {
    Group struct {
        Stats map[string]interface{} `json:"stats"`
    } `json:"group"`
    HA struct {
        Masters map[string]string `json:"masters"`
    } `json:"sentinels"`
}

var groupStat Stats
var dashBoardUrl string
var dataFile string = "./data.log"
var Master int = 0

func main() {
    flag.StringVar(&dashBoardUrl, "d", "10.66.103.196:18015", "dashboard url")
    flag.StringVar(&dataFile, "data", "./data.log", "data file")
    flag.IntVar(&Master, "master", 0, "master or slave")
    flag.Parse()

    csvFile, err := os.OpenFile(dataFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_SYNC, 0644)
    if err != nil {
        fmt.Printf("create file error: %v\n", err)
        os.Exit(1)
    }
    defer csvFile.Close()
    csvWriter := bufio.NewWriterSize(csvFile, 1<<20) //1MB

    response, err := http.Get("http://" + dashBoardUrl + "/topom/stats")
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    //fmt.Println(string(body))

    if err = json.Unmarshal([]byte(body), &groupStat); err != nil { //集群要有sentinel才行
        panic(err)
    }

    var redisAddr []string
    var masterAddr []string
    var slaveAddr []string
    for item, _ := range groupStat.Group.Stats { //all
        redisAddr = append(redisAddr, item)
    }
    for _, item := range groupStat.HA.Masters {
        masterAddr = append(masterAddr, item)
    }
    mMap := make(map[string]int)
    for _, m := range masterAddr { //用map去重 可以用空结构体实现set
        mMap[m] = 1
    }
    for _, item := range redisAddr {
        if _, exists := mMap[item]; !exists {
            slaveAddr = append(slaveAddr, item)
        }
    }
    if Master == 0 {
        redisAddr = slaveAddr[:]
    } else {
        redisAddr = masterAddr[:]
    }

    for _, item :=range redisAddr{
        csvWriter.WriteString(strings.Split(item, ":")[0]+"\n")
    }

    fmt.Printf("redis addr len:%v\n", len(redisAddr))
    csvWriter.Flush()
    csvFile.Sync()
}
