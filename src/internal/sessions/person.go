package sessions

type Person struct {
	UserID uint `json:"user_id"`
	Rank   uint `json:"rank"`
	Elo    uint `json:"elo"`
	Wins   uint `json:"wins"`
	Losses uint `json:"losses"`
	Draws  uint `json:"draws"`
}

func NewPerson(id uint, rank uint) Person {
	newPerson := Person{
		UserID: id,
		Rank:   rank,
		Elo:    1500,
		Wins:   0,
		Losses: 0,
		Draws:  0,
	}
	return newPerson
}
