package game

var MonsterHealth = new(int)
var PlayersHealth = new(int)

var playerAttackDmg, monsterAttackDmg, healValue = new(int), new(int), new(int)

var Winner = new(string)

const PLAYER_ATTACK_MIN_DMG = 5
const PLAYER_MAX_ATTACK_DMG = 10
const PLAYER_SELF_HEAL_MINVALUE = 8
const PLAYER_SELF_HEAL_MAXVALUE = 13
const MONSTER_ATTACK_MIN_DMG = 9
const MONSTER_ATTACK_MAX_DMG = 15
