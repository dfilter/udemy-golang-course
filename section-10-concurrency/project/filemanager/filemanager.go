package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadData() ([]string, error) {
	// Interesting note: the context for the path seems to be main package,
	// and not the package calling os.Open()
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var strings = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) < 1 {
			continue
		}
		strings = append(strings, text)
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return strings, nil
}

func (fm *FileManager) WriteData(stringMap any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(stringMap)
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
