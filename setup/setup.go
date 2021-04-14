package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	message := []byte("Writing from setup")
	logFile := os.Getenv("LOGFILE")
	err := ioutil.WriteFile(logFile, message, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final message: %s", message)
}
