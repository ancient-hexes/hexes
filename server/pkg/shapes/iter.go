package shapes

import (
	pb "hexes/proto/go/hexes/v1"
	"hexes/server/pkg/coordinates"
	objectTypes "hexes/server/pkg/object-types"
)

func ExtractTileCoordinates(r *objectTypes.Registry, a *pb.Map_Object) []*coordinates.Wrapper {
	t := r.MustGet(a.Type)

	result := make([]*coordinates.Wrapper, 0, t.Height*t.Width)
	if len(t.Shape) > 0 {
		for _, c := range t.Shape {
			result = append(
				result,
				coordinates.Wrap(c).MoveTo(a.Coordinate),
			)
		}
	} else {
		for x := uint32(0); x < t.Width; x++ {
			for y := uint32(0); y < t.Height; y++ {
				result = append(result, coordinates.New2D(x, y).MoveTo(a.Coordinate))
			}
		}
	}
	return result
}
