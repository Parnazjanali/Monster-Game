package game

import (
	"fmt"
)

func InitHealth() {
	*MonsterHealth = 100
	*PlayersHealth = 100

}
func Play() {
	PrintInstructions()

	switch GetLayerChoice() {
	case "Attack":
		*playerAttackDmg = PlayersAttack()
		*monsterAttackDmg = MonsterAttack()
		fmt.Printf("You attacked the monster and dealt %d damage.\n", *playerAttackDmg)
		fmt.Printf("The monster attacked you and dealt %d damage.\n", *monsterAttackDmg)

	case "heal":
		*healValue = PlayerSelfHeal()
		*monsterAttackDmg = MonsterAttack()
		fmt.Printf("You healed yourself for %d health.\n", *healValue)
		fmt.Printf("The monster attacked you and dealt %d damage.\n", *monsterAttackDmg)
	default:
		fmt.Println("Your choice is invalid type!")
		return

	}
	GetHealthAmount()
	fmt.Printf("Player's health: %d\n", *PlayersHealth)
	fmt.Printf("Monster's health: %d\n", *MonsterHealth)

	if *PlayersHealth <= 0 {
		fmt.Println("You lost! The monster won.")

	} else if *MonsterHealth <= 0 {
		fmt.Println("Monster lost! You won the game.")

	} else {
		fmt.Println("continue!")
	}
}
