package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
)

func Run(targetDir string) error {
	structs, err := parseStructs(targetDir)
	if err != nil {
		return fmt.Errorf("failed to parse structs: %w", err)
	}

	for _, structInfo := range structs {
		if err := generateCode(structInfo, targetDir); err != nil {
			return fmt.Errorf("failed to generate code: %w", err)
		}
	}

	return nil
}

func generateCode(structInfo *StructInfo, outputDir string) error {
	var buf bytes.Buffer
	err := codeTemplate.Execute(&buf, structInfo)
	if err != nil {
		return fmt.Errorf("failed to execute template %w", err)
	}

	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format code: %w", err)
	}

	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_gen.go", uncapitalize(structInfo.Name)))

	err = os.WriteFile(outputFile, formattedCode, 0644)
	if err != nil {
		return fmt.Errorf("failed to write code: %w", err)
	}

	fmt.Printf("generated code for struct %s at %s\n", structInfo.Name, outputFile)
	return nil
}
