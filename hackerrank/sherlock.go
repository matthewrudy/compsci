package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func writeSolution(writer *bufio.Writer, fives int, threes int) {
	for i := 0; i < fives; i++ {
		writer.WriteString("5")
	}

	for i := 0; i < threes; i++ {
		writer.WriteString("3")
	}
	writer.WriteString("\n")
	writer.Flush()
}

func parseInt(line string) (parsed int) {
	parsed, _ = strconv.Atoi(line)
	return parsed
}

func readLine(reader *bufio.Reader) (line string) {
	line, _ = reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInt(reader *bufio.Reader) (read int) {
	line := readLine(reader)
	return parseInt(line)
}

func writeLine(writer *bufio.Writer, line string) {
	writer.WriteString(line)
	writer.WriteString("\n")
}

func writeInt(writer *bufio.Writer, number int) {
	line := strconv.Itoa(number)
	writeLine(writer, line)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	count := readInt(reader)

	for i := 0; i < count; i++ {
		n := readInt(reader)
		solved, fives, threes := solve(n)

		if solved {
			writeSolution(writer, fives, threes)
		} else {
			writeLine(writer, "-1")
		}
	}
}
