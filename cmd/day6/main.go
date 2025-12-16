package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation int

const (
	OperationPlus Operation = iota
	OperationTimes
	OperationUnknown
)

func ParseOp(s byte) Operation {
	switch s {
	case '*':
		return OperationTimes
	case '+':
		return OperationPlus
	default:
		log.Fatalf("Encountered unrecognized operation %v", s)
		return OperationUnknown
	}
}

type Problem struct {
	op       Operation
	operands []int
}

func (p Problem) Eval() int {
	var result int
	switch p.op {
	case OperationPlus:
		result = 0
	case OperationTimes:
		result = 1
	}

	for _, operand := range p.operands {
		switch p.op {
		case OperationPlus:
			result += operand
		case OperationTimes:
			result *= operand
		}
	}

	return result
}

func main() {
	part := flag.Int("part", 1, "which part to run")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	problems := make([]Problem, 0)
	switch *part {
	case 1:
		tokens := make([][]string, 0)
		for _, line := range lines {
			fields := strings.Fields(line)
			tokens = append(tokens, fields)
		}

		for i := 0; i < len(tokens[0]); i++ {
			operator := tokens[len(tokens)-1][i]
			op := ParseOp(operator[0])

			var operands []int
			for j := 0; j < len(tokens)-1; j++ {
				val, err := strconv.Atoi(tokens[j][i])
				if err != nil {
					log.Fatalf("Encountered error parsing operand: %v", err)
				}

				operands = append(operands, val)
			}

			problems = append(problems, Problem{op, operands})
		}

	case 2:
		// determine bounds for safe col-based indexing
		minLength := len(lines[0])
		maxLength := len(lines[0])
		for i := 1; i < len(lines)-1; i++ {
			minLength = min(minLength, len(lines[i]))
			maxLength = max(maxLength, len(lines[i]))
		}

		// identify which columns are problem delimiters
		delimCols := make([]int, 0)
		for i := 0; i < minLength; i++ {
			colIsEmpty := true
			for j := 0; j < len(lines)-1; j++ {
				if lines[j][i] != ' ' {
					colIsEmpty = false
				}
			}
			if colIsEmpty {
				delimCols = append(delimCols, i)
			}
		}

		// determine column ranges for each problem
		type Range struct {
			start int
			end   int
		}
		ranges := make([]Range, 0)
		start := 0
		for _, end := range delimCols {
			ranges = append(ranges, Range{start, end})
			start = end + 1
		}
		ranges = append(ranges, Range{start, maxLength})

		// parse each problem
		for _, r := range ranges {
			op := ParseOp(lines[len(lines)-1][r.start])
			operands := make([]int, 0)
			for c := r.start; c < r.end; c++ {
				operand := 0
				for i := 0; i < len(lines)-1; i++ {
					if c >= len(lines[i]) || lines[i][c] == ' ' {
						continue
					}
					operand *= 10
					operand += int(lines[i][c] - '0')
				}
				operands = append(operands, operand)
			}
			problems = append(problems, Problem{op, operands})
		}
	default:
		log.Fatalf("Unrecognized part %q", *part)
	}

	total := 0
	for _, problem := range problems {
		total += problem.Eval()
	}

	fmt.Println(total)
}
