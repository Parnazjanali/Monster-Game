package actions

import (
	"math/rand"
	"time"
)

var ranSource = rand.NewSource(time.Now().UnixNano())
var randgenerator = rand.New(ranSource)
var playerhealth = 100
var monsterhealth = 100

func Playerattack() int {
	minAttackPlayer := 5
	maxAttackPlayer := 10

	dmgValue := RandomNumber(minAttackPlayer, maxAttackPlayer)
	monsterhealth -= dmgValue
	return dmgValue

}
func Monsterattack() int {
	minAttackMonster := 9
	maxAttackMonster := 15

	dmgValue := RandomNumber(minAttackMonster, maxAttackMonster)
	playerhealth -= dmgValue
	return dmgValue

}
func Playerheal() int {
	minHealPlayer := 8
	maxHealPlayer := 16
	healValue := RandomNumber(minHealPlayer, maxHealPlayer)
	if playerhealth+healValue > 100 {
		healValue = 100 - playerhealth // مقدار قابل اضافه کردن محدود شود
	}
	playerhealth += healValue
	return healValue
}
func GethealthAmount() (int, int) {
	return playerhealth, monsterhealth

}

func RandomNumber(min int, max int) int {
	return randgenerator.Intn(max-min+1) + min

}
