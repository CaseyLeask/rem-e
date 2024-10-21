package robot

type Command interface {
	isCommand()
}

type Place struct {
	X         int
	Y         int
	Direction Direction
}

func (Place) isCommand() {}

type Move struct{}

func (Move) isCommand() {}

type Left struct{}

func (Left) isCommand() {}

type Right struct{}

func (Right) isCommand() {}

type Report struct{}

func (Report) isCommand() {}
