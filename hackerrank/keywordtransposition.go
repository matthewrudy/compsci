package main

import (
	"bufio"
	"fmt"
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

func removeDuplicateRunes(runes []rune) (deduped []rune) {
	var j int
	var changed bool

	deduped = runes

	for i := 0; i < len(deduped)-2; i++ {
		for {
			changed = false

			for j = i + 1; j < len(deduped)-1; j++ {
				if deduped[i] == deduped[j] {
					deduped = removeRuneAt(deduped, j)
					changed = true
					break
				}
			}

			if changed == false {
				break
			}
		}
	}
	return
}

// return a new slice exluding the found rune
func removeRune(array []rune, target rune) (removed bool, without []rune) {
	for i, found := range array {
		if found == target {
			removed = true
			without = removeRuneAt(array, i)
			return
		}
	}

	removed = false
	return
}

func removeRuneAt(array []rune, i int) (without []rune) {
	if i >= len(array) {
		panic("index beyond array")
	}

	if i < 0 {
		panic("negative index")
	}

	front := array[0 : i-1]

	if len(array) == i+1 {
		return front
	}

	back := array[i+1:]
	return makeSlice(front, back)
}

// take two slices and join them together
func makeSlice(front []rune, back []rune) (joined []rune) {
	joined = make([]rune, len(front)+len(back))

	i := 0

	// copy the front in the front
	for _, rune := range front {
		joined[i] = rune
		i++
	}

	// then the back to the ...
	for _, rune := range back {
		joined[i] = rune
		i++
	}

	return
}

func createCypher(keyword string) (cypher map[rune]rune) {
	// iterators
	var i int
	var j int

	allRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	kwRunes := []rune(keyword)

	kwRunes = removeDuplicateRunes(kwRunes)

	// remove kwRunes from allRunes
	remainingRunes := allRunes

	for _, kwRune := range kwRunes {
		removed, slice := removeRune(remainingRunes, kwRune)

		if removed {
			remainingRunes = slice
		} //else {
		//	panic(fmt.Sprintf("rune %v is missing from %v, %v", string(kwRune), string(remainingRunes), string(kwRunes)))
		//}
	}

	// store the adjusted kwlength for later
	kwLength := len(kwRunes)

	// now we'll go through all the letters and put them
	// into arrays, one for each column
	matrix := make([][]rune, kwLength)

	// initialize the sub arrays
	for i, _ := range matrix {
		matrix[i] = make([]rune, 26/kwLength+1)

		// stick the keyword letters in place
		matrix[i][0] = kwRunes[i]
	}

	i = 0
	j = 1

	// stick the remaining letters in place
	for _, rune := range remainingRunes {
		matrix[i][j] = rune

		i++
		if i == kwLength {
			i = 0
			j++
		}
	}

	// now we can make our cypher
	cypher = make(map[rune]rune, 26)

	kwOrder := orderRunes(kwRunes)

	runeIndex := 0
	var rune rune

	for i, _ = range kwOrder {
		for j = 0; j < len(matrix[i]); j++ {
			rune = allRunes[runeIndex]
			cypher[rune] = matrix[i][j]
			runeIndex += 1

			if runeIndex == 26 {
				return
			}
		}
	}
	return
}

func orderRunes(runes []rune) (ordered []int) {
	ordered = make([]int, len(runes))

	var i int
	var r1 int
	var r2 int
	var changed bool

	// initialize
	for i, _ = range runes {
		ordered[i] = i
	}

	// sort
	for {
		changed = false
		for i = 1; i < len(runes); i++ {
			r1 = ordered[i-1]
			r2 = ordered[i]

			if runes[r1] > runes[r2] {
				ordered[i-1] = r2
				ordered[i] = r1

				changed = true
			}
		}

		if changed == false {
			return
		}
	}
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
