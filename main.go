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
	statWeighting := 6

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
				statWeighting++
			case "d":
				statWeighting--
			default:
				fmt.Println("Use \"u\" to incremement stat weighting up\nAnd \"d\" to incrememnt stat weighting down.")
			}

			// display weightings
			for _, x := range attributes {
				for stat := stats.Stats(0); stat < stats.Limit; stat++ {
					m := x.mappings[stat](statWeighting)
					fmt.Printf("%v\n", m)
				}
			}
		}
	}
}