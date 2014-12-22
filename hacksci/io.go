package hacksci

import (
	"bufio"
	"strconv"
	"strings"
)

func ParseInt(line string) (parsed int) {
	parsed, _ = strconv.Atoi(line)
	return parsed
}

func ReadLine(reader *bufio.Reader) (line string) {
	line, _ = reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func ReadInt(reader *bufio.Reader) (read int) {
	line := ReadLine(reader)
	return ParseInt(line)
}

func ReadInts(reader *bufio.Reader) []int {
	line := ReadLine(reader)
	chars := strings.Split(line, " ")

	ints := make([]int, len(chars))
	for i, n := range chars {
		ints[i] = ParseInt(n)
	}
	return ints
}

func WriteLine(writer *bufio.Writer, line string) {
	writer.WriteString(line)
	writer.WriteString("\n")
}

func WriteInt(writer *bufio.Writer, number int) {
	line := strconv.Itoa(number)
	WriteLine(writer, line)
}

func WriteInts(writer *bufio.Writer, numbers []int) {
	for i, number := range numbers {
		string := strconv.Itoa(number)
		writer.WriteString(string)

		if i < len(numbers)-1 {
			writer.WriteString(" ")
		}
	}

	writer.WriteString("\n")
}
