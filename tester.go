package main

import (
	"os/exec"
	"log"
	"runtime"
)

func main()  {

	log.Println(runtime.GOOS)
	cmd := exec.Command("cmd", "/C", "echo luowen")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	log.Println(string(output))
}
