package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10 and press ENTER when ready.")
	reader.ReadString('\n')

}