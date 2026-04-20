package main

import (
	site "Atrium/internal/core"
	"log"
)

func main() {
	generator := site.NewGenerator()

	if err := generator.Generate(); err != nil {
		log.Fatal(err)
	}
}
