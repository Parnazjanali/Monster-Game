package game

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)
var ranSource = rand.NewSource(time.Now().UnixNano())
var randgenerator = rand.New(ranSource)

func PrintInstructions() {
	fmt.Println("Welcome to the game!")
	fmt.Println("Let's play, choose an action")
	fmt.Println("1)Attack")
	fmt.Println("2)Heal")
}

func GetUserinput() (string, error) {
	fmt.Print("Enter your choice: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("An error occurred while reading the input")
	}
	cleanedInput := strings.TrimSpace(userInput)

	return cleanedInput, nil
}
func GetUserChoice() string {
	for {
		userchoice, _ := GetUserinput()

		if userchoice == "1" {
			return "Attack"
		} else if userchoice == "2" {
			return "Heal"
		}
		fmt.Println("fetching the user input failed, please try again.")
	}
}

func MonsterAttack() int {
	minAttackMonster := MONSTER_ATTACK_MIN_DMG
	maxAttackMonster := MONSTER_ATTACK_MAX_DMG

	dmgValue := RandomNumber(minAttackMonster, maxAttackMonster)
	*PlayersHealth -= dmgValue
	if *PlayersHealth < 0 {
		*PlayersHealth = 0
	}
	return dmgValue

}
func PlayersAttack() int {
	minAttackPlayer := PLAYER_ATTACK_MIN_DMG
	maxAttackPlayer := PLAYER_MAX_ATTACK_DMG

	dmgValue := RandomNumber(minAttackPlayer, maxAttackPlayer)
	*MonsterHealth -= dmgValue
	if *MonsterHealth < 0 {
		*MonsterHealth = 0
	}
	return dmgValue

}
func PlayerSelfHeal() int {
	minHealPlayer := PLAYER_SELF_HEAL_MINVALUE
	maxHealPlayer := PLAYER_SELF_HEAL_MAXVALUE
	healValue := RandomNumber(minHealPlayer, maxHealPlayer)
	if *PlayersHealth+healValue > 100 {
		healValue = 100 - *PlayersHealth
	}
	*PlayersHealth += healValue
	return healValue

}

func RandomNumber(min int, max int) int {
	return randgenerator.Intn(max-min+1) + min
}
func GetHealthAmount() (int, int) {
	return *PlayersHealth, *MonsterHealth
}
func PrintResult(Winner string) {
	fmt.Println("--------------")
	fmt.Println("Game Over!")
	fmt.Printf("%v Won!", Winner)
}
