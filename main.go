package main

import (
	"log"

	"github.com/krok-o/krokctl/cmd"
	_ "github.com/krok-o/krokctl/cmd/commands"
	_ "github.com/krok-o/krokctl/cmd/events"
	_ "github.com/krok-o/krokctl/cmd/platforms"
	_ "github.com/krok-o/krokctl/cmd/repositories"
	_ "github.com/krok-o/krokctl/cmd/settings"
	_ "github.com/krok-o/krokctl/cmd/vault"
	_ "github.com/krok-o/krokctl/cmd/vcs"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
