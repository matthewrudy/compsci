package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func readInts(reader *bufio.Reader) []int {
	line := readLine(reader)
	chars := strings.Split(line, " ")

	ints := make([]int, len(chars))
	for i, n := range chars {
		ints[i] = parseInt(n)
	}
	return ints
}

func readIntPair(reader *bufio.Reader) (a int, b int) {
	ints := readInts(reader)
	a = ints[0]
	b = ints[1]

	return
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	n, count := readIntPair(reader)

	countries := make([]int, n)

	for i := 0; i < n; i++ {
		countries[i] = i
	}

	for i := 0; i < count; i++ {
		a, b := readIntPair(reader)

		aCountry := countries[a]
		bCountry := countries[b]

		for j := 0; j < n; j++ {
			if countries[j] == bCountry {
				countries[j] = aCountry
			}
		}
	}

	countryCounts := make([]int, n)

	for i := 0; i < n; i++ {
		country := countries[i]
		countryCounts[country] += 1
	}

	result := 0
	rem := n
	var this int

	for i := 0; i < n-1; i++ {
		this = countryCounts[i]
		if this > 0 {
			rem -= this
			result += this * rem
		}
	}

	writeInt(writer, result)
}
