package robot

import "fmt"

type Robot struct {
	X      int
	Y      int
	facing Direction
}

func NewRobot() Robot {
	return Robot{}
}

const lowerBound = 0
const upperbound = 4

func (r Robot) IsOnTable() bool {
	return r.facing != nil &&
		r.X >= lowerBound &&
		r.Y <= upperbound &&
		r.Y >= lowerBound &&
		r.X <= upperbound
}

func (r Robot) Run(command Command) (Robot, *string) {
	switch command.(type) {
	case Place:
		place := command.(Place)
		newRobot := Robot{
			place.X,
			place.Y,
			place.Direction,
		}
		if newRobot.IsOnTable() {
			return newRobot, nil
		} else {
			return r, nil
		}

	case Move:
		newRobot := Robot{
			r.X + r.facing.dx(),
			r.Y + r.facing.dy(),
			r.facing,
		}
		if newRobot.IsOnTable() {
			return newRobot, nil
		} else {
			return r, nil
		}

	case Left:
		return Robot{
			r.X,
			r.Y,
			r.facing.counterClockwise(),
		}, nil

	case Right:
		return Robot{
			r.X,
			r.Y,
			r.facing.clockwise(),
		}, nil

	case Report:
		if r.IsOnTable() {
			s := fmt.Sprintf("Output: %d,%d,%s", r.X, r.Y, r.facing)
			return r, &s
		} else {
			return r, nil
		}

	default:
		return r, nil
	}
}
