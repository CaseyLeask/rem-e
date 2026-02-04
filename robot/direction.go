package robot

type Direction interface {
	dx() int
	dy() int

	clockwise() Direction
	counterClockwise() Direction

	String() string
}

type North struct{}

func (North) dx() int {
	return 0
}

func (North) dy() int {
	return 1
}

func (North) clockwise() Direction {
	return East{}
}

func (North) counterClockwise() Direction {
	return West{}
}

func (North) String() string {
	return "NORTH"
}

type East struct{}

func (East) dx() int {
	return 1
}

func (East) dy() int {
	return 0
}

func (East) clockwise() Direction {
	return South{}
}

func (East) counterClockwise() Direction {
	return North{}
}

func (East) String() string {
	return "EAST"
}

type South struct{}

func (South) dx() int {
	return 0
}

func (South) dy() int {
	return -1
}

func (South) clockwise() Direction {
	return West{}
}

func (South) counterClockwise() Direction {
	return East{}
}

func (South) String() string {
	return "SOUTH"
}

type West struct{}

func (West) dx() int {
	return -1
}

func (West) dy() int {
	return 0
}

func (West) clockwise() Direction {
	return North{}
}

func (West) counterClockwise() Direction {
	return South{}
}

func (West) String() string {
	return "WEST"
}
