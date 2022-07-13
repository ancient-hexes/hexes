package objectTypes

import (
	"fmt"

	pb "github.com/ancient-hexes/proto/generated/go/hexes/hexes/v1"
)

type Registry struct {
	m map[string]*pb.Map_Object_Type
}

func NewRegistry() *Registry {
	return &Registry{
		m: make(map[string]*pb.Map_Object_Type),
	}
}

func (r *Registry) Add(tt ...*pb.Map_Object_Type) error {
	for _, t := range tt {
		if _, ok := r.m[t.Id]; ok {
			return fmt.Errorf("duplicate map object registration: '%s'", t.Id)
		}
		r.m[t.Id] = t
	}
	return nil
}

func (r *Registry) MustAdd(tt ...*pb.Map_Object_Type) {
	if err := r.Add(tt...); err != nil {
		panic(err)
	}
}

func (r *Registry) Get(id string) (*pb.Map_Object_Type, bool) {
	t, ok := r.m[id]
	return t, ok
}

func (r *Registry) MustGet(id string) *pb.Map_Object_Type {
	t, ok := r.m[id]
	if !ok {
		panic(fmt.Sprintf("unregistered object ID: %s", id))
	}
	return t
}

func (r *Registry) Validate() error {
	// TODO
	return nil
}
