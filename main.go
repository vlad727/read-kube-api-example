package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// read file
	audit, err := os.ReadFile("/config/kube-apiserver.yaml")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(audit))
}
