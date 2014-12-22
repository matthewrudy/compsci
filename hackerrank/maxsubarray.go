package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// IO methods

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

func readInts(reader *bufio.Reader) []int {
	line := readLine(reader)
	chars := strings.Split(line, " ")

	ints := make([]int, len(chars))
	for i, n := range chars {
		ints[i] = parseInt(n)
	}
	return ints
}

func writeLine(writer *bufio.Writer, line string) {
	writer.WriteString(line)
	writer.WriteString("\n")
}

func writeInt(writer *bufio.Writer, number int) {
	line := strconv.Itoa(number)
	writeLine(writer, line)
}

func writeInts(writer *bufio.Writer, numbers []int) {
	for i, number := range numbers {
		string := strconv.Itoa(number)
		writer.WriteString(string)

		if i < len(numbers)-1 {
			writer.WriteString(" ")
		}
	}

	writer.WriteString("\n")
}

// main

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	count := readInt(reader)

	for i := 0; i < count; i++ {
		_ = readInt(reader)
		list := readInts(reader)

		contiguous := MaxContiguous(list)
		nonContiguous := MaxNonContiguous(list)

		writeInts(writer, []int{contiguous, nonContiguous})
	}
}

// contiguous

func MaxContiguous(list []int) int {
	maxSum := list[0]
	thisSum := list[0]

	for _, element := range list[1:] {
		if thisSum+element > element {
			thisSum += element
		} else {
			thisSum = element
		}

		if thisSum > maxSum {
			maxSum = thisSum
		}
	}

	return maxSum
}

// not-contiguous

func MaxNonContiguous(list []int) (sum int) {
	inserted := 0
	maxElement := list[0]

	for _, element := range list {
		if element > 0 {
			sum += element
			inserted += 1
		}

		if element > maxElement {
			maxElement = element
		}
	}

	// if no positive ints were foun
	if inserted == 0 {
		return maxElement
	}

	return
}
