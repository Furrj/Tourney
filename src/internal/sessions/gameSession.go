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
	draw  = 3
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
