package vlc

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/flexer2006/simpleArchiver-golang/internal/application"
	"github.com/spf13/cobra"
)

const packedExtension = "vlc"

var Vlc = &cobra.Command{
	Use:   "vlc [file_path]",
	Short: "Pack file using variable-length code",
	Run: func(cmd *cobra.Command, args []string) {
		application.HandlePanic(func() {
			application.HandleError(func() error {
				if len(args) == 0 || args[0] == "" {
					return application.ErrEmptyPath
				}
				return pack(args[0])
			})
		})
	},
}

func pack(filePath string) error {
	file, err := openFile(filePath)
	if err != nil {
		return err
	}

	defer application.HandleError(func() error {
		return file.Close()
	})

	log.Printf("File %s opened successfully", filePath)

	data, err := readFile(file)
	if err != nil {
		return err
	}
	log.Printf("Read %d bytes from file %s", len(data), filePath)

	// Преобразуем данные через Encode
	encodedData := Encode(string(data))
	log.Printf("Encoded data: %s", encodedData)

	outputFileName := packedFileName(filePath)

	err = writeFile(outputFileName, []byte(encodedData))
	if err != nil {
		return err
	}

	log.Printf("Packed file saved as %s", outputFileName)
	return nil
}

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func readFile(file *os.File) ([]byte, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func writeFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {

	application.HandlePanic(func() {
		application.HandleError(func() error {
			application.RootCmd.AddCommand(Vlc)
			return nil
		})
	})
}
