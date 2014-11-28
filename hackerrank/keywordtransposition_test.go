package main

import "testing"

func TestRemoveDuplicateRunes(t *testing.T) {
	var output []rune

	// leave non-duplicates alone
	output = RemoveDuplicateRunes([]rune("ABC"))
	ExpectRunes(t, "ABC", output)

	// remove multiple duplicates
	output = RemoveDuplicateRunes([]rune("ABCAB"))
	ExpectRunes(t, "ABC", output)
}

func TestRemoveRune(t *testing.T) {
	var removed bool
	var output []rune

	removed, output = RemoveRune([]rune("ABC"), 'A')

	if removed == false {
		t.Error("should have been removed")
	}

	ExpectRunes(t, "BC", output)
}

func TestRemoveRuneAt(t *testing.T) {
	var output []rune

	// at the start
	output = RemoveRuneAt([]rune("ABCDE"), 0)
	ExpectRunes(t, "BCDE", output)

	// at the end
	output = RemoveRuneAt([]rune("ABCDE"), 4)
	ExpectRunes(t, "ABCD", output)

	// in the middle
	output = RemoveRuneAt([]rune("ABCDE"), 3)
	ExpectRunes(t, "ABCE", output)
}

func TestOrderRunes(t *testing.T) {
	var output []int

	// already ordered
	output = OrderRunes([]rune("ABCDE"))
	ExpectInts(t, []int{0, 1, 2, 3, 4}, output)

	// reverse ordered
	output = OrderRunes([]rune("EDCBA"))
	ExpectInts(t, []int{4, 3, 2, 1, 0}, output)

	// random ordered
	output = OrderRunes([]rune("SPORT"))
	ExpectInts(t, []int{2, 1, 3, 0, 4}, output)
}

func TestCreateCypher(t *testing.T) {
	var cypher map[rune]rune

	cypher = CreateCypher("SPORT")

	// SPORT   PFAKV
	// ABCDE   QGBLW
	// FGHIJ   RHCMX
	// KLMNQ   SIDNY
	// UVWXY   TJEOZ
	// Z       U

	// cypher
	// A: Q
	// B: G
	// C: B
	// D: L
	// E: E
	// F: R
	// G: H
	// H: C
	// I: M
	// J: X
	// K: S
	// L: I
	// M: D
	// N: N
	// O: A
	// P: F
	// Q: Y
	// R: K
	// S: P
	// T: V
	// U: T
	// V: J
	// W: E
	// X: O
	// Y: Z
	// Z: U

	if cypher['A'] != 'Q' {
		t.Errorf("A isn't Q, it's %v", string(cypher['A']))
	}

	if cypher['B'] != 'G' {
		t.Errorf("B isn't G, it's %v", string(cypher['B']))
	}

	if cypher['O'] != 'A' {
		t.Errorf("O isn't A, it's %v", string(cypher['O']))
	}

	if cypher['Z'] != 'U' {
		t.Errorf("Z isn't U, it's %v", string(cypher['Z']))
	}
}

func ExpectRunes(t *testing.T, expected string, received []rune) {
	if expected != string(received) {
		t.Errorf("Got %v, expected %v", string(received), expected)
	}
}

func ExpectInts(t *testing.T, expected []int, received []int) {
	if len(expected) != len(received) {
		t.Errorf("lengths are not the same, %v, %v", expected, received)
		return
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != received[i] {
			t.Errorf("Got %v, expected %v", received, expected)
			return
		}
	}
}
