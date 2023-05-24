package grid

type Direction struct{ dx, dy float64 }

func NewDirection(dx, dy float64) Direction {
	return Direction{dx: dx, dy: dy}
}

func (this Direction) Dx() float64 { return this.dx }
func (this Direction) Dy() float64 { return this.dy }

func (this Direction) TurnRight() Direction { return Clockwise[this] }
func (this Direction) TurnLeft() Direction  { return CounterClockwise[this] }

var (
	Static Direction

	Right = NewDirection(1, 0)
	Left  = NewDirection(-1, 0)
	Up    = NewDirection(0, 1)
	Down  = NewDirection(0, -1)

	TopRight    = NewDirection(1, 1)
	TopLeft     = NewDirection(-1, 1)
	BottomRight = NewDirection(1, -1)
	BottomLeft  = NewDirection(-1, -1)

	Neighbors4 = []Direction{Right, Left, Up, Down}
	Neighbors8 = append(Neighbors4, []Direction{TopRight, TopLeft, BottomRight, BottomLeft}...)

	Clockwise = map[Direction]Direction{
		Down:  Left,
		Left:  Up,
		Up:    Right,
		Right: Down,
	}
	CounterClockwise = map[Direction]Direction{
		Down:  Right,
		Right: Up,
		Up:    Left,
		Left:  Down,
	}
)
