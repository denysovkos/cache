package main

import (
	"fmt"
	"semrush/cache/cache"
)

func main() {
	// Create a new cache with a capacity of 2
	myCache, err := cache.NewCache(2)
	if err != nil {
		fmt.Println("Error creating cache:", err)
		return
	}

	// Add some items to the cache
	myCache.Set("key1", "value1")
	myCache.Set("key2", "value2")

	// Retrieve and print items from the cache
	if val, ok := myCache.Get("key1"); ok {
		fmt.Println("key1:", val)
	} else {
		fmt.Println("key1 not found")
	}

	// Add another item, causing an eviction
	myCache.Set("key3", "value3")

	// Check the cache status after eviction
	if _, ok := myCache.Get("key2"); ok {
		fmt.Println("key2 is still in the cache")
	} else {
		fmt.Println("key2 has been evicted")
	}

	if val, ok := myCache.Get("key3"); ok {
		fmt.Println("key3:", val)
	} else {
		fmt.Println("key3 not found")
	}

	if val, ok := myCache.Get("key1"); ok {
		fmt.Println("key1:", val)
	} else {
		fmt.Println("key1 has been evicted")
	}
}
