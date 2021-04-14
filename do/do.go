package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	logFile := os.Getenv("LOGFILE")
	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		log.Fatal(err)
	}
	upd := append(content, []byte(" and from do")...)
	err = ioutil.WriteFile(logFile, upd, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final message: %s", upd)
}
