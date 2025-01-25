package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/flexer2006/simpleArchiver-golang/internal/application"
	"github.com/flexer2006/simpleArchiver-golang/pkg/table"
)

type BinaryChunks []BinaryChunk
type BinaryChunk string
type encodingTable = table.EncodingTable
type HexChunks []HexChunk
type HexChunk string

const chunksSize = int(8)

// Encode encodes a string into binary representation split into chunks.
func Encode(str string) string {
	str = prepareText(str)

	chunks := splitByChunks(encodeBinary(str), chunksSize)

	return chunks.ToHex().ToString()
}

func (hcs HexChunks) ToString() string {

	const sep = " "

	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, hc := range hcs[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(hc))
	}
	return buf.String()
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		hexChunk := chunk.ToHex()
		res = append(res, hexChunk)
	}

	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	var result HexChunk

	application.HandleError(func() error {
		value, err := strconv.ParseUint(string(bc), 2, chunksSize)
		if err == nil {

			result = HexChunk(strings.ToUpper(fmt.Sprintf("%X", value)))
		}
		return err
	})
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

// splitByChunks splits a binary string into chunks of the specified size.
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunkCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunkCount++
	}
	res := make(BinaryChunks, 0, chunkCount)

	var buf strings.Builder
	for i, ch := range bStr {
		buf.WriteRune(ch)
		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-buf.Len())
		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}

// prepareText prepares the input string by handling uppercase letters.
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!') // Add an exclamation mark for uppercase
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

// encodeBinary converts the prepared string into its binary representation.
func encodeBinary(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}
	return buf.String()
}

// bin retrieves the binary representation of a rune from the encoding table.
func bin(ch rune) string {
	tableUnicode := getEncodingTable()

	res, ok := tableUnicode[ch]
	if !ok {
		panic("character not found in encoding table: " + string(ch))
	}
	return res
}

// getEncodingTable loads the encoding table.
func getEncodingTable() encodingTable {
	return table.BuildEncodingTable()
}
