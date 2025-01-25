// Package chunks provides utilities for splitting binary data into fixed-size chunks
// and converting between binary and hexadecimal representations. It ensures consistent
// chunk sizes for reliable encoding/decoding operations.
package chunks

import (
	"fmt"
	"strconv"
	"strings"
)

// BinaryChunks represents a sequence of fixed-size binary chunks.
// Each chunk is guaranteed to be exactly ChunkSize bits long (8 by default).
type BinaryChunks []BinaryChunk

// BinaryChunk is an 8-bit binary string representation (e.g., "00110101").
// Implements validation to ensure proper format during conversions.
type BinaryChunk string

// HexChunks represents a sequence of hexadecimal-encoded chunks.
// Each chunk is a 2-character string representing 8 bits of data.
type HexChunks []HexChunk

// HexChunk is a 2-character hexadecimal string representation (e.g., "A5").
// Validates format and value range during conversions.
type HexChunk string

const (
	// ChunkSize defines the fixed bit-length for all binary chunks
	ChunkSize    = 8
	hexChunkSep  = " "
	hexChunkSize = 2 // Hex representation size for 8 bits (2 hex characters)
)

// ToString joins HexChunks into a space-separated string.
// Example: HexChunks{"A1", "FF"} => "A1 FF".
func (hcs HexChunks) ToString() string {
	if len(hcs) == 0 {
		return ""
	}

	// Convert HexChunks to []string
	strChunks := make([]string, len(hcs))
	for i, chunk := range hcs {
		strChunks[i] = string(chunk)
	}

	return strings.Join(strChunks, hexChunkSep)
}

// ToBinary converts HexChunks to BinaryChunks with validation.
// Returns error for invalid hex values or incorrect chunk sizes.
func (hcs HexChunks) ToBinary() (BinaryChunks, error) {
	res := make(BinaryChunks, 0, len(hcs))
	for _, chunk := range hcs {
		bc, err := chunk.ToBinary()
		if err != nil {
			return nil, fmt.Errorf("conversion failed: %w", err)
		}
		res = append(res, bc)
	}
	return res, nil
}

// ToBinary converts a HexChunk to BinaryChunk with validation.
// Ensures proper hex format and 8-bit value range.
func (hc HexChunk) ToBinary() (BinaryChunk, error) {
	if len(hc) != hexChunkSize {
		return "", fmt.Errorf("invalid hex chunk size: want %d, got %d", hexChunkSize, len(hc))
	}

	value, err := strconv.ParseUint(string(hc), 16, ChunkSize)
	if err != nil {
		return "", fmt.Errorf("invalid hex value %q: %w", hc, err)
	}

	return BinaryChunk(fmt.Sprintf("%08b", value)), nil
}

// ToHex converts BinaryChunks to validated HexChunks.
// Returns error for invalid binary formats or conversion issues.
func (bcs BinaryChunks) ToHex() (HexChunks, error) {
	res := make(HexChunks, 0, len(bcs))
	for _, chunk := range bcs {
		hc, err := chunk.ToHex()
		if err != nil {
			return nil, fmt.Errorf("conversion failed: %w", err)
		}
		res = append(res, hc)
	}
	return res, nil
}

// ToHex converts BinaryChunk to HexChunk with validation.
// Ensures 8-bit length and proper binary format (0s and 1s).
func (bc BinaryChunk) ToHex() (HexChunk, error) {
	if len(bc) != ChunkSize {
		return "", fmt.Errorf("invalid binary chunk size: want %d, got %d", ChunkSize, len(bc))
	}

	if strings.TrimFunc(string(bc), func(r rune) bool {
		return r == '0' || r == '1'
	}) != "" {
		return "", fmt.Errorf("invalid binary characters in %q", bc)
	}

	value, err := strconv.ParseUint(string(bc), 2, ChunkSize)
	if err != nil {
		return "", fmt.Errorf("conversion error: %w", err)
	}

	return HexChunk(fmt.Sprintf("%02X", value)), nil
}

// SplitByChunks splits a binary string into 8-bit chunks with validation and padding.
// - Validates input contains only 0s and 1s
// - Pads with trailing zeros to reach multiples of ChunkSize
// Returns error for invalid characters.
func SplitByChunks(bStr string) (BinaryChunks, error) {
	// Validate binary format
	if strings.TrimFunc(bStr, func(r rune) bool { return r == '0' || r == '1' }) != "" {
		return nil, fmt.Errorf("invalid binary string: contains non-binary characters")
	}

	// Calculate padding needed
	strLen := len(bStr)
	padding := (ChunkSize - (strLen % ChunkSize)) % ChunkSize
	paddedStr := bStr + strings.Repeat("0", padding)

	// Split into chunks
	chunkCount := (len(paddedStr) + ChunkSize - 1) / ChunkSize
	res := make(BinaryChunks, chunkCount)
	for i := 0; i < chunkCount; i++ {
		start := i * ChunkSize
		end := start + ChunkSize
		if end > len(paddedStr) {
			end = len(paddedStr)
		}
		res[i] = BinaryChunk(paddedStr[start:end])
	}

	return res, nil
}

// Join concatenates BinaryChunks into a single binary string.
// Preserves original chunk order and padding.
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder
	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}
	return buf.String()
}

// NewHexChunks creates validated HexChunks from a space-separated string.
// Validates each chunk is a 2-character hex value.
// Returns error for invalid format or values.
func NewHexChunks(str string) (HexChunks, error) {
	if str == "" {
		return HexChunks{}, nil
	}

	chunks := strings.Fields(str)
	hexChunks := make(HexChunks, 0, len(chunks))

	for _, chunk := range chunks {
		if len(chunk) != hexChunkSize {
			return nil, fmt.Errorf("invalid hex chunk %q: wrong size", chunk)
		}
		if _, err := strconv.ParseUint(chunk, 16, ChunkSize); err != nil {
			return nil, fmt.Errorf("invalid hex value %q: %w", chunk, err)
		}
		hexChunks = append(hexChunks, HexChunk(chunk))
	}

	return hexChunks, nil
}
