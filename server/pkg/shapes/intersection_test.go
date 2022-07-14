package shapes_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	pb "hexes/proto/go/hexes/v1"
	objectTypes "hexes/server/pkg/object-types"
	"hexes/server/pkg/shapes"
)

func TestIntersection(t *testing.T) {
	r := objectTypes.NewRegistry()
	r.MustAdd(objectTypes.CreatureBanks...)
	r.MustAdd(objectTypes.Treasuries...)

	a := &pb.Map_Object{
		Type:       "Base/CreatureBanks/DragonFlyHive",
		Coordinate: &pb.Coordinate{X: 0, Y: 0, Z: 0},
	}
	b := &pb.Map_Object{
		Type:       "Base/Treasuries/DragonUtopia",
		Coordinate: &pb.Coordinate{X: 0, Y: 0, Z: 0},
	}
	intersection := shapes.Intersection(r, a, b)
	require.Equal(t, 0, len(intersection))

	a.Coordinate.X += 2
	intersection = shapes.Intersection(r, a, b)
	require.Equal(t, 1, len(intersection))
	require.Equal(t, uint32(2), intersection[0].X)
	require.Equal(t, uint32(0), intersection[0].Y)
}
