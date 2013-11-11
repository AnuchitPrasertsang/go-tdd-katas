package rover

import (
	"fmt"
)

const (
	North = iota
	East  = iota
	South = iota
	West  = iota
)

type Coordinates struct {
	X int
	Y int
}

type Rover struct {
	Coords Coordinates
	Facing int
	Grid   *Grid
}

/*
 * Just called New because external users will
 * use rover.New()
 */
func New(coords Coordinates, facing int) *Rover {
	rover := new(Rover)
	rover.Coords = coords
	rover.Facing = facing
	rover.Grid = NewGrid(100, 100)
	rover.Grid.Snap(rover)
	return rover
}

func (rover *Rover) TurnRight() {
	switch direction := rover.Facing; direction {
	case North:
		rover.Facing = East
	case East:
		rover.Facing = South
	case South:
		rover.Facing = West
	case West:
		rover.Facing = North
	default:
		panic(fmt.Sprintf("WTF is this direction: %d", direction))
	}
}

func (rover *Rover) TurnLeft() {
	// turn left == turn right 3 times
	// screw efficiency, let's be DRY
	for i := 0; i < 3; i++ {
		rover.TurnRight()
	}
}

func (rover *Rover) TurnBack() {
	// using TurnRight because TurnLeft is slower
	rover.TurnRight()
	rover.TurnRight()
}

func (rover *Rover) Advance() {
	rover.Grid.Insert(rover.Coords, NOTHING)
	switch direction := rover.Facing; direction {
	case North:
		rover.Coords.Y += 1
	case East:
		rover.Coords.X += 1
	case West:
		rover.Coords.X -= 1
	case South:
		rover.Coords.Y -= 1
	}
	rover.Grid.Snap(rover)
}

func (rover *Rover) Retreat() {
	rover.TurnBack()
	rover.Advance()
	rover.TurnBack()
}