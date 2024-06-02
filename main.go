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
	for {
		fmt.Print("myshell> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "exit"{
			break
		}

		cmd := exec.Command(input)
		output, err := cmd.Output()

		if err != nil {
			fmt.Println("Error executing command:", err)
			continue
		}

		// Print the output
		fmt.Println(string(output))
	}
}