package main

import (
	"log"
	"os"

	"github.com/kqns91/go-generator/internal/generator"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("no target directory")
	}
	targetDir := os.Args[1]
	if err := generator.Run(targetDir); err != nil {
		log.Fatalf("failed to generate: %v", err)
	}
}
