package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInt(line string) (parsed int) {
	parsed, _ = strconv.Atoi(strings.TrimSpace(line))
	return parsed
}

func solve(n int) (solved bool, fives int, threes int) {
	fives = n

	remainder := fives % 3

	if remainder == 1 {
		fives -= 10
		threes += 10
	}

	if remainder == 2 {
		fives -= 5
		threes += 5
	}

	if fives < 0 {
		solved = false
		return
	} else {
		solved = true
		return
	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	// grab the count
	_, _ = reader.ReadString('\n')

	for {
		line, err := reader.ReadString('\n')
		n := parseInt(line)
		solved, fives, threes := solve(n)

		if solved {
			for i := 0; i < fives; i++ {
				writer.WriteString("5")
			}
			for i := 0; i < threes; i++ {
				writer.WriteString("3")
			}
		} else {
			writer.WriteString("-1")
		}
		writer.WriteString("\n")
		writer.Flush()

		if err != nil {
			break
		}
	}
}
