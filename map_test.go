package tsyncmap

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testKey struct {
	id int
}

type testValue struct {
	name string
}

func TestMap_Load_LoadsOnStored(t *testing.T) {
	t.Parallel()
	// Arrange
	key := testKey{id: 5}
	value := testValue{name: "abc"}

	item := &Map[testKey, testValue]{}
	item.Store(key, value)

	// Act
	result, ok := item.Load(key)

	// Assert
	assert.True(t, ok)
	assert.Equal(t, value, result)
}

func TestMap_Load_DoesNotLoadOnNotStored(t *testing.T) {
	t.Parallel()
	// Arrange
	key := testKey{id: 5}

	item := &Map[testKey, testValue]{}

	// Act
	result, ok := item.Load(key)

	// Assert
	assert.False(t, ok)
	assert.Empty(t, result)
}

func TestMap_Store_Stores(t *testing.T) {
	t.Parallel()
	// Arrange
	key := testKey{id: 5}
	value := testValue{name: "abc"}

	item := &Map[testKey, testValue]{}

	// Act
	item.Store(key, value)

	// Assert
	result, ok := item.Load(key)

	assert.True(t, ok)
	assert.Equal(t, value, result)
}

func TestMap_Delete_Deletes(t *testing.T) {
	t.Parallel()
	// Arrange
	key := testKey{id: 5}
	value := testValue{name: "abc"}

	item := &Map[testKey, testValue]{}
	item.Store(key, value)

	// Act
	item.Delete(key)

	// Assert
	result, ok := item.Load(key)

	assert.False(t, ok)
	assert.Empty(t, result)
}

func TestMap_LoadOrStore_StoresOnNotExist(t *testing.T) {
	t.Parallel()
	// Arrange
	key := testKey{id: 5}
	value := testValue{name: "abc"}

	item := &Map[testKey, testValue]{}

	// Act
	result, ok := item.LoadOrStore(key, value)

	// Assert
	assert.False(t, ok)
	assert.Equal(t, value, result)
}

func TestMap_LoadOrStore_LoadsIfExists(t *testing.T) {
	t.Parallel()
	// Arrange
	key := testKey{id: 5}
	value := testValue{name: "abc"}

	item := &Map[testKey, testValue]{}
	item.Store(key, value)

	// Act
	result, ok := item.LoadOrStore(key, value)

	// Assert
	assert.True(t, ok)
	assert.Equal(t, value, result)
}

func TestMap_LoadAndDelete_LoadAndDeletesIfExists(t *testing.T) {
	t.Parallel()
	// Arrange
	key := testKey{id: 5}
	value := testValue{name: "abc"}

	item := &Map[testKey, testValue]{}
	item.Store(key, value)

	// Act
	result, ok := item.LoadAndDelete(key)

	// Assert
	assert.True(t, ok)
	assert.Equal(t, value, result)
}

func TestMap_LoadAndDelete_LoadAndDeletesDoesNothingOnNotExists(t *testing.T) {
	t.Parallel()
	// Arrange
	key := testKey{id: 5}

	item := &Map[testKey, testValue]{}

	// Act
	result, ok := item.LoadAndDelete(key)

	// Assert
	assert.False(t, ok)
	assert.Empty(t, result)
}

func TestMap_Range_Ranges(t *testing.T) {
	t.Parallel()
	// Arrange
	key1 := testKey{id: 5}
	key2 := testKey{id: 6}
	key3 := testKey{id: 7}
	value1 := testValue{name: "a"}
	value2 := testValue{name: "b"}
	value3 := testValue{name: "c"}

	item := &Map[testKey, testValue]{}
	item.Store(key1, value1)
	item.Store(key2, value2)
	item.Store(key3, value3)

	var calledWithKey []testKey
	var calledWithValue []testValue

	// Act
	item.Range(func(k testKey, v testValue) bool {
		calledWithKey = append(calledWithKey, k)
		calledWithValue = append(calledWithValue, v)

		return true
	})

	// Assert
	assert.ElementsMatch(t, []testKey{key1, key2, key3}, calledWithKey)
	assert.ElementsMatch(t, []testValue{value1, value2, value3}, calledWithValue)
}

// Benchmarks

func BenchmarkMap_Load(b *testing.B) {
	object := new(Map[testKey, testValue])

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	object.Store(key, value)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.Load(key)
	}
}

func BenchmarkSyncMap_Load(b *testing.B) {
	object := new(sync.Map)

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	object.Store(key, value)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.Load(key)
	}
}

func BenchmarkMap_Store(b *testing.B) {
	object := new(Map[testKey, testValue])

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.Store(key, value)
	}
}

func BenchmarkSyncMap_Store(b *testing.B) {
	object := new(sync.Map)

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.Store(key, value)
	}
}

func BenchmarkMap_LoadOrStore(b *testing.B) {
	object := new(Map[testKey, testValue])

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.LoadOrStore(key, value)
	}
}

func BenchmarkSyncMap_LoadOrStore(b *testing.B) {
	object := new(sync.Map)

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.LoadOrStore(key, value)
	}
}

func BenchmarkMap_Delete(b *testing.B) {
	object := new(Map[testKey, testValue])

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	object.Store(key, value)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.Delete(key)
	}
}

func BenchmarkSyncMap_Delete(b *testing.B) {
	object := new(sync.Map)

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	object.Store(key, value)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.Delete(key)
	}
}

func BenchmarkMap_LoadAndDelete(b *testing.B) {
	object := new(Map[testKey, testValue])

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	object.Store(key, value)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.LoadAndDelete(key)
	}
}

func BenchmarkSyncMap_LoadAndDelete(b *testing.B) {
	object := new(sync.Map)

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	object.Store(key, value)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.LoadAndDelete(key)
	}
}

func BenchmarkMap_Range(b *testing.B) {
	object := new(Map[testKey, testValue])

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	object.Store(key, value)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.Range(func(key testKey, value testValue) bool {
			return true
		})
	}
}

func BenchmarkSyncMap_Range(b *testing.B) {
	object := new(sync.Map)

	key := testKey{id: 5}
	value := testValue{name: "abc"}

	object.Store(key, value)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		object.Range(func(key any, value any) bool {
			return true
		})
	}
}
