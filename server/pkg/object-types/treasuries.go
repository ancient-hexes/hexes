package objectTypes

import (
	pb "github.com/ancient-hexes/proto/generated/go/hexes/hexes/v1"
)

var Treasuries = []*pb.Map_Object_Type{
	{
		Id:     "Base/Treasuries/DragonUtopia",
		Height: 3,
		Width:  6,
		Shape: []*pb.Coordinate{
			{X: 0, Y: 1},
			{X: 1, Y: 1},
			{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2},
			{X: 3, Y: 0}, {X: 3, Y: 1}, {X: 3, Y: 2},
			{X: 4, Y: 0}, {X: 4, Y: 1}, {X: 4, Y: 2},
			{X: 5, Y: 2},
		},
		Entrance: &pb.Coordinate{X: 4, Y: 0},
	},
}
