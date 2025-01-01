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
	*Winner = ""
	for *Winner == "" {
		*Winner = RoundsGame()
	}
	PrintResult(*Winner)
}

func RoundsGame() string {
	switch GetUserChoice() {
	case "Attack":
		*playerAttackDmg = PlayersAttack()
		*monsterAttackDmg = MonsterAttack()
		fmt.Printf("You attacked the monster and dealt %d damage.\n", *playerAttackDmg)
		fmt.Printf("The monster attacked you and dealt %d damage.\n", *monsterAttackDmg)

	case "Heal":
		*healValue = PlayerSelfHeal()
		*monsterAttackDmg = MonsterAttack()
		fmt.Printf("You healed yourself for %d health.\n", *healValue)
		fmt.Printf("The monster attacked you and dealt %d damage.\n", *monsterAttackDmg)
	default:
		fmt.Println("Your choice is invalid type!")
		return ""

	}
	GetHealthAmount()
	fmt.Printf("Player's health: %d\n", *PlayersHealth)
	fmt.Printf("Monster's health: %d\n", *MonsterHealth)

	if *PlayersHealth <= 0 {
		return "Monster"
	} else if *MonsterHealth <= 0 {
		return "Player"

	} else {
		fmt.Println("continue!")
	}
	return ""
}
