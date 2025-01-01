package game

import "time"

var CurrentRound = 0
var RoundHistory = []RoundData{}
var CurrentMonsterHealth = new(int)
var CurrentPlayerHealth = new(int)
var Winner = new(string)
var PlayerStatus = new(StatusData)
var PlayerAttackDmg, MonsterAttackDmg, PlayerhealValue = new(int), new(int), new(int)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerSelfHeal   int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}
type StatusData struct {
	TotalRounds  int
	TotalTime    time.Duration
	LastPlayedAt time.Time
}
type GameState struct {
	PlayerRounds  []RoundData
	GameStartTime time.Time
}

var Game = GameState{
	PlayerRounds:  []RoundData{},
	GameStartTime: time.Now(),
}

const (
	PLAYER_ATTACK_MIN_DMG     = 5
	PLAYER_MAX_ATTACK_DMG     = 10
	PLAYER_SELF_HEAL_MINVALUE = 8
	PLAYER_SELF_HEAL_MAXVALUE = 13
	MONSTER_ATTACK_MIN_DMG    = 9
	MONSTER_ATTACK_MAX_DMG    = 15
)
