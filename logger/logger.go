package logger

import (
	"log"
	"os"
)

var file *os.File
var err error

func Log(msg string) {
	log.SetOutput(file)
	log.Println(msg)
}

func CloseFile() {
	file.Close()
}

func InitialzeLogger() {
	file, err = os.OpenFile("sheddit_log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
