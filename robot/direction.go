package robot

type Direction interface {
	isDirection()
}

type North struct{}

func (North) isDirection() {}

type East struct{}

func (East) isDirection() {}

type South struct{}

func (South) isDirection() {}

type West struct{}

func (West) isDirection() {}
