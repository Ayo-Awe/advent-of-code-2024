package main

type Direction string

const (
	Up    Direction = "^"
	Down  Direction = "v"
	Right Direction = ">"
	Left  Direction = "<"
)

func (d Direction) Rotate90Clockwise() Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		return Up
	}
}

type Guard struct {
	x         int
	y         int
	direction Direction
}

func (g *Guard) LookAhead() (int, int) {
	switch g.direction {
	case Up:
		return g.x, g.y - 1
	case Down:
		return g.x, g.y + 1
	case Left:
		return g.x - 1, g.y
	case Right:
		return g.x + 1, g.y
	default:
		return g.x, g.y
	}
}

func (g *Guard) Forward() {
	g.x, g.y = g.LookAhead()
}

func (g *Guard) Rotate90Clockwise() {
	g.direction = g.direction.Rotate90Clockwise()
}

func isGuard(v string) bool {
	switch Direction(v) {
	case Up, Down, Left, Right:
		return true
	default:
		return false
	}
}
