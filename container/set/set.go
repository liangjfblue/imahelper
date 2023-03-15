package set

// Set 集合
type Set[T comparable] struct {
	data map[T]struct{}
}

// New 实例化Set
func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

// NewWith 实例化Set
func NewWith[T comparable](items ...T) *Set[T] {
	s := &Set[T]{
		data: make(map[T]struct{}, len(items)),
	}
	for _, item := range items {
		s.data[item] = struct{}{}
	}
	return s
}

// Add 添加
func (s *Set[T]) Add(v T) {
	s.data[v] = struct{}{}
}

// Remove 移除
func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

// Contains 是否包含
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.data[v]
	return ok
}

// Equals 是否相等
func (s *Set[T]) Equals(other *Set[T]) bool {
	if other == nil || len(s.data) != len(other.data) {
		return false
	}

	for v := range other.data {
		if !s.Contains(v) {
			return false
		}
	}

	return true
}

// Len 集合元素个数
func (s *Set[T]) Len() int {
	return len(s.data)
}

// Items 集合元素
func (s *Set[T]) Items() []T {
	list := make([]T, 0, len(s.data))
	for v := range s.data {
		list = append(list, v)
	}
	return list
}

// Clear 清空集合
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}
