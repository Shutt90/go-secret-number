package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	time := rand.NewSource(time.Now().UnixNano())
	newNum := rand.New(time)
	var lives int = 5

	var randomNum int = newNum.Intn(100)

	fmt.Println("Guess the secret number(max 100): ")

	var guess int

	for guess != randomNum {
		fmt.Scanln(&guess)

		if guess > randomNum {
			fmt.Println("Try lower")
		} else if guess < randomNum {
			fmt.Println("Try higher")
		}

		if guess == randomNum {
			fmt.Println("You win")
		} else {
			lives--
			if lives == 0 {
				fmt.Println("You lose")
				fmt.Scanf("h")
				break
			}
			fmt.Println(lives, "lives left")
		}

	}

}
