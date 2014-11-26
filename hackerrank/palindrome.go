package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func checkPalindrome(word string, start int, end int) (success bool) {
	for {
		if word[start] != word[end] {
			return false
		}
		start++
		end--

		if start >= end {
			return true
		}
	}
}

func solve(word string) (index int) {
	length := len(word)

	var j int

	for i := 0; i < length/2; i++ {
		j = length - 1 - i

		if word[i] != word[j] {
			if checkPalindrome(word, i+1, j) {
				return i
			}

			if checkPalindrome(word, i, j-1) {
				return j
			}

			return -1
		}
	}
	return -1

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
	writer.Flush()
}

func writeInt(writer *bufio.Writer, number int) {
	line := strconv.Itoa(number)
	writeLine(writer, line)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	count := readInt(reader)

	for i := 0; i < count; i++ {
		line := readLine(reader)
		index := solve(line)
		writeInt(writer, index)
	}
}
