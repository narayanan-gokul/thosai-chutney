package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
// Program has been converted to Go via Copilot
func login(filename string) bool {
	fmt.Println("Welcome back to Fighting Food Fragments!")
	loggedIn := false

	for !loggedIn {
		fmt.Print("Please enter your user name: ")
		var inputUsername string
		fmt.Scanln(&inputUsername)

		fmt.Print("Please enter your password: ")
		var inputPassword string
		fmt.Scanln(&inputPassword)

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error: Users database file not found")
			return false
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				username := strings.TrimSpace(parts[0])
				password := strings.TrimSpace(parts[1])
				if inputUsername == username && inputPassword == password {
					loggedIn = true
					break
				}
			}
		}

		if !loggedIn {
			fmt.Println("User name or password invalid, please check and try again.")
		}
	}

	return true
}

func main() {
	success := login("users_database.txt")
	if success {
		fmt.Println("Login successful!")
	}
}
