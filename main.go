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
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RollConfig struct {
	NumDice  int
	DieType  int
	Modifier int
	Minimum  int
	Maximum  int
}

func main() {
	help := flag.Bool("h", false, "View command usage")
	flag.Parse()

	if *help == true {
		printUsage()
		return
	}

	var input string
	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	}

	config := parseRollString(input)

	switch config.NumDice {
	case 0:
		fmt.Println("error: cannot roll 0 dice")
		os.Exit(1)
	case 1:
		result := 0
		r := rand.IntN(config.DieType) + 1

		if r < config.Minimum {
			r = config.Minimum
		}
		if config.Maximum != -1 && r > config.Maximum {
			r = config.Maximum
		}

		result += r
		fmt.Printf("rolling a d%v\n", config.DieType)

		if config.Modifier != 0 {
			switch config.Modifier > 0 {
			case true:
				fmt.Printf("result: %v + %v\n", result, config.Modifier)
			case false:
				fmt.Printf("result: %v - %v\n", result, config.Modifier)
			}
		}

		result += config.Modifier
		fmt.Printf("result: %v\n", result)
		return
	default:
		result := 0
		for i := range config.NumDice {
			r := rand.IntN(config.DieType) + 1

			if r < config.Minimum {
				r = config.Minimum
			}
			if config.Maximum != -1 && r > config.Maximum {
				r = config.Maximum
			}
			result += r
			fmt.Printf("roll %v: %v\n", i+1, r)
		}
		if config.Modifier != 0 {
			switch config.Modifier > 0 {
			case true:
				fmt.Printf("result: %v + %v\n", result, config.Modifier)
			case false:
				fmt.Printf("result: %v - %v\n", result, config.Modifier)
			}
		}

		result += config.Modifier
		fmt.Printf("total: %v\n", result)
	}

	if len(os.Args) == 1 {
		printUsage()
	}
}

func printUsage() {
	fmt.Println("\nUsage: roll <roll specification>")
	fmt.Println("Examples:")
	fmt.Println("  roll")
	fmt.Println("  roll 2d6")
	fmt.Println("  roll +4")
	fmt.Println("  roll min3")
	fmt.Println("  roll 2d20+12 max16")
}

func parseRollString(input string) RollConfig {
	config := RollConfig{NumDice: 1, DieType: 6, Modifier: 0, Minimum: 1, Maximum: -1}

	if input == "" {
		return config
	}

	// Parse the main roll part
	rollRegex := regexp.MustCompile(`(\d+)?d?(\d+)?([+-]\d+)?`)
	matches := rollRegex.FindStringSubmatch(input)

	if len(matches) > 1 && matches[1] != "" {
		config.NumDice, _ = strconv.Atoi(matches[1])
	}
	if len(matches) > 2 && matches[2] != "" {
		config.DieType, _ = strconv.Atoi(matches[2])
	}
	if len(matches) > 3 && matches[3] != "" {
		config.Modifier, _ = strconv.Atoi(matches[3])
	}

	// Parse the minimum if present
	minRegex := regexp.MustCompile(`min(\d+)`)
	minMatch := minRegex.FindStringSubmatch(input)
	if len(minMatch) > 1 {
		config.Minimum, _ = strconv.Atoi(minMatch[1])
	}

	// Parse the maximum if present
	maxRegex := regexp.MustCompile(`max(\d+)`)
	maxMatch := maxRegex.FindStringSubmatch(input)
	if len(maxMatch) > 1 {
		config.Maximum, _ = strconv.Atoi(maxMatch[1])
	}

	return config
}
