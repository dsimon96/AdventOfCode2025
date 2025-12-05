package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type Range struct {
	lower int
	upper int
}

func main() {
	part := flag.Int("part", 1, "which part to run")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	ranges := make([]Range, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		var lower int
		var upper int
		_, err := fmt.Sscanf(line, "%d-%d", &lower, &upper)
		if err != nil {
			log.Fatalf("Encountered an error parsing range: %v", err)
		}

		ranges = append(ranges, Range{lower, upper})
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if n := cmp.Compare(a.lower, b.lower); n != 0 {
			return n
		}
		return cmp.Compare(a.upper, b.upper)
	})

	switch *part {
	case 1:
		ids := make([]int, 0)
		for scanner.Scan() {
			line := scanner.Text()
			id, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Encountered an error parsing ID: %v", err)
			}

			ids = append(ids, id)
		}
		slices.Sort(ids)

		nextRangeIdx := 0
		fresh := 0
		for _, id := range ids {
			// find the first range that could contain this id
			for nextRangeIdx < len(ranges) && ranges[nextRangeIdx].upper < id {
				nextRangeIdx++
			}

			if nextRangeIdx == len(ranges) {
				// stop early, cannot have any more fresh
				break
			}

			if ranges[nextRangeIdx].lower <= id && id <= ranges[nextRangeIdx].upper {
				fresh++
			}

		}

		fmt.Println(fresh)
	case 2:
		if len(ranges) == 0 {
			fmt.Println(0)
			return
		}
		prev := ranges[0]
		total := 0
		for i := 1; i < len(ranges); i++ {
			if prev.upper < ranges[i].lower {
				// no overlap - add entire range
				total += prev.upper + 1 - prev.lower
				prev = ranges[i]
			} else if prev.upper < ranges[i].upper {
				// partial overlap - add the portion of prev range which doesn't overlap
				total += ranges[i].lower - prev.lower
				prev = ranges[i]
			}
			// prev completely covers the new range - ignore the new range
		}
		total += prev.upper + 1 - prev.lower

		fmt.Println(total)
	default:
		log.Fatalf("Invalid part %d", *part)
	}

}
