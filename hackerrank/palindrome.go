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

func main() {

	reader := bufio.NewReader(os.Stdin)

	// grab the count
	_, _ = reader.ReadString('\n')

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		index := solve(line)

		fmt.Println(index)

		if err != nil {
			break
		}
	}
}
