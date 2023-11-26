package sessions

import "fmt"

type Person struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Rank   uint   `json:"rank"`
	Elo    uint   `json:"elo"`
	Wins   uint   `json:"wins"`
	Losses uint   `json:"losses"`
	Draws  uint   `json:"draws"`
}

func NewPerson(id uint, name string, rank uint) Person {
	newPerson := Person{
		UserID: id,
		Name:   name,
		Rank:   rank,
		Elo:    1500,
		Wins:   0,
		Losses: 0,
		Draws:  0,
	}
	return newPerson
}

func (p *Person) String() string {
	outStr := fmt.Sprintf("Id: %d\nName: %s\n", p.UserID, p.Name)
	outStr += fmt.Sprintf("Rank: %d\n", p.Rank)
	outStr += fmt.Sprintf("Elo: %d\n", p.Elo)
	outStr += fmt.Sprintf("Record: %d/%d/%d\n", p.Wins, p.Losses, p.Draws)

	return outStr
}
