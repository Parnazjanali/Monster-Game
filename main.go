package main

import (
	interactions "github.com/parnazjanali/Monstergame/Interactions"
	"github.com/parnazjanali/Monstergame/actions"
)

func main() {
	startGame()
	playingGame()

}
func startGame() {
	interactions.PrintInstructions()

}
func playingGame() {
	userChoice := interactions.GetLayerchoice()
	var playerAttackdmg int
	var monsterAttackdmg int
	var healPlayer int

	if userChoice == "Attack" {
		playerAttackdmg = actions.Playerattack()
	} else {
		healPlayer = actions.Playerheal()
	}
	monsterAttackdmg = actions.Monsterattack()

	playerhealth, monsterhealth := actions.GethealthAmount()

	data := interactions.Output{
		Action:        userChoice,
		Playerhealth:  playerhealth,
		Monsterhealth: monsterhealth,
		AttackMonster: monsterAttackdmg,
		AttackPlayer:  playerAttackdmg,
		HealPlayer:    healPlayer,
	}

	interactions.OutputResult(&data)

}
