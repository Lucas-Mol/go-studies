package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("failed to read line in file")
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	fm.OutputFilePath = strings.TrimSpace(fm.OutputFilePath)
	if !strings.HasSuffix(fm.OutputFilePath, ".json") {
		fm.OutputFilePath += ".json"
	}

	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to encode data to JSON")
	}

	return nil
}
