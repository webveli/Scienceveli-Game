package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Bitcoin int
	Farms   int
	Animals int
	Rewards []string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	player := Player{Bitcoin: 0, Farms: 0, Animals: 0, Rewards: []string{}}

	fmt.Println("Welcome to Bitcoin Farm Game with Rewards!")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Plant Bitcoin Seeds (+10 Bitcoin)")
		fmt.Println("2. Buy a Farm (50 Bitcoin)")
		fmt.Println("3. Buy Animals (100 Bitcoin)")
		fmt.Println("4. Check Rewards")
		fmt.Println("5. Exit")
		fmt.Print("Choose an action: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			player.Bitcoin += 10
			fmt.Println("You planted Bitcoin seeds! Current Bitcoin:", player.Bitcoin)
			checkForReward(&player)
		case 2:
			if player.Bitcoin >= 50 {
				player.Farms++
				player.Bitcoin -= 50
				fmt.Println("You bought a farm! Total Farms:", player.Farms)
				checkForReward(&player)
			} else {
				fmt.Println("Not enough Bitcoin to buy a farm!")
			}
		case 3:
			if player.Bitcoin >= 100 {
				player.Animals++
				player.Bitcoin -= 100
				fmt.Println("You bought animals! Total Animals:", player.Animals)
				checkForReward(&player)
			} else {
				fmt.Println("Not enough Bitcoin to buy animals!")
			}
		case 4:
			fmt.Println("Your Rewards:", player.Rewards)
		case 5:
			fmt.Println("Thanks for playing! Final Stats:", player)
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func checkForReward(player *Player) {
	if player.Farms > 0 && len(player.Rewards) == 0 {
		player.Rewards = append(player.Rewards, "First Farm Bonus: +50 Bitcoin!")
		player.Bitcoin += 50
		fmt.Println("Congratulations! You received a reward: First Farm Bonus!")
	}

	if player.Animals > 2 && len(player.Rewards) == 1 {
		player.Rewards = append(player.Rewards, "Animal Lover Bonus: +100 Bitcoin!")
		player.Bitcoin += 100
		fmt.Println("Congratulations! You received a reward: Animal Lover Bonus!")
	}
}
