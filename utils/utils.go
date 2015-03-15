package utils

import (
	"encoding/json"
	"fmt"
	"github.com/milliyang/dice"
	"os"
)

const (
	NUM_OF_DICE = 3
	FACES       = 6
)

func CheckCasinoPoint(all []*dice.DiceRoll) {
	fmt.Println("testCasinoPoint:")

	resultMap := make(map[int]int)
	evenMap := make(map[int]int)

	samples := len(all)

	for i := 0; i < samples; i++ {

		diceroll := all[i]

		point := 0
		for _, onePoint := range diceroll.Rolls {
			point += onePoint
		}
		sum, ok := resultMap[point]
		if ok {
			resultMap[point] = sum + 1
		} else {
			resultMap[point] = 1
		}

		// simple check
		if diceroll.Rolls[0] == diceroll.Rolls[1] &&
			diceroll.Rolls[1] == diceroll.Rolls[2] {
			point = diceroll.Rolls[0]
			// fmt.Println(diceroll)
			sum, ok := evenMap[point]
			if ok {
				evenMap[point] = sum + 1
			} else {
				evenMap[point] = 1
			}
		}
	}

	fmt.Println("\nPOINT\t  TIMES")
	for i := 3; i <= NUM_OF_DICE*FACES; i++ {
		value, ok := resultMap[i]
		if ok {
			fmt.Println(i, "\t:", value)
		} else {
			fmt.Println(i, "\t: 0")
		}
	}

	fmt.Println("\nEVEN:")
	fmt.Println("WeiSai\t  TIMES")
	for i := 1; i <= FACES; i++ {
		value, ok := evenMap[i]
		if ok {
			fmt.Println(i, i, i, "\t:", value)
		} else {
			fmt.Println(i, i, i, "\t: 0")
		}
	}

	fmt.Println("totalRound:", samples)
}

func CheckRandom(all []*dice.DiceRoll) {
	fmt.Println("\n\nCheckRandom:")

	resultMap := make(map[int]int)

	samples := len(all)

	for i := 0; i < samples; i++ {

		diceroll := all[i]
		for _, onePoint := range diceroll.Rolls {
			sum, ok := resultMap[onePoint]
			if ok {
				resultMap[onePoint] = sum + 1
			} else {
				resultMap[onePoint] = 1
			}
		}
	}

	fmt.Println("Face\t  TIMES")
	totalDice := 0
	for i := 0; i < NUM_OF_DICE*FACES; i++ {
		value, ok := resultMap[i]
		if ok {
			fmt.Println(i, "\t:", value)
			totalDice += value
		}
	}
	fmt.Println("totalRound:", samples)
	fmt.Println("totalDice:", totalDice)
}

/*
RollP() generates a new DiceRoll based on the specified parameters.
*/
func CasinoRoll() *dice.DiceRoll {
	return dice.RollP(NUM_OF_DICE, FACES, 0, false)
}

func JsonPrint(obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	os.Stdout.Write(b)
	fmt.Println("")
}
