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

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	count := readInt(reader)

	for i := 0; i < count; i++ {
		keyword := readLine(reader)
		text := readLine(reader)

		cypher := createCypher(keyword)
		text = translateText(text, cypher)

		writeLine(writer, text)
	}
}

func createCypher(keyword string) (cypher map[rune]rune) {
	allGlyphs := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	cypher = make(map[rune]rune, 26)

	for i, a := range allGlyphs {
		cypher[a] = allGlyphs[25-i]
	}

	return
}

func translateText(text string, cypher map[rune]rune) (translated string) {
	runes := make([]rune, len(text))

	var nilRune rune

	for i, rune := range text {

		trune := cypher[rune]

		if trune != nilRune {
			runes[i] = trune
		} else {
			runes[i] = rune
		}
	}
	return string(runes)
}
