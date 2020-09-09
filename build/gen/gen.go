package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/filecoin-project/statediff/build"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Must specify source directory")
		os.Exit(1)
	}

	cmd := exec.Command("npm", "install")
	cmd.Dir = os.Args[1]
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to install frontend dependencies: %v", err)
		os.Exit(1)
	}

	data := build.Compile(os.Args[1])
	if len(os.Args) < 3 {
		fmt.Printf("%s\n", data)
	}
	ioutil.WriteFile(os.Args[2], []byte(data), 0644)
}
