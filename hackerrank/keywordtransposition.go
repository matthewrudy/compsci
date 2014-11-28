package main

import (
	"bufio"
	"fmt"
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

		cypher := CreateCypher(keyword)
		text = translateText(text, cypher)

		writeLine(writer, text)
	}
}

// Rune Dedupe methods

func RemoveDuplicateRunes(runes []rune) (deduped []rune) {
	// ABBA
	// ij
	// i j
	// i  j <- remove at j
	// ABB
	// i  j
	// *j hit the end*
	// ABB
	//  ij <- remove at j
	// AB
	//  ij
	// *j hit the end*

	i := 0

	// the i loop
	for {

		if i >= len(runes)-1 {
			break
		}

		j := i + 1

		// the j loop
		for {

			if j >= len(runes) {
				break
			}

			if runes[i] == runes[j] {
				fmt.Print(i, ",", j, "==", runes, "\n")
				runes = RemoveRuneAt(runes, j)
				// keep j the same
			} else {
				j++
			}
		}

		i++
	}

	return runes
}

// return a new slice exluding the found rune
func RemoveRune(array []rune, target rune) (removed bool, without []rune) {
	for i, found := range array {
		if found == target {
			removed = true
			without = RemoveRuneAt(array, i)
			return
		}
	}

	removed = false
	return
}

func RemoveRuneAt(array []rune, i int) (without []rune) {
	if i >= len(array) {
		panic("index beyond array")
	}

	if i < 0 {
		panic("negative index")
	}

	// remove the first element
	if i == 0 {
		return array[1:]
	}

	// remove the last element
	if len(array) == i+1 {
		return array[0:i]
	}

	front := array[0:i]
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

func CreateCypher(keyword string) (cypher map[rune]rune) {
	// iterators
	var i int
	var j int

	allRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	kwRunes := []rune(keyword)

	// remove duplicates from keyword
	kwRunes = RemoveDuplicateRunes(kwRunes)

	// remove kwRunes from allRunes
	remainingRunes := allRunes

	for _, kwRune := range kwRunes {
		removed, slice := RemoveRune(remainingRunes, kwRune)

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

	kwOrder := OrderRunes(kwRunes)

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

func OrderRunes(runes []rune) (ordered []int) {
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
