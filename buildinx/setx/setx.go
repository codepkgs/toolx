package setx

import (
	"fmt"
)

type Set struct {
	m map[string]string
}

func NewSet(keys ...string) *Set {
	m := make(map[string]string)
	for _, k := range keys {
		m[k] = ""
	}

	return &Set{
		m: m,
	}
}

func (s *Set) String() string {
	return fmt.Sprintf("%q", s.Keys())
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) Keys() []string {
	keys := make([]string, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}
func (s *Set) Has(key string) bool {
	_, ok := s.m[key]
	return ok
}

func (s *Set) Add(keys ...string) {
	for _, k := range keys {
		s.m[k] = ""
	}
}

func (s *Set) Del(key string) {
	if s.Len() == 0 {
		return
	}
	delete(s.m, key)
}

func (s *Set) Intersection(o *Set) []string {
	keys := make([]string, 0)
	if s.Len() == 0 || o.Len() == 0 {
		return keys
	}

	for k := range s.m {
		if o.Has(k) {
			keys = append(keys, k)
		}
	}
	return keys
}

func (s *Set) Union(o *Set) []string {
	m := make(map[string]string)
	if s.Len() == 0 {
		return s.Keys()
	}

	if o.Len() == 0 {
		return o.Keys()
	}

	for k := range s.m {
		m[k] = ""
	}

	for k := range o.m {
		m[k] = ""
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
