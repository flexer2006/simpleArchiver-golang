// Package vlcPack provides functionality for encoding text into binary format
// using variable-length codes (VLC). It prepares the text, encodes it into binary,
// and converts the binary data into hexadecimal chunks for storage or transmission.
package vlcPack

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/flexer2006/simpleArchiver-golang/pkg/chunks"
	"github.com/flexer2006/simpleArchiver-golang/pkg/table"
)

// Encode takes a string, prepares it for encoding, converts it to binary using
// a predefined encoding table, and returns the result as a space-separated
// hexadecimal string.
//
// Parameters:
//   - str: The input string to encode.
//
// Returns:
//   - string: The encoded hexadecimal string.
//   - error: An error if encoding fails (e.g., due to an undefined character).
func Encode(str string) (string, error) {
	// Prepare the text by handling uppercase letters
	prepared := prepareText(str)

	// Encode the prepared text into binary
	encoded, err := encodeToBinary(prepared)
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}

	// Split the binary string into chunks and convert to hexadecimal
	binaryChunks, err := chunks.SplitByChunks(encoded)
	if err != nil {
		return "", fmt.Errorf("split binary into chunks: %w", err)
	}

	hexChunks, err := binaryChunks.ToHex()
	if err != nil {
		return "", fmt.Errorf("convert binary chunks to hex: %w", err)
	}

	// Join the hexadecimal chunks into a space-separated string
	return hexChunks.ToString(), nil
}

// prepareText processes the input string to handle uppercase letters.
// Uppercase letters are prefixed with '!' and converted to lowercase.
//
// Parameters:
//   - str: The input string to prepare.
//
// Returns:
//   - string: The processed string with uppercase letters handled.
func prepareText(str string) string {
	var buf strings.Builder
	for _, r := range str {
		if unicode.IsUpper(r) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// encodeToBinary converts a string into a binary string using a predefined
// encoding table. Each character in the string is replaced with its corresponding
// binary code from the table.
//
// Parameters:
//   - str: The input string to encode.
//
// Returns:
//   - string: The binary-encoded string.
//   - error: An error if a character in the string is not found in the encoding table.
func encodeToBinary(str string) (string, error) {
	table := table.BuildEncodingTable()
	var builder strings.Builder

	for _, r := range str {
		code, ok := table[r]
		if !ok {
			return "", fmt.Errorf("undefined character: %U", r)
		}
		builder.WriteString(code)
	}

	return builder.String(), nil
}
