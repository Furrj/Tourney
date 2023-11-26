package sessions

import "fmt"

var names = [10]string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Henry", "Ivy", "Jack"}

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

func InitSessionManager(g game, peopleCount uint) *SessionManager {
	s := SessionManager{
		Game:           g,
		UserIDIterator: 0,
		SessionCount:   0,
	}

	var i uint
	for i = 1; i <= peopleCount; i++ {
		s.UserIDIterator++
		newCompetitor := NewPerson(s.UserIDIterator, names[i-1], s.UserIDIterator)
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

func (s *SessionManager) GetPeopleAsString() string {
	outStr := ""

	for i := 0; i < len(s.AllCompetitors); i++ {
		outStr += fmt.Sprint(&s.AllCompetitors[i])
		outStr += "\n"
	}

	return outStr
}
