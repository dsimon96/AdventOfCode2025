package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func countDigits(n int) int {
	d := 0
	for n > 0 {
		n /= 10
		d++
	}
	return d
}

func pow10(n int) int {
	p := 1
	for i := 0; i < n; i++ {
		p *= 10
	}
	return p
}

func main() {
	part := flag.Int("part", 1, "which part to run")
	flag.Parse()

	if *part != 1 && *part != 2 {
		log.Fatalf("Invalid part %d", *part)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	total := 0
	for _, p := range strings.Split(line, ",") {
		var lower int
		var upper int

		_, err := fmt.Sscanf(p, "%d-%d", &lower, &upper)
		if err != nil {
			log.Fatalf("Failed to parse %q as range (%v)", p, err)
		}

		digits := countDigits(lower)
		nextPow := pow10(digits)
		maxPatternLen := digits / 2
		for i := lower; i <= upper; i++ {
			if i == nextPow {
				digits++
				nextPow *= 10
				maxPatternLen = digits / 2
			}

			isInvalid := false
			if *part == 1 {
				if digits%2 == 1 {
					continue
				}
				divisor := pow10(digits / 2)
				lower := i % divisor
				upper := i / divisor

				if lower == upper {
					isInvalid = true
				}
			} else {
				var minPatternLen int
				if *part == 1 {
					minPatternLen = maxPatternLen
				} else {
					minPatternLen = 1
				}

				for patternLen := minPatternLen; patternLen <= maxPatternLen; patternLen++ {
					if digits%patternLen != 0 {
						continue
					}
					repetitions := digits / patternLen

					divisor := pow10(patternLen)

					pattern := i % divisor
					rem := i
					patternRepeats := true
					for j := 1; j < repetitions; j++ {
						rem /= divisor
						if rem%divisor != pattern {
							patternRepeats = false
							break
						}
					}

					if patternRepeats {
						isInvalid = true
						break
					}
				}

			}
			if isInvalid {
				total += i
			}
		}
	}

	fmt.Println(total)
}
