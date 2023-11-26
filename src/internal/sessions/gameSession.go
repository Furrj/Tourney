package sessions

type GameSession struct {
	SessionID uint
	P1        *Person
	P2        *Person
	Game      game
	Winner    uint
}

const (
	p1Win = 1
	p2Win = 2
)

func InitGameSession(sessionId uint, p1, p2 *Person, g game) GameSession {
	return GameSession{
		SessionID: sessionId,
		P1:        p1,
		P2:        p2,
		Game:      g,
		Winner:    0,
	}
}

func (g *GameSession) ExecuteSession(iterations uint) {
	result := g.Game.RunGame(iterations)

	switch result {
	case p1Win:
		g.P1.Wins++
		g.P2.Losses++
	case p2Win:
		g.P2.Wins++
		g.P2.Losses++
	default:
		g.P1.Draws++
		g.P2.Draws++
	}
}
