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
		parts := strings.Split(input, " ")
		command := parts[0]
		parts = parts[1:]

		switch command {

		
		case "exit":
			os.Exit(0)
			
			
		case "cd":
			err = os.Chdir(parts[0])
			if err != nil {
				fmt.Println("Error in changing directory:", err)
			}

			continue
			

		case "pwd" :
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error in changing directory:", err)
			}
			fmt.Println("Current working directory:", dir)
			continue
				
			
		default:

			cmd := exec.Command(command,parts...)
			output, err := cmd.Output()

			if err != nil {
				fmt.Println("Error executing command:", err)
				continue
			}

		// Print the output
		fmt.Println(string(output))}
	}
}