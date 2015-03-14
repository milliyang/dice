package main

import (
	"flag"
	"fmt"
	"github.com/milliyang/dice"
	"os"
)

const (
	ROUNDS = 3
	FACES  = 6

	TEST_SAMPLES = 100000
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  "+os.Args[0]+" [<roll description>...]\n")
	fmt.Fprintf(os.Stderr, "  -help\tprints this help message\n")
	flag.PrintDefaults()
}

var samples int

func init() {
	flag.IntVar(&samples, "n", TEST_SAMPLES, "sample size")
}

func main() {
	flag.Parse()

	testCasinoPoint()
	testRandom()
}

func testCasinoPoint() {
	fmt.Println("testCasinoPoint:")

	resultMap := make(map[int]int)
	evenMap := make(map[int]int)

	for i := 0; i < samples; i++ {

		diceroll := casinoRoll()

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

	for i := 3; i <= ROUNDS*FACES; i++ {
		value, ok := resultMap[i]
		if ok {
			fmt.Println(i, ":", value)
		} else {
			fmt.Println(i, ": 0")
		}
	}

	fmt.Println("EVEN:")
	for i := 1; i <= FACES; i++ {
		value, ok := evenMap[i]
		if ok {
			fmt.Println(i, i, i, ":", value)
		} else {
			fmt.Println(i, i, i, ": 0")
		}
	}

	fmt.Println("totalRound:", samples)
}

func testRandom() {
	fmt.Println("\n\ntestRandom:")

	resultMap := make(map[int]int)

	for i := 0; i < samples; i++ {

		diceroll := casinoRoll()
		for _, onePoint := range diceroll.Rolls {
			sum, ok := resultMap[onePoint]
			if ok {
				resultMap[onePoint] = sum + 1
			} else {
				resultMap[onePoint] = 1
			}
		}
	}

	totalDice := 0
	for i := 0; i < ROUNDS*FACES; i++ {
		value, ok := resultMap[i]
		if ok {
			fmt.Println(i, ":", value)
			totalDice += value
		}
	}
	fmt.Println("totalRound:", samples)
	fmt.Println("totalDice:", totalDice)
}

/*
RollP() generates a new DiceRoll based on the specified parameters.
*/
func casinoRoll() *dice.DiceRoll {
	return dice.RollP(ROUNDS, FACES, 0, false)
}
