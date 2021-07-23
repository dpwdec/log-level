package main

import (
	"fmt"
	"bufio"
    "os"
	"strings"
	"github.com/dpwdec/log-level/stats"
	"github.com/dpwdec/log-level/utils"
)

type mapping func(stat int) int

type attribute struct {
	name string
	mappings map[stats.Stats]mapping
}

func main() {
	charStats := map[stats.Stats]int {
		stats.CONSTITUTION: 6,
		stats.TOUGHNESS: 7,
	}

	// initialise attribute mappings
	attributes := [2]attribute {
		attribute {
			name: "Health",
			mappings: map[stats.Stats]mapping {
				stats.CONSTITUTION: func(stat int) int {
					return utils.Round(utils.Log(float64(stat - 1), 1.4) * 21)
				},
				stats.TOUGHNESS: func(stat int) int {
					return utils.Round(utils.Log(float64(stat - 1), 2.1) * 10 - 21)
				},
			},
		},
		attribute {
			name: "Stamina",
			mappings: map[stats.Stats]mapping {
				stats.CONSTITUTION: func(stat int) int {
					return 0
				},
				stats.TOUGHNESS: func(stat int) int {
					return utils.Round(utils.Log(float64(stat - 1), 1.4) * 21)
				},
			},
		},
	}

	buf := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := buf.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		} else {
			// handle input
			switch strings.ReplaceAll(input, "\n", "") {
			case "u":
				for stat := stats.Stats(0); stat < stats.Limit; stat++ {
					charStats[stat]++
				}
			case "d":
				for stat := stats.Stats(0); stat < stats.Limit; stat++ {
					charStats[stat]--
				}
			default:
				fmt.Println("Use \"u\" to incremement stat weighting up\nAnd \"d\" to incrememnt stat weighting down.")
			}

			for _, attr := range attributes {
				attrAmount := 0
				for stat, statMap := range attr.mappings {
					attrAmount += statMap(charStats[stat])
				}
				fmt.Printf("%v: %v\n", attr.name, attrAmount)
			}
		}
	}
}