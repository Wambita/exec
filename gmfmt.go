package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// read go code from file
	filePath := "test1.txt"
	code, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error", err)
	}

	// write code to the temporary file
	err = os.WriteFile("temp.go", code, 0o644)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// compile temporary go file
	cmd := exec.Command("gofmt", "-e", "temp.go")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
	} else {
		err = os.WriteFile("output.txt", []byte("well formatted\n"), 0o644)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("successfully wrote to output.txt")
	}
	// Clean up by removing temporary file

	err = os.Remove("temp.go")
	if err != nil {
		fmt.Println("Error removing temporay file")
	}
}
