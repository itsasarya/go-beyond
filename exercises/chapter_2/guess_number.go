package exercise

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func Play_game() {
	solution := rand.Intn(100) + 1
	try := 1
	fmt.Println("Guess a number between 1 and 100:")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Enter a number.")
			continue
		}

		if guess < 1 || guess > 100 {
			fmt.Println("Enter number between 1 and 100")
			continue
		}

		if guess > solution {
			fmt.Println("Too high")
		}
		if guess < solution {
			fmt.Println("Too low")
		}
		if guess == solution {
			fmt.Println("Correct!")
			break
		}
		try += 1
	}
	fmt.Printf("You guessed in %d attenpts.\n", try)
}
