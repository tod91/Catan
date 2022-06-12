package player

type Player struct {
	Cities      uint32
	Settlements uint32
	Roads       uint32
}

type Turn interface {
	DrawDevCard()
	RollDice()
	Build()
	Trade()
}

func (p *Player) DrawDevCard() {

}

func (p *Player) RollDice() {

}

func (p *Player) Build() {

}

func (p *Player) Trade() {

}
