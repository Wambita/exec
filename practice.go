package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// create a program that reads an input file and writes the file contents to a go file, runs the go file and stores the output in a onutput file
// additional : destroy the go file after use

// 1.declare temp file, executable, outputfile
// 2.read the file with go code
// 3. write the code on to the temporary file
// 4.build the file
// 5.Execute the build file
// 6.Print output
// 7.Remove temporary files
// try to handle multiple programs

func main() {
	// declare variables
	inputFile := "test1.txt"
	temporaryFile := "temp.go"
	executable := "temp"
	outputfile := "output.txt"

	// read file with code

	fileinfo, err := os.Stat(inputFile)
	if os.IsNotExist(err) {
		fmt.Println("Error, file does not exist")
	} else if fileinfo.Size() == 0 {
		fmt.Println("File exists but is empty")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		code, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println(err)
		}

		// write to temporary file
		err = os.WriteFile(temporaryFile, code, 0o644)
		if err != nil {
			fmt.Println(err)
		}

		fileinfo, err := os.Stat(outputfile)
		if os.IsNotExist(err) {
			fmt.Println("Error, file does not exist")
		} else if fileinfo.Size() == 0 {
			fmt.Println("File exists but is empty")
		}

		// compile the file
		command := exec.Command("go", "build", "-o", executable, temporaryFile)

		// buffer to capture stdout and stderr
		var out bytes.Buffer
		command.Stdout = &out
		command.Stderr = &out

		err = command.Run()
		if err != nil {
			fmt.Printf("Compilation failed: %s\n", err)
			fmt.Printf("Compiler output: \n%s", out.String())
		}
		// compile temporary go file

		// Run generated executable and redirect output
		command = exec.Command("./temp")
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("error", err)
		}
		// Write output to desired file
		err = os.WriteFile("output.txt", []byte("wellformated\n"), 0o644)
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = os.WriteFile("outputs.txt", output, 0o644)
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("Successfully wrote output to output.txt")

		// Clean up by removing temporary file

		err = os.Remove("temp.go")
		if err != nil {
			fmt.Println("Error removing temporay file")

			// check if executable is creat
		}
	}
}
