package morse

import (
	"errors"
)

// Symbol type for morse symbol
type Symbol uint8

// List of morse symbol
const (
	Ti Symbol = iota
	Ta
)

// Morse Alphabet
var (
	Alphabet = map[rune][]Symbol{
		'A': {Ti, Ta},
		'B': {Ta, Ti, Ti, Ti},
		'C': {Ta, Ti, Ta, Ti},
		'D': {Ta, Ti, Ti},
		'E': {Ti},
		'F': {Ti, Ti, Ta, Ti},
		'G': {Ta, Ta, Ti},
		'H': {Ti, Ti, Ti, Ti},
		'I': {Ti, Ti},
		'J': {Ti, Ta, Ta, Ta},
		'K': {Ta, Ti, Ta},
		'L': {Ti, Ta, Ti, Ti},
		'M': {Ta, Ta},
		'N': {Ta, Ti},
		'O': {Ta, Ta, Ta},
		'P': {Ti, Ta, Ta, Ti},
		'Q': {Ta, Ta, Ti, Ta},
		'R': {Ti, Ta, Ti},
		'S': {Ti, Ti, Ti},
		'T': {Ta},
		'U': {Ti, Ti, Ta},
		'V': {Ti, Ti, Ti, Ta},
		'W': {Ti, Ta, Ta},
		'X': {Ta, Ti, Ti, Ta},
		'Y': {Ta, Ti, Ta, Ta},
		'Z': {Ta, Ta, Ti, Ti},
		'1': {Ti, Ta, Ta, Ta, Ta},
		'2': {Ti, Ti, Ta, Ta, Ta},
		'3': {Ti, Ti, Ti, Ta, Ta},
		'4': {Ti, Ti, Ti, Ti, Ta},
		'5': {Ti, Ti, Ti, Ti, Ti},
		'6': {Ta, Ti, Ti, Ti, Ti},
		'7': {Ta, Ta, Ti, Ti, Ti},
		'8': {Ta, Ta, Ta, Ti, Ti},
		'9': {Ta, Ta, Ta, Ta, Ti},
		'0': {Ta, Ta, Ta, Ta, Ta},
	}

	errInvalidChar = errors.New("input contains invalid char")
)

// CheckInput returns error if a rune does not belong to morse alphabet
func CheckInput(r []rune) error {
	for i := range r {
		if _, ok := Alphabet[r[i]]; !ok && r[i] != ' ' {
			return errInvalidChar
		}
	}
	return nil
}
