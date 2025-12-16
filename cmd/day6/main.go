package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		part := flag.Int("part", 1, "which part to run")
		flag.Parse()
	*/

	scanner := bufio.NewScanner(os.Stdin)

	tokens := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		tokens = append(tokens, fields)
	}

	total := 0
	for i := 0; i < len(tokens[0]); i++ {
		operator := tokens[len(tokens)-1][i]
		var isMul bool
		switch operator {
		case "*":
			isMul = true
		case "+":
			isMul = false
		default:
			log.Fatalf("Encountered unrecognized operation %q")
		}
		var result int
		if isMul {
			result = 1
		} else {
			result = 0
		}

		for j := 0; j < len(tokens)-1; j++ {
			val, err := strconv.Atoi(tokens[j][i])
			if err != nil {
				log.Fatalf("Encountered error parsing operand: %v", err)
			}
			if isMul {
				result *= val
			} else {
				result += val
			}
		}

		total += result
	}

	fmt.Println(total)
}
