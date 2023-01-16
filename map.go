package tsyncmap

import "sync"

// Map mimics sync.Map but with generics
type Map[K any, V any] struct {
	// sync.Map is exported for flexibility, so you can still
	// use it if required
	sync.Map
}

// Load mimics sync.Map{}'s Load
func (m *Map[K, V]) Load(key K) (V, bool) {
	result, ok := m.Map.Load(key)
	if ok {
		return result.(V), true
	}

	return *new(V), false
}

// Store mimics sync.Map{}'s Store
func (m *Map[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}

// LoadOrStore mimics sync.Map{}'s LoadOrStore
func (m *Map[K, V]) LoadOrStore(key K, value V) (V, bool) {
	result, ok := m.Map.LoadOrStore(key, value)
	if ok {
		return result.(V), true
	}

	return value, false
}

// LoadAndDelete mimics sync.Map{}'s LoadAndDelete
func (m *Map[K, V]) LoadAndDelete(key K) (V, bool) {
	item, ok := m.Map.LoadAndDelete(key)

	if ok {
		return item.(V), true
	}

	return *new(V), false
}

// Delete mimics sync.Map{}'s Delete
func (m *Map[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

// Range mimics sync.Map{}'s Range
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
