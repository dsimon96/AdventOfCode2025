package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func isAccessible(grid [][]bool, rows int, cols int, r int, c int) bool {
	adjacent := 0
	for nr := r - 1; nr <= r+1; nr++ {
		for nc := c - 1; nc <= c+1; nc++ {
			// count if (nr, nc) is inbounds, is not the same as (r, c), and is occupied
			if 0 <= nr && nr < rows && 0 <= nc && nc < cols && !(nr == r && nc == c) && grid[nr][nc] {
				adjacent += 1
			}
		}
	}

	return adjacent < 4
}

func main() {
	part := flag.Int("part", 1, "which part to run")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]bool
	rows := 0
	cols := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			log.Fatal("Encountered empty line")
		}

		if cols == 0 {
			cols = len(line)
		} else if len(line) != cols {
			log.Fatalf("Encountered a line with length %d (expected %d)", len(line), cols)
		}

		row := make([]bool, cols)
		for i := 0; i < cols; i++ {
			switch line[i] {
			case '.':
				row[i] = false
			case '@':
				row[i] = true
			default:
				log.Fatalf("Encountered invalid character %q in input", line[i])
			}
		}

		grid = append(grid, row)
		rows++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	switch *part {
	case 1:
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if !grid[r][c] {
					continue
				}

				if isAccessible(grid, rows, cols, r, c) {
					count++
				}
			}
		}
	case 2:
		for {
			removed := false
			for r := 0; r < rows; r++ {
				for c := 0; c < cols; c++ {
					if !grid[r][c] {
						continue
					}

					if isAccessible(grid, rows, cols, r, c) {
						removed = true
						grid[r][c] = false
						count++
					}
				}
			}

			if !removed {
				break
			}
		}
	default:
		log.Fatalf("Invalid part %d", *part)
	}

	fmt.Println(count)
}
