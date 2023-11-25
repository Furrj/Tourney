package sessions

type Person struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Rank   uint   `json:"rank"`
	Elo    uint   `json:"elo"`
}

func NewPerson() Person {
	var newPerson Person
	newPerson.Elo = 1500
	return newPerson
}
