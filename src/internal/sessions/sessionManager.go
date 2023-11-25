package sessions

type SessionManager struct {
	AllSessions    []GameSession
	AllCompetitors []Person
	UserIDIterator uint
	SessionCount   uint
	Game           game
}

type game interface {
	RunGame(iterations uint) uint8
}

func InitSessionManager(g game) *SessionManager {
	s := SessionManager{
		Game:           g,
		UserIDIterator: 0,
		SessionCount:   0,
	}

	var i uint
	for i = 1; i <= 10; i++ {
		s.UserIDIterator++
		newCompetitor := NewPerson(s.UserIDIterator, s.UserIDIterator)
		s.AllCompetitors = append(s.AllCompetitors, newCompetitor)
	}

	return &s
}

func (s *SessionManager) SpawnSession(p1, p2 *Person) *GameSession {
	s.SessionCount++
	newSession := InitGameSession(s.SessionCount, p1, p2, s.Game)
	s.AllSessions = append(s.AllSessions, newSession)
	return &s.AllSessions[s.SessionCount-1]
}
