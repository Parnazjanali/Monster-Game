package interactions

import "fmt"

type Output struct {
	Action        string
	Playerhealth  int
	Monsterhealth int
	AttackMonster int
	AttackPlayer  int
	HealPlayer    int
}

func PrintInstructions() {
	fmt.Println("Welcome to the game!")
	fmt.Println("Let's play, choose an action")
	fmt.Println("1)Attack")
	fmt.Println("2)Heal")
}
func OutputResult(data *Output) {
	if data.Action == "Attack" {
		fmt.Printf("Player attackted Monster for %d damage\n", data.AttackPlayer)
	} else {
		fmt.Printf("Player healed for %d value\n", data.HealPlayer)
	}
	fmt.Printf("Monster attacked for %d damage\n", data.AttackMonster)
	fmt.Printf("Player's health=%d\n", data.Playerhealth)
	fmt.Printf("Monster's health=%d\n", data.Monsterhealth)

}
