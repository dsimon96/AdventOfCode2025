package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	part := flag.Int("part", 1, "which part to run")
	flag.Parse()

	if *part != 1 && *part != 2 {
		log.Fatalf("Invalid part %d", *part)
	}

	pos := 50
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var dir rune
		var n int

		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%c%d", &dir, &n)

		if err != nil {
			log.Fatalf("Failed to parse line %q (%v)", line, err)
		}

		if dir != 'L' && dir != 'R' {
			log.Fatalf("Invalid direction %q", dir)
		}

		fullRotations := n / 100
		n %= 100

		if *part == 2 {
			count += fullRotations
			if dir == 'L' && pos > 0 && n >= pos {
				count += 1
			} else if dir == 'R' && pos+n >= 100 {
				count += 1
			}
		}

		if dir == 'L' {
			n = 100 - n
		}
		pos = (pos + n) % 100

		if *part == 1 && pos == 0 {
			count += 1
		}
	}

	fmt.Println(count)
}
