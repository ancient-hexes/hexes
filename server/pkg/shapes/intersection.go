package shapes

import (
	pb "hexes/proto/go/hexes/v1"
	"hexes/server/pkg/coordinates"
	objectTypes "hexes/server/pkg/object-types"
)

// Intersection returns a slice of mutual coordinates for two objects.
func Intersection(r *objectTypes.Registry, a, b *pb.Map_Object) []*coordinates.Wrapper {
	tiles := coordinates.NewSet(ExtractTileCoordinates(r, a)...)

	intersection := make([]*coordinates.Wrapper, 0, tiles.Size())
	for _, w := range ExtractTileCoordinates(r, b) {
		if tiles.Has(w) {
			intersection = append(intersection, w)
		}
	}
	return intersection
}
