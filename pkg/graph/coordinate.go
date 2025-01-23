package graph

import (
	"fmt"
	"math"
)

type Coordinate struct {
	Row   int
	Col   int
	Value string
}

func NewCoordinate() Coordinate {
	return Coordinate{Row: -1, Col: -1, Value: ""}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf(
		"Coordinate(Row: %d, Col: %d, Val: <%s>)",
		c.Row, c.Col, c.Value,
	)
}

func (c *Coordinate) HashKey() string {
	return fmt.Sprintf("(%d,%d)", c.Row, c.Col)
}

func (c *Coordinate) Subtract(other Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row - other.Row,
		Col: c.Col - other.Col,
	}
}

func (c *Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row + other.Row,
		Col: c.Col + other.Col,
	}
}

func (c *Coordinate) Distance(other Coordinate) float32 {
	return float32(math.Sqrt(
		math.Pow((float64(c.Row-other.Row)), 2) + math.Pow(float64(c.Col-other.Col), 2),
	))
}

func (c *Coordinate) Multiply(n int) Coordinate {
	return Coordinate{
		Row: c.Row * n,
		Col: c.Col * n,
	}
}

func (c *Coordinate) Equals(other Coordinate) bool {
	return c.Row == other.Row && c.Col == other.Col
}

func (c *Coordinate) NextCoordinates(
	maxRow int, maxCol int, grid *[][]string, obstruction string,
) []Coordinate {
	retval := []Coordinate{}
	// up
	if c.Row-1 >= 0 && (*grid)[c.Row-1][c.Col] != obstruction {
		retval = append(retval, Coordinate{
			Row: c.Row - 1, Col: c.Col, Value: (*grid)[c.Row-1][c.Col],
		})
	}
	// down
	if c.Row+1 < maxRow && (*grid)[c.Row+1][c.Col] != obstruction {
		retval = append(retval, Coordinate{
			Row: c.Row + 1, Col: c.Col, Value: (*grid)[c.Row+1][c.Col],
		})
	}
	// left
	if c.Col-1 >= 0 && (*grid)[c.Row][c.Col-1] != obstruction {
		retval = append(retval, Coordinate{
			Row: c.Row, Col: c.Col - 1, Value: (*grid)[c.Row][c.Col-1],
		})
	}
	//right
	if c.Col+1 < maxCol && (*grid)[c.Row][c.Col+1] != obstruction {
		retval = append(retval, Coordinate{
			Row: c.Row, Col: c.Col + 1, Value: (*grid)[c.Row][c.Col+1],
		})
	}

	return retval
}
