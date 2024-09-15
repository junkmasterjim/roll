/*
rolls a selected number of specified dice

to install this script run the following
	- $ go run build -o roll
and place the executable
wherever you keep your shell scripts
(usr/local/bin)
*/

package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

var (
	loose bool
)

func main() {

	for i := range len(os.Args) {
		if strings.Contains(os.Args[i], "-l") || strings.Contains(os.Args[i], "--loose") {
			loose = true
		} else {
			loose = false
		}

		if strings.Contains(os.Args[i], "-h") || strings.Contains(os.Args[i], "--help") {
			PrintHelp()
			return
		}
	}

	switch len(os.Args) {
	case 1:
		Roll(1, 6)
		return

	default:
		arg := os.Args[1]

		// If arg starts with d (ex: d20)
		if strings.Index(arg, "d") == 0 {
			args := strings.Split(arg, "d")
			diceToRoll := GetDice(args[1])
			if diceToRoll == 0 {
				return
			}

			Roll(1, diceToRoll)
			return
		} else {
			// Else we have a string in front of d*
			args := strings.Split(arg, "d")

			if strings.ContainsAny(args[0], "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvqxyz") {
				fmt.Println("err: cannot use char as int")
				fmt.Println("usage: roll 3d20 (rolls #dice)")
				return
			} else {

				if len(args) == 1 && strings.Contains(args[0], "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvqxyz") != true {
					qty, _ := strconv.ParseInt(args[0], 0, 0)
					Roll(int(qty), 6)
					return
				}

				qty, _ := strconv.ParseInt(args[0], 0, 0)
				diceToRoll := GetDice(args[1])
				if diceToRoll == 0 {
					return
				}
				Roll(int(qty), diceToRoll)
			}
		}
		return
	}
}

// ----------------------------------------------
// helper functions
// ----------------------------------------------

func GetDice(d string) int {
	switch d {
	case "4":
		return 4
	case "6":
		return 6
	case "8":
		return 8
	case "10":
		return 10
	case "12":
		return 12
	case "20":
		return 20
	default:
		if loose == true {
			i, _ := strconv.ParseInt(d, 0, 0)
			return int(i)
		} else {
			fmt.Printf("d%s doesn't exist! use --loose (-l) to use any value of die\n", d)
			return 0
		}
	}
}

func PrintHelp() {
	fmt.Println("usage: roll #diceToRoll")

	fmt.Println("\nflags:")
	fmt.Println("  --loose -l: enable loose mode (any number of sided die)")
	fmt.Println("  --help -h: help / info")

	fmt.Println("\nexamples:")
	fmt.Println("  --  roll (rolls a d6)")
	fmt.Println("  --  roll 2d8 (rolls 2 d8)")
	fmt.Println("  --  roll 10d10 (rolls 10 d10)")
	fmt.Println("  --  roll 2d5 -l (rolls 2 d5 in loose mode)")
	fmt.Println("  --  roll 7 (rolls 7 d6)")
}

func Roll(qty int, dice int) {
	if qty < 1 {
		fmt.Println("err: cannot roll less than 1 die")
		return
	}
	if dice < 1 {
		fmt.Println("err: cannot roll a d%a", dice)
		return
	}

	switch qty {
	case 1:
		result := 0
		r := rand.IntN(dice) + 1
		result += r
		fmt.Printf("roll: %v\n", r)

	default:
		fmt.Printf("rolling %d d%d\n", qty, dice)
		result := 0
		for i := range qty {
			r := rand.IntN(dice) + 1
			result += r
			fmt.Printf("roll %v: %v\n", i+1, r)
		}
		fmt.Println("---")
		fmt.Printf("total roll: %v\n", result)

	}

}
