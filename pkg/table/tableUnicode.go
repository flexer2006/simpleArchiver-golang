// Package table provides functionality for encoding characters into binary strings
// using a predefined encoding table. This is useful for compression or custom encoding schemes.
package table

// EncodingTable maps runes (characters) to their corresponding binary string representations.
// The keys are Unicode characters, and the values are binary strings of varying lengths.
type EncodingTable map[rune]string

// BuildEncodingTable initializes and returns a predefined EncodingTable.
// The table includes mappings for:
//   - Basic lowercase letters (e.g., 'e', 't', 'a')
//   - Digits (0-9)
//   - Special characters (e.g., ' ', '.', ',', '!')
//   - Uppercase markers and additional letters (e.g., 'd', 'l', 'c')
//
// Example:
//
//	'e' -> "000"
//	't' -> "0010"
//	' ' -> "1010010"
//
// This table is designed for efficient encoding of common characters in English text.
func BuildEncodingTable() EncodingTable {
	return EncodingTable{
		// Basic letters
		'e': "000",
		't': "0010",
		'a': "0011",
		'o': "0100",
		'n': "0101",
		's': "0110",
		'r': "0111",
		'h': "10000",
		'i': "10001",

		// Digits
		'0': "1001000",
		'1': "1001001",
		'2': "1001010",
		'3': "1001011",
		'4': "1001100",
		'5': "1001101",
		'6': "1001110",
		'7': "1001111",
		'8': "1010000",
		'9': "1010001",

		// Special characters
		' ': "1010010",
		'.': "1010011",
		',': "1010100",
		'!': "1010101",
		'?': "1010110",
		'-': "1010111",
		'_': "1011000",
		'@': "1011001",
		'#': "1011010",
		'$': "1011011",
		'%': "1011100",
		'^': "1011101",
		'&': "1011110",
		'*': "1011111",
		'(': "1100000",
		')': "1100001",

		// Uppercase markers and letters
		'd': "1100010",
		'l': "1100011",
		'c': "1100100",
		'u': "1100101",
		'm': "1100110",
		'w': "1100111",
		'f': "1101000",
		'g': "1101001",
		'y': "1101010",
		'p': "1101011",
		'b': "1101100",
		'v': "1101101",
		'k': "1101110",
		'j': "1101111",
		'x': "1110000",
		'q': "1110001",
		'z': "1110010",
	}
}
