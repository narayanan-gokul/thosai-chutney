package main
//Program has been converted into Go from Python
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func signup(filename string) bool {
	fmt.Println("Welcome to Fighting Food Fragments!")

	var typeValid, postcodeValid bool
	var typeUser int

	// Get user type
	for !typeValid {
		fmt.Println(`What are your needs? Press:
1 - I am an individual in need of food.
2 - I am a food supplier that has food to give away.
3 - I am a food hub that needs more food.`)

		var input string
		fmt.Scanln(&input)
		num, err := strconv.Atoi(input)
		if err == nil && num >= 1 && num <= 3 {
			typeUser = num
			typeValid = true
		} else {
			fmt.Println("Please try again (1–3)")
		}
	}

	// Get username
	fmt.Print("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Get postcode
	var postcode int
	for !postcodeValid {
		fmt.Print("Where are you? Please enter your postcode: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err == nil && num > 2000 && len(input) == 4 {
			postcode = num
			postcodeValid = true
		} else {
			fmt.Println("Please enter a valid postcode.")
		}
	}

	// Get password
	fmt.Print("Please enter your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Printf("Your user name will be %s and your password is %s.\n", username, password)

	// Write to file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return false
	}
	defer file.Close()

	entry := fmt.Sprintf("%d,%s,%d,%s\n", typeUser, username, postcode, password)
	if _, err := file.WriteString(entry); err != nil {
		fmt.Println("Error saving user data:", err)
		return false
	}

	return true
}

func main() {
	success := signup("user_data.txt")
	if success {
		fmt.Println("Signup successful!")
	}
}
