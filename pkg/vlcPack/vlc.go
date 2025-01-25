// Package vlcPack provides functionality for packing files using variable-length code (VLC) encoding.
// It includes a CLI command to encode a file and save the result with a `.vlc` extension.
package vlcPack

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/flexer2006/simpleArchiver-golang/internal/application"
	"github.com/spf13/cobra"
)

const (
	// packedExtension is the file extension used for packed files.
	packedExtension = "vlc"
)

// VlcPackCmd is the Cobra command for packing files using variable-length code.
// Usage: vlcPack [file_path]
// Short: Pack file using variable-length code.
var VlcPackCmd = &cobra.Command{
	Use:   "vlcPack [file_path]",
	Short: "Pack file using variable-length code",
	Run: func(cmd *cobra.Command, args []string) {
		application.HandlePanic(func() {
			if err := validateAndPack(args); err != nil {
				log.Fatalf("Error: %v", err)
			}
		})
	},
}

// validateAndPack validates the input arguments and initiates the packing process.
// Returns an error if the file path is empty or if packing fails.
func validateAndPack(args []string) error {
	return application.HandleError(func() error {
		if len(args) == 0 || args[0] == "" {
			return application.ErrEmptyPath
		}
		return pack(args[0])
	})
}

// pack reads the file at the given path, encodes its contents using variable-length code,
// and writes the encoded data to a new file with a `.vlc` extension.
// Returns an error if any step fails.
func pack(filePath string) error {
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

	encoded, err := Encode(string(data))
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	outputPath := generateOutputPath(filePath)
	if err := os.WriteFile(outputPath, []byte(encoded), 0644); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	log.Printf("File successfully packed: %s", outputPath)
	return nil
}

// generateOutputPath generates the output file path by replacing the original file's
// extension with `.vlc`.
func generateOutputPath(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base)) + "." + packedExtension
}

// init registers the VlcPackCmd with the root command during package initialization.
func init() {
	application.HandlePanic(func() {
		application.RootCmd.AddCommand(VlcPackCmd)
	})
}
