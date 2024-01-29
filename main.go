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
		exec.Command("echo", fmt.Sprintf("outname=%v", args[1])).Run()
		// os.Setenv("GITHUB_OUTPUT", fmt.Sprintf("\noutname=%v", args[1]))
		return
	}
	log.Println("Hello world!!")
}
