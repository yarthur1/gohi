package main

import (
	"fmt"
	"github.com/go-redis/redis"

	//"strconv"
	"strings"
	"time"
	"reflect"
)

func main(){
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println(time.Now().Format("2006-01-02"))
	//client.SAdd("yxj","adc").Result()
	t,_:=client.ConfigGet("slaveof").Result()
	if t[1]==""{
		fmt.Println("master")
	}else{
		fmt.Println("slave")
	}
	s,_:=client.Info("server").Result()
	fmt.Println(reflect.TypeOf(s).String())
	fmt.Println(s)
	version:=2
	tmp:=strings.Index(s, "redis_version:")
	fmt.Println(tmp)
	fmt.Println(s[tmp+14:tmp+14+6])
	fmt.Println(s[tmp+14:tmp+14+1])
	//lineSplit := strings.Split(s, "\\t")
	//fmt.Println(lineSplit)
	//fmt.Println("888888******",lineSplit[0])
	//if len(lineSplit)>1{
	//	fmt.Println(lineSplit[1])
	//	versionSplit := strings.Split(lineSplit[1], ":")
	//	if len(versionSplit)>1{
	//		tmp:=strings.Split(versionSplit[1], ".")[0]
	//		fmt.Println(tmp)
	//		version,_=strconv.Atoi(tmp)
	//	}
	//}
	fmt.Println(version)
}