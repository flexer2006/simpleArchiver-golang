package vlcUnpack

import (
	"github.com/flexer2006/simpleArchiver-golang/pkg/decodingTree"
	"testing"
)

func TestDecodeInvalidHexChunks(t *testing.T) {
	invalidHexData := "xyz123"
	_, err := Decode(invalidHexData)
	if err == nil {
		t.Errorf("Expected an error for invalid hex data, but got nil")
	}
}

func TestDecodeEmptyData(t *testing.T) {
	result, err := Decode("")
	if err != nil {
		t.Errorf("Expected no error for empty data, but got %v", err)
	}
	if result != "" {
		t.Errorf("Expected empty result for empty data, but got %v", result)
	}
}

func TestBuildDecodingTree(t *testing.T) {
	encodingTable := map[rune]string{
		'a': "00",
		'b': "01",
		'c': "10",
	}

	tree, err := decodingTree.BuildDecodingTree(encodingTable)
	if err != nil {
		t.Errorf("BuildDecodingTree() failed: %v", err)
	}

	decoded, err := tree.Decode("000110")
	if err != nil {
		t.Errorf("Decode() failed: %v", err)
	}
	if decoded != "abc" {
		t.Errorf("Decode() result = %v, want abc", decoded)
	}
}
