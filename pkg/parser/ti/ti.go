package ti

import (
	"fmt"
	"github.com/bazdalaz/cpi/internal/config"
	"os"
)

// TIParser represents a parser for TDC files.
type TIParser struct {
	// Add any fields necessary for the parser here.
}

// New creates a new instance of the TIParser.
func New() *TIParser {
	return &TIParser{}
}

// ParseDir parses the given directory containing FSS files.
func (p *TIParser) ParseDir() error {
	// This is a simplified example. Replace it with actual parsing logic.
	fileDir := config.Config{}.TiPath
	files, err := os.ReadDir(fileDir)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %w", fileDir, err)

	}
	for _, file := range files {
		fmt.Println(file.Name())

	}
	fmt.Printf("Successfully parsed TI directory: %s\n", fileDir)
	return nil
}

//// ParseFile parses the given FSS file.
//func (p *TIParser) ParseFile(filePath string) error {
//
//	// This is a simplified example. Replace it with actual parsing logic.
//	data, err := os.ReadFile(filePath)
//	if err != nil {
//		return fmt.Errorf("error reading file %s: %w", filePath, err)
//	}
//
//	fmt.Printf("Successfully parsed TI file: %s\n", filePath)
//	fmt.Println("File content:", string(data))
//	return nil
//}
