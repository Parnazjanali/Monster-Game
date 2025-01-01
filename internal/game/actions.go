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
	*CurrentPlayerHealth -= dmgValue
	if *CurrentPlayerHealth < 0 {
		*CurrentPlayerHealth = 0
	}
	return dmgValue

}
func PlayersAttack() int {
	minAttackPlayer := PLAYER_ATTACK_MIN_DMG
	maxAttackPlayer := PLAYER_MAX_ATTACK_DMG

	dmgValue := RandomNumber(minAttackPlayer, maxAttackPlayer)
	*CurrentMonsterHealth -= dmgValue
	if *CurrentMonsterHealth < 0 {
		*CurrentMonsterHealth = 0
	}
	return dmgValue

}
func PlayerSelfHeal() int {
	minHealPlayer := PLAYER_SELF_HEAL_MINVALUE
	maxHealPlayer := PLAYER_SELF_HEAL_MAXVALUE

	healValue := RandomNumber(minHealPlayer, maxHealPlayer)
	if *CurrentPlayerHealth+healValue > 100 {
		healValue = 100 - *CurrentPlayerHealth
	}
	*CurrentPlayerHealth += healValue
	return healValue

}

func RandomNumber(min int, max int) int {
	return randgenerator.Intn(max-min+1) + min
}
func GetHealthAmount() (int, int) {
	return *CurrentPlayerHealth, *CurrentMonsterHealth
}
func PrintRounds(Round *RoundData) {
	if Round.Action == "Attack" {
		fmt.Printf("Player Attacked Monster for %v damage.\n", Round.PlayerAttackDmg)
	} else {
		fmt.Printf("Player Healed self for %v value.\n", Round.PlayerSelfHeal)
	}
	fmt.Printf("Monster Attacked Player for %v damage.\n", Round.MonsterAttackDmg)
	fmt.Printf("Player's Health:%v\n", Round.PlayerHealth)
	fmt.Printf("Monster's Health:%v\n", Round.MonsterHealth)
}

func PrintStatus(PlayerStatus *StatusData) {
	fmt.Println("Game Status:")
	fmt.Printf("Total Rounds: %d\n  Total Time: %s\n  Last Played At: %s\n",
		PlayerStatus.TotalRounds,
		PlayerStatus.TotalTime.String(),
		PlayerStatus.LastPlayedAt.Format("2006-01-02 15:04:05"),
	)
	fmt.Println("------------------------")
}
func CalculateStatus(rounds []*RoundData, StartAt time.Time) StatusData {
	totalRounds := len(rounds)
	totalTime := time.Duration(0)
	lastPlayedAt := StartAt

	return StatusData{
		TotalRounds:  totalRounds,
		TotalTime:    totalTime,
		LastPlayedAt: lastPlayedAt,
	}
}
func PrintResult(Winner string) {
	fmt.Println("Game Over!")
	fmt.Printf("%v Won!", Winner)
}
