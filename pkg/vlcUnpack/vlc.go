// Package vlcUnpack provides functionality for unpacking files encoded with variable-length code (VLC).
// It reads a `.vlc` file, decodes its contents, and writes the decoded text to a new `.txt` file.
package vlcUnpack

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/flexer2006/simpleArchiver-golang/internal/application"
	"github.com/flexer2006/simpleArchiver-golang/pkg/chunks"
	"github.com/flexer2006/simpleArchiver-golang/pkg/decodingTree"
	"github.com/flexer2006/simpleArchiver-golang/pkg/table"
	"github.com/spf13/cobra"
)

const (
	// unpackedExtension is the file extension used for unpacked files.
	unpackedExtension = "txt"
)

// VlcUnpackCmd is the Cobra command for unpacking files encoded with variable-length code.
// Usage: vlcUnpack [file_path]
// Short: Unpack file using variable-length code.
var VlcUnpackCmd = &cobra.Command{
	Use:   "vlcUnpack [file_path]",
	Short: "Unpack file using variable-length code",
	Run: func(cmd *cobra.Command, args []string) {
		application.HandlePanic(func() {
			err := application.HandleError(func() error {
				if len(args) == 0 || args[0] == "" {
					return application.ErrEmptyPath
				}
				return unpack(args[0])
			})
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		})
	},
}

// unpack reads the file at the given path, decodes its contents using variable-length code,
// and writes the decoded text to a new file with a `.txt` extension.
// Returns an error if any step fails.
func unpack(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("Warning: failed to close file: %v", closeErr)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	decoded, err := Decode(string(data))
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	outputPath := generateOutputPath(filePath)
	if err := os.WriteFile(outputPath, []byte(decoded), 0644); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	log.Printf("File successfully unpacked: %s", outputPath)
	return nil
}

// generateOutputPath generates the output file path by replacing the original file's
// extension with `.txt`.
func generateOutputPath(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + "." + unpackedExtension
}

// Decode converts a space-separated hexadecimal string into its original text form.
// It validates the hex chunks, converts them to binary, and uses a decoding tree
// to reconstruct the original text.
//
// Parameters:
//   - encodedData: The space-separated hexadecimal string to decode.
//
// Returns:
//   - string: The decoded text.
//   - error: An error if decoding fails (e.g., invalid hex chunks or decoding tree issues).
func Decode(encodedData string) (string, error) {
	if encodedData == "" {
		return "", nil
	}

	// Parse the hex chunks and validate them
	hexChunks, err := chunks.NewHexChunks(encodedData)
	if err != nil {
		return "", fmt.Errorf("parse hex chunks: %w", err)
	}

	// Convert hex chunks to binary
	binaryChunks, err := hexChunks.ToBinary()
	if err != nil {
		return "", fmt.Errorf("convert hex to binary: %w", err)
	}

	// Join binary chunks into a single binary string
	binaryData := binaryChunks.Join()

	// Build the decoding tree from the encoding table
	encodingTable := table.BuildEncodingTable()
	tree, err := decodingTree.BuildDecodingTree(encodingTable)
	if err != nil {
		return "", fmt.Errorf("build decoding tree: %w", err)
	}

	// Decode the binary data using the tree
	decoded, err := tree.Decode(binaryData)
	if err != nil {
		return "", fmt.Errorf("decode binary data: %w", err)
	}

	// Restore the original case of the text
	return restoreCase(decoded), nil
}

// restoreCase processes the decoded text to restore uppercase letters.
// Uppercase letters are prefixed with '!' in the encoded data, so this function
// converts the next character to uppercase when '!' is encountered.
//
// Parameters:
//   - str: The decoded text to process.
//
// Returns:
//   - string: The text with uppercase letters restored.
func restoreCase(str string) string {
	var buf strings.Builder
	var capitalizeNext bool

	for _, r := range str {
		if capitalizeNext && unicode.IsLetter(r) {
			buf.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			buf.WriteRune(r)
		}

		if r == '!' {
			capitalizeNext = true
		}
	}

	return buf.String()
}

// init registers the VlcUnpackCmd with the root command during package initialization.
func init() {
	application.HandlePanic(func() {
		application.RootCmd.AddCommand(VlcUnpackCmd)
	})
}
