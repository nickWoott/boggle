package player

import "errors"

type Player struct {
	Score    int
	WordList []string
}

func NewPlayer() *Player {

	return &Player{
		Score:    0,
		WordList: []string{},
	}

}

func (p *Player) AddWord(word string) error {

	if len(word) < 3 {
		return errors.New("word too short")
	}

	switch length := len(word); {
	case length >= 8:
		p.Score += 11
	case length == 7:
		p.Score += 4
	case length == 6:
		p.Score += 3
	case length == 5:
		p.Score += 2
	case length == 3 || length == 4:
		p.Score += 1
	}

	return nil
}
