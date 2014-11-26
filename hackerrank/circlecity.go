package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func countEdges(r int) (count int) {
	// check in just one quadrant

	var i int
	var j int
	var j2 int

	for i = 0; i*i < r; i++ {
		j2 = r - (i * i)
		j = int(math.Sqrt(float64(j2)))

		if i*i+j*j == r {
			count += 4
		}
	}

	return
}

func solve(r int, k int) (possible bool) {
	return k >= countEdges(r)
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
		line := readLine(reader)
		ints := strings.Split(line, " ")

		r := parseInt(ints[0])
		k := parseInt(ints[1])

		possible := solve(r, k)

		if possible {
			writeLine(writer, "possible")
		} else {
			writeLine(writer, "impossible")
		}
	}
}
