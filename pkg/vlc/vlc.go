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

var vlc = &cobra.Command{
	Use:   "vlc [file_path]",
	Short: "Pack file using variable-length code",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		application.HandlePanic(func() {
			application.HandleError(func() error {
				// Запуск основной логики упаковки файла
				return pack(args[0])
			})
		})
	},
}

func pack(filePath string) error {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	// Обрабатываем ошибку при закрытии файла
	defer func() {
		application.HandleError(func() error {
			return file.Close()
		})
	}()

	log.Printf("File %s opened successfully", filePath)

	// Читаем данные из файла
	data, err := readFile(file)
	if err != nil {
		return err
	}
	log.Printf("Read %d bytes from file %s", len(data), filePath)

	// Генерация имени для сохранённого файла
	outputFileName := packedFileName(filePath)

	// Запись закодированных данных в файл
	err = writeFile(outputFileName, data)
	if err != nil {
		return err
	}

	log.Printf("Packed file saved as %s", outputFileName)
	return nil
}

// Чтение данных из файла
func readFile(file *os.File) ([]byte, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Запись данных в файл
func writeFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Генерация имени упакованного файла
func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}
