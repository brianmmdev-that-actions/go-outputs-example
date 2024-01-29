package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		log.Printf("Hello %v!!", args[1])
		err := exec.Command("echo", fmt.Sprintf("\"outname=%v\" >> $GITHUB_OUTPUT", args[1])).Run()
		if err != nil {
			log.Fatal(err)
		}
		err = exec.Command("echo", fmt.Sprintf("::set-output name=outname::%v", args[1])).Run()
		if err != nil {
			log.Fatal(err)
		}
		// os.Setenv("GITHUB_OUTPUT", fmt.Sprintf("\noutname=%v", args[1]))
		return
	}
	log.Println("Hello world!!")
}
