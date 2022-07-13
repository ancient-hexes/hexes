package coordinates

type Set struct {
	m map[Wrapper]struct{}
}

func NewSet(ww ...*Wrapper) *Set {
	s := &Set{
		m: make(map[Wrapper]struct{}),
	}
	s.Add(ww...)
	return s
}

func (s *Set) Add(ww ...*Wrapper) {
	for _, w := range ww {
		s.m[*w] = struct{}{}
	}
}

func (s *Set) Has(w *Wrapper) bool {
	_, ok := s.m[*w]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}
