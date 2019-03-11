package jsonapi

import (
	"errors"
	"sync"
)

// SoftCollection ...
type SoftCollection struct {
	typ string
	res []*SoftResource

	m sync.Mutex
}

// AddAttr ...
func (s *SoftCollection) AddAttr(attr Attr) {
	for i := range s.res {
		s.res[i].AddAttr(attr)
	}
}

// AddRel ...
func (s *SoftCollection) AddRel(rel Rel) {
	for i := range s.res {
		s.res[i].AddRel(rel)
	}
}

// Type ...
func (s *SoftCollection) Type() string {
	return s.typ
}

// Len ...
func (s *SoftCollection) Len() int {
	return len(s.res)
}

// Elem ...
func (s *SoftCollection) Elem(i int) Resource {
	if i > 0 && i < len(s.res) {
		return s.res[i]
	}
	return nil
}

// Add ...
func (s *SoftCollection) Add(r Resource) {
	if sr, ok := r.(*SoftResource); ok {
		s.res = append(s.res, sr)
	} else {
		panic("jsonapi: can only add SoftResource to SoftCollection")
	}
}

// Sample ...
func (s *SoftCollection) Sample() Resource {
	if len(s.res) > 0 {
		return s.res[0].New()
	}
	return &SoftResource{}
}

// UnmarshalJSON ...
func (s *SoftCollection) UnmarshalJSON(payload []byte) error {
	return errors.New("jsonapi: SoftCollection.UnmarshalJSON unimplemented")
}
