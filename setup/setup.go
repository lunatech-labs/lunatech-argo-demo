package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Writing from setup")
	err := ioutil.WriteFile("/app/data/log.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final message: %s", message)
}
