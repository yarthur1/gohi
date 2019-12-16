package main

import (
    "log"

    "github.com/natefinch/lumberjack"
)

func main() {
    hook := &lumberjack.Logger{
        Filename:   "/Users/yarthur/log/test", //filePath
        MaxSize:    500,                       // megabytes
        MaxBackups: 1,
        MaxAge:     30,    //days
        Compress:   false, // disabled by default
    }
    defer hook.Close()

    _log := log.New(hook, "", log.LstdFlags)
    _log.Println("Start...")
}
