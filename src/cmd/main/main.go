package main

import (
	"fmt"
	"github.com/Furrj/Tourney/src/internal/games"
	"github.com/Furrj/Tourney/src/internal/sessions"
)

func main() {
	sessionManager := sessions.InitSessionManager(games.LoHigh{})

	fmt.Printf("%+v\n", sessionManager)
}
