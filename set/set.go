package set

var _defValue = 1

type Set struct {
	hash map[interface{}]int
}

func New(initial ...interface{}) *Set {
	s := &Set{make(map[interface{}]int)}
	for _, v := range initial {
		s.hash[v] = _defValue
	}
	return s
}

func (s *Set) Has(element interface{}) bool {
	_, exists := s.hash[element]
	return exists
}

func (s *Set) Insert(element interface{}) {
	s.hash[element] = _defValue
}

func (s *Set) Remove(element interface{}) {
	delete(s.hash, element)
	return
}

func (s *Set) Len() int {
	return len(s.hash)
}

func (s *Set) Do(f func(interface{})) {
	for k, _ := range s.hash {
		f(k)
	}
}

// option

func (s *Set) Union(set *Set) *Set {
	result := make(map[interface{}]int)
	for k, _ := range s.hash {
		result[k] = _defValue
	}
	for k, _ := range set.hash {
		result[k] = _defValue
	}
	return &Set{result}
}

func (s *Set) Intersection(set *Set) *Set {
	result := make(map[interface{}]int)
	for k, _ := range s.hash {
		if _, exists := set.hash[k]; exists {
			result[k] = _defValue
		}
	}
	return &Set{result}
}

func (s *Set) Difference(set *Set) *Set {
	result := make(map[interface{}]int)
	for k, _ := range s.hash {
		if _, exists := set.hash[k]; !exists {
			result[k] = _defValue
		}
	}
	return &Set{result}
}

func (s *Set) SymmetricDifference(set *Set) *Set {
	a := s.Difference(set)
	b := set.Difference(s)
	return a.Union(b)
}
