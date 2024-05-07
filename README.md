# Go Program File Executor

This Go program reads a program from a file, stores it in a Go file, compiles and executes the file, and writes the output to another file.

## Usage

1. Create a text file containing the Go code you want to execute (e.g., `test1.txt`).
2. Run the Go program `main.go`.
3. The program will read the code from `test1.txt`, compile it, execute it, and write the output to `output.txt`.

## Requirements

- Go installed on your system.
- Ensure that the input file (`test1.txt`) exists and contains valid Go code.

## Usage Example

Assuming you have a file named `test1.txt` containing the following Go code:

```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println("The sum is:", result)
}

```
After running the Go program main.go, it will generate an executable temp, execute it, and write the output to output.txt, which will contain:
```bash
The sum is: 8
```
### Note

    This program is designed for educational purposes and should be used responsibly.
    Ensure that the code you're executing is safe and does not pose a security risk.
    The generated executable file (temp) and the temporary Go file (temp.go) are cleaned up automatically after execution.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Authors
- [Wambita](https://github.com/Wambita)
