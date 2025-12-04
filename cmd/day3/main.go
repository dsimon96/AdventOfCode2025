package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	part := flag.Int("part", 1, "which part to run")
	flag.Parse()

	var numDigits int
	switch *part {
	case 1:
		numDigits = 2
	case 2:
		numDigits = 12
	default:
		log.Fatalf("Invalid part %d", *part)
	}

	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]int, len(line))
		for i, c := range line {
			if !unicode.IsDigit(c) {
				log.Fatalf("Encountered non-numeric character %q in input", c)
			}
			digits[i] = int(c - '0')
		}

		joltage := 0
		prevIdx := -1
		for i := 0; i < numDigits; i++ {
			curIdx := prevIdx + 1
			val := digits[curIdx]
			for j := curIdx + 1; j < len(line)-numDigits+1+i; j++ {
				if digits[j] > val {
					curIdx = j
					val = digits[j]
				}
			}

			prevIdx = curIdx
			joltage = joltage*10 + val
		}

		total += joltage
	}

	fmt.Println(total)
}
