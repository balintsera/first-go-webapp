package main

import (
    "log"
    "os"
)

var (
	Log      *log.Logger
)

func init() {
    file, _ := os.OpenFile("go-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    Log = log.New(file,
    "TWITTER-EPUB: ",
    log.Ldate|log.Ltime|log.Lshortfile)
}
