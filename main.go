package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	// read and copy file kube-apiserver.yaml to kube-apiserver.yaml.bac
	audit, err := os.ReadFile("/config/kube-apiserver.yaml")
	if err != nil {
		log.Println(err)
	}
	log.Println("Success read the file...")

	// Write data to dstination file
	err = os.WriteFile("/config/kube-apiserver.yaml.bac", audit, 0600)
	if err != nil {
		log.Printf("Can't create copy for file kube-apiserver.yaml\n")
		log.Println(err)
	}

	// list files
	files, err := os.ReadDir("/config/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
	log.Println("Success list the files...")

	//go to a sleep
	time.Sleep(3600 * time.Second)
}


