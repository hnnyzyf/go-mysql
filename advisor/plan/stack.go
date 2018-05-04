package plan

type stack struct {
	element []interface{}
	size    int
}

func newStack() *stack {
	return &stack{
		element: []interface{}{},
		size:    0,
	}
}
func (s *stack) top() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.element[s.size-1]
}

func (s *stack) push(e interface{}) {
	if len(s.element) > s.size {
		s.element[s.size] = e
	} else {
		s.element = append(s.element, e)
	}
	s.size++
}

func (s *stack) pop() (interface{}, bool) {
	if s.size <= 0 {
		return nil, false
	}
	element := s.top()
	s.size--
	return element, true
}

func (s *stack) popAll() {
	s.size = 0
}

func (s *stack) all() []interface{} {
	if s.size == 0 {
		return []interface{}{}
	} else {
		return s.element[0:s.size]
	}
}
