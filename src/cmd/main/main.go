package main

import (
	"fmt"
	"github.com/Furrj/Tourney/src/internal/games"
	"github.com/Furrj/Tourney/src/internal/sessions"
)

func main() {
	manager := sessions.InitSessionManager(games.LoHigh{}, 10)

	sesh := manager.SpawnSession(&manager.AllCompetitors[0], &manager.AllCompetitors[1])

	sesh.ExecuteSession(10)

	fmt.Println(manager.GetPeopleAsString())
}
