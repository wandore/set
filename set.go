package set

import "sync"

var empty = struct{}{}

type Set struct {
	data map[interface{}]struct{}
	mu   *sync.RWMutex
}

// 添加元素
func (s *Set) Add(elem interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[elem] = empty

	return nil
}

// 判断是否包含元素
func (s *Set) Contain(elem interface{}) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[elem]

	return ok
}

// 获取 set 大小
func (s *Set) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}

// 清空 set
func (s *Set) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[interface{}]struct{})
}

// 构造 set 实例
func New() *Set {
	return &Set{
		data: make(map[interface{}]struct{}),
		mu:   &sync.RWMutex{},
	}
}
