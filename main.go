package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(100)

	var guess int
	fmt.Println("Guess the number!")

	for guess != rand {
		fmt.Scanln(&guess)
		if guess > rand {
			fmt.Println("Too high!")
		} else if guess < rand {
			fmt.Println("Too low!")
		}
	}

	fmt.Println("You win!")
}
