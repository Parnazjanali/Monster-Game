package game

import (
	"fmt"
	"time"
)

func InitHealth() {
	*CurrentMonsterHealth = 100
	*CurrentPlayerHealth = 100

}
func Play() {
	PrintInstructions()

	Winner := ""
	GameStartTime := time.Now()
	for Winner == "" {
		Winner = executerounds()
	}

	PlayerStatus.TotalRounds = CurrentRound
	PlayerStatus.TotalTime = time.Since(GameStartTime)
	PlayerStatus.LastPlayedAt = time.Now()

	PrintResult(Winner)

	PrintStatus(PlayerStatus)
}

func executerounds() string {
	CurrentRound++
	UserChoice := GetUserChoice()

	switch UserChoice {
	case "Attack":
		*PlayerAttackDmg = PlayersAttack()
		*MonsterAttackDmg = MonsterAttack()
	case "Heal":
		*PlayerhealValue = PlayerSelfHeal()
		*MonsterAttackDmg = MonsterAttack()
	default:
		fmt.Println("Your choice is Invalid type!")
	}
	*CurrentPlayerHealth, *CurrentMonsterHealth = GetHealthAmount()

	RoundData := RoundData{
		Action:           UserChoice,
		PlayerAttackDmg:  *PlayerAttackDmg,
		PlayerSelfHeal:   *PlayerhealValue,
		MonsterAttackDmg: *MonsterAttackDmg,
		PlayerHealth:     *CurrentPlayerHealth,
		MonsterHealth:    *CurrentMonsterHealth,
	}

	RoundHistory = append(RoundHistory, RoundData)
	PrintRounds(&RoundData)

	if *CurrentPlayerHealth <= 0 {
		return "Monster"
	} else if *CurrentMonsterHealth <= 0 {
		return "Player"

	} else {
		fmt.Println("continue!")
	}
	return ""
}
