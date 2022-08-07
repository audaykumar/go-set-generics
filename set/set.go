package set

// Creating the generic Set data type
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(values ...T) {

	for _, value := range values {
		s[value] = struct{}{}
	}
}

func (s Set[T]) Remove(values ...T) {
	for _, value := range values {
		delete(s, value)
	}
}

func (s Set[T]) Contains(value T) (T, bool) {
	if _, ok := s[value]; ok {
		return value, ok
	}
	return *new(T), false
}

func (s Set[T]) RemoveAll(value T) {

	for value := range s {
		delete(s, value)
	}

}

func (s Set[T]) Len() {

}

func (s Set[T]) Print() {

}
