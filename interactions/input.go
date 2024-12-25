package interactions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetUserinput() (string, error) {
	fmt.Print("Enter your choice: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("An error occurred while reading the input")
	}
	cleanedInput := strings.TrimSpace(userInput)

	return cleanedInput, nil
}
func GetLayerchoice() string {
	for {
		userchoice, _ := GetUserinput()

		if userchoice == "1" {
			return "Attack"
		} else if userchoice == "2" {
			return "Heal"
		}
		fmt.Println("fetching the user input failed, please try again.")
	}
}
