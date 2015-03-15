package main

import (
	"flag"
	"fmt"
	"github.com/milliyang/dice"
	"github.com/milliyang/dice/utils"
	"os"
)

const (
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

	var allDiceRolls []*dice.DiceRoll

	for i := 0; i < samples; i++ {
		diceroll := utils.CasinoRoll()
		allDiceRolls = append(allDiceRolls, diceroll)

	}

	utils.CheckCasinoPoint(allDiceRolls)
	utils.CheckRandom(allDiceRolls)
}
