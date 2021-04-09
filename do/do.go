package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("/app/data/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	upd := append(content, []byte(" and from do")...)
	err = ioutil.WriteFile("/app/data/log.txt", upd, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final message: %s", upd)
}
