package actions

import (
	"math/rand"
	"time"
)

var ranSource = rand.NewSource(time.Now().UnixNano())
var randgenerator = rand.New(ranSource)

var playerHealth = new(int)
var monsterHealth = new(int)

func InitHealth() {
	*playerHealth = 100
	*monsterHealth = 100
}
func Playerattack() int {
	minAttackPlayer := 5
	maxAttackPlayer := 10

	dmgValue := RandomNumber(minAttackPlayer, maxAttackPlayer)
	*monsterHealth -= dmgValue
	if *monsterHealth < 0 {
		*monsterHealth = 0
	}
	return dmgValue
}
func Monsterattack() int {
	minAttackMonster := 9
	maxAttackMonster := 15

	dmgValue := RandomNumber(minAttackMonster, maxAttackMonster)
	*playerHealth -= dmgValue
	if *playerHealth < 0 {
		*playerHealth = 0
	}
	return dmgValue
}
func Playerheal() int {
	minHealPlayer := 8
	maxHealPlayer := 16
	healValue := RandomNumber(minHealPlayer, maxHealPlayer)
	if *playerHealth+healValue > 100 {
		healValue = 100 - *playerHealth
	}
	*playerHealth += healValue
	return healValue
}
func GetHealthAmount() (int, int) {
	return *playerHealth, *monsterHealth
}
func RandomNumber(min int, max int) int {
	return randgenerator.Intn(max-min+1) + min
}
