package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"os/exec"
)

func main() {
	reader := bufio.NewReader(os.Stdin) 
    fmt.Print("myshell> ")      

    input, err := reader.ReadString('\n') 
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    input = strings.TrimSpace(input)

	cmd := exec.Command(input)
	output, err := cmd.Output()

	if err != nil {
        fmt.Println("Error executing command:", err)
        return
    }

    // Print the output
    fmt.Println(string(output))
}