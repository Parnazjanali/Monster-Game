package main

import (
	"fmt"

	interactions "github.com/parnazjanali/Monstergame/Interactions"
	"github.com/parnazjanali/Monstergame/actions"
)

func main() {
	actions.InitHealth()
	interactions.PrintInstructions()

	result := interactions.GetLayerchoice()
	var playerAttackDmg, monsterAttackDmg, healAmount int

	switch result {
	case "Attack":
		playerAttackDmg = actions.Playerattack()
		monsterAttackDmg = actions.Monsterattack()

		fmt.Printf("You attacked the monster and dealt %d damage.\n", playerAttackDmg)
		fmt.Printf("The monster attacked you and dealt %d damage.\n", monsterAttackDmg)

	case "Heal":
		healAmount = actions.Playerheal()
		monsterAttackDmg = actions.Monsterattack()

		fmt.Printf("You healed yourself for %d health.\n", healAmount)
		fmt.Printf("The monster attacked you and dealt %d damage.\n", monsterAttackDmg)

	default:
		fmt.Println("Your choice is invalid type!")
		return
	}
	playerHealth, monsterHealth := actions.GetHealthAmount()

	fmt.Printf("Player's health: %d\n", playerHealth)
	fmt.Printf("Monster's health: %d\n", monsterHealth)

	if playerHealth <= 0 {
		fmt.Println("You lost! The monster won.")

	} else if monsterHealth <= 0 {
		fmt.Println("Monster lost! You won the game.")

	} else {
		fmt.Println("continue!")
	}
}
