package coordinates

import (
	"fmt"

	pb "github.com/ancient-hexes/proto/generated/go/hexes/hexes/v1"
)

type Wrapper struct {
	X uint32
	Y uint32
	Z uint32
}

func Wrap(c *pb.Coordinate) *Wrapper {
	return &Wrapper{
		X: c.X,
		Y: c.Y,
		Z: c.Z,
	}
}

func New2D(x, y uint32) *Wrapper {
	return &Wrapper{X: x, Y: y}
}

func (w *Wrapper) String() string {
	return fmt.Sprintf("(%d, %d, %d)", w.X, w.Y, w.Z)
}

func (w *Wrapper) MoveTo(c *pb.Coordinate) *Wrapper {
	w.X += c.X
	w.Y += c.Y
	w.Z = c.Z
	return w
}
