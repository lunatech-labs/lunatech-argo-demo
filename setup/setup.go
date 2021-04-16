package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	logFile := os.Getenv("LOGFILE")
	message := "Writing to " + logFile + " from setup"
	err := ioutil.WriteFile(logFile, []byte(message), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("First message: %s", message)
}
