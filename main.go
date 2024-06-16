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
		piped_commands := strings.Split(input, "|")
		temp_output := ""

		for i := 0; i < len(piped_commands); i++ {
			temp_input := strings.TrimSpace(piped_commands[i])
			parts := strings.Split(temp_input, " ")
			command := parts[0]
			parts = parts[1:]
			if (temp_output) != "" {
				parts = append(parts,temp_output)
			}

			switch command {
			
			case "exit":
				os.Exit(0)

			case "cd":
				err = os.Chdir(parts[0])
				if err != nil {
					fmt.Println("Error in changing directory:", err)
					continue
				}

				

			case "pwd":
				output, err := os.Getwd()
				if err != nil {
					fmt.Println("Error in changing directory:", err)
					continue
				}
				temp_output = string(output)
						
				
			default:
				cmd := exec.Command(command,parts...)
				output, err := cmd.Output()

				if err != nil {
					fmt.Println("Error executing command:", err)
					continue
				}
				temp_output = string(output)
			}
				
			
		}
		if temp_output != "" {
			fmt.Println(temp_output)
		}
	}
}