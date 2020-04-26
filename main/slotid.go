package main

import (
    "bytes"
    "fmt"
    "hash/crc32"
)

func Hash(key []byte) uint32 {
    const (
        TagBeg = '{'
        TagEnd = '}'
    )
    if beg := bytes.IndexByte(key, TagBeg); beg >= 0 {
        if end := bytes.IndexByte(key[beg+1:], TagEnd); end >= 0 {
            key = key[beg+1 : beg+1+end]
        }
    }
    return crc32.ChecksumIEEE(key)
}

func main(){
    fmt.Print(Hash([]byte("GIFT_CALC_INFO_V2_15870113787667300157")) % 1024)
}