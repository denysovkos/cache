package cache

import (
	"fmt"
	"sync"
	"testing"
)

// Test for cache initialization
func TestNewCache(t *testing.T) {
	_, err := NewCache(0)
	if err == nil {
		t.Fatal("expected error for zero capacity")
	}

	cache, err := NewCache(2)
	if err != nil {
		t.Fatal(err)
	}

	if cache.capacity != 2 {
		t.Fatalf("expected capacity to be 2 but got %d", cache.capacity)
	}
}

// Positive test for cache operations
func TestCachePositive(t *testing.T) {
	cache, err := NewCache(2)
	if err != nil {
		t.Fatal(err)
	}

	// Add items to the cache
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Retrieve items and check values
	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Fatalf("expected key1 to be 'value1' but got %v", val)
	}

	// Override an existing key
	cache.Set("key1", "new_value1")
	if val, ok := cache.Get("key1"); !ok || val != "new_value1" {
		t.Fatalf("expected key1 to be 'new_value1' but got %v", val)
	}

	// Add another item to cause eviction
	cache.Set("key3", "value3")
	if _, ok := cache.Get("key2"); ok {
		t.Fatal("expected key2 to be evicted")
	}

	// Add one more item to ensure capacity handling
	cache.Set("key4", "value4")
	if val, ok := cache.Get("key3"); !ok || val != "value3" {
		t.Fatalf("expected key3 to be 'value3' but got %v", val)
	}
}

// Negative test for cache operations
func TestCacheNegative(t *testing.T) {
	cache, err := NewCache(2)
	if err != nil {
		t.Fatal(err)
	}

	// Check non-existent key
	if _, ok := cache.Get("non_existent"); ok {
		t.Fatal("expected non_existent key to not be found")
	}

	// Add and evict items
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	// Check eviction of key1
	if _, ok := cache.Get("key1"); ok {
		t.Fatal("expected key1 to be evicted")
	}
}

// Test to check if key already exists and overriding it
func TestCacheOverride(t *testing.T) {
	cache, err := NewCache(2)
	if err != nil {
		t.Fatal(err)
	}

	cache.Set("key1", "value1")
	cache.Set("key1", "value2")

	if val, ok := cache.Get("key1"); !ok || val != "value2" {
		t.Fatalf("expected key1 to be 'value2' but got %v", val)
	}
}

// Test to remove a key manually
func TestCacheRemove(t *testing.T) {
	cache, err := NewCache(2)
	if err != nil {
		t.Fatal(err)
	}

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Remove key1 manually
	if removed := cache.Remove("key1"); !removed {
		t.Fatal("expected key1 to be removed")
	}

	// Check if key1 is actually removed
	if _, ok := cache.Get("key1"); ok {
		t.Fatal("expected key1 to be not found after removal")
	}

	// Add another item to ensure cache behaves correctly after removal
	cache.Set("key3", "value3")
	if val, ok := cache.Get("key3"); !ok || val != "value3" {
		t.Fatalf("expected key3 to be 'value3' but got %v", val)
	}
}

// Test to check thread safety
func TestCacheThreadSafety(t *testing.T) {
	cache, err := NewCache(1000)
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup

	// Writer goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				cache.Set(fmt.Sprintf("key-%d-%d", i, j), fmt.Sprintf("value-%d-%d", i, j))
			}
		}(i)
	}

	// Reader goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				cache.Get(fmt.Sprintf("key-%d-%d", i, j))
			}
		}(i)
	}

	wg.Wait()

	// Check some values to ensure the cache is still functioning
	for i := 0; i < 10; i++ {
		if val, ok := cache.Get(fmt.Sprintf("key-%d-999", i)); !ok || val != fmt.Sprintf("value-%d-999", i) {
			t.Logf("expected key-%d-999 to be 'value-%d-999' but got %v", i, i, val)
		}
	}
}
