package main

import (
	"log"

	"github.com/krok-o/krokctl/cmd"
	_ "github.com/krok-o/krokctl/cmd/repositories"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
