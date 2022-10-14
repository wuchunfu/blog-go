package set

import "encoding/gob"

type Set struct {
	M map[any]struct{}
}

func init() {
	gob.Register(Set{})
}

func (s *Set) Init() {
	s.M = make(map[any]struct{})
}

func (s *Set) Add(val any) {
	s.M[val] = struct{}{}
}

func (s *Set) Del(val any) {
	delete(s.M, val)
}

func (s *Set) Size() int {
	return len(s.M)
}

func (s *Set) Clear() {
	s.M = make(map[any]struct{})
}

func (s *Set) Exist(val any) (ok bool) {
	_, ok = s.M[val]
	return
}

func (s *Set) List() (list []any) {
	for v := range s.M {
		list = append(list, v)
	}
	return
}
