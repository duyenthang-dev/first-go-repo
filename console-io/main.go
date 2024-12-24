package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

type User struct {
	firstName      string
	lastName       string
	age            int
	favoriteNumber float32
}

var coffees = make(map[int]string)

var reader *bufio.Reader

func displayMenu() {
	fmt.Println("--  Menu --")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Espresso")
	fmt.Println("3 - Latte")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("Q - Quit")
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	KeyinUserInfo()
}

func KeyinUserInfo() {
	var user User
	user.firstName = readString("Enter your first name")
	user.lastName = readString("Enter your last name")
	user.age = readInt("Enter your age")
	user.favoriteNumber = readFloat("Enter your favorite number")

	fmt.Printf("Hello %s %s, you are %d years old and your favorite number is %.2f\n",
		user.firstName, user.lastName, user.age, user.favoriteNumber)
}

func ChooseCoffe() {
	displayMenu()
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees[1] = "Cappuccino"
	coffees[2] = "Espresso"
	coffees[3] = "Latte"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"

	displayMenu()
	fmt.Printf("Press any key, or press ESC to exit\n")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		choice, _ := strconv.Atoi(string(char))
		fmt.Printf("You chose %s\n", coffees[choice])

		if char == 'q' || char == 'Q' {
			break
		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	fmt.Println("Exiting...")
}

func Prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		Prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		if userInput != "" {
			return userInput
		} else {
			fmt.Println("Please enter a value")
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		Prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			log.Fatal(err)
			fmt.Println("Please enter a valid number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float32 {
	for {
		fmt.Println(s)
		Prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 32)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Please enter a valid float number")
		} else {
			return float32(num)
		}
	}
}
