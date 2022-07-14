package objectTypes

import (
	pb "hexes/proto/go/hexes/v1"
)

var CreatureBanks = []*pb.Map_Object_Type{
	{
		Id:       "Base/CreatureBanks/DragonFlyHive",
		Height:   1,
		Width:    1,
		Entrance: &pb.Coordinate{X: 0, Y: 0},
	},
	{
		Id:       "Base/CreatureBanks/GriffinConservatory",
		Height:   1,
		Width:    2,
		Entrance: &pb.Coordinate{X: 1, Y: 0},
	},
}
