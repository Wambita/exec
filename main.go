package main

import (
	"fmt"
	"os"
	"os/exec"
)

// program to read a prog from A file, store it in a go file, execute the file and write file output in an output file

func main() {
	// read go code from file

	code, err := os.ReadFile("test1.txt")
	if err != nil {
		fmt.Println("Error", err)
	}

	// write code to the temporary file
	err = os.WriteFile("temp.go", code, 0o644)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// compile temporary go file
	cmd := exec.Command("go", "build", "-o", "temp", "temp.go")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
	}

	// Run generated executable and redirect output
	cmd = exec.Command("./temp")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error", err)
	}

	// Write output to desired file
	err = os.WriteFile("output.txt", output, 0o644)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Successfully wrote output to output.txt")

	// Clean up by removing temporary file

	err = os.Remove("temp.go")
	if err != nil {
		fmt.Println("Error removing temporay file")
	}
}
