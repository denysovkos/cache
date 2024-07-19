# LRU Cache Implementation

This repository contains a generic, reusable, and easy-to-integrate LRU (Least Recently Used) cache module implemented in Go. This cache module is designed to improve application performance by holding heavily accessed (read/written) application-specific objects.

## Features

- **Generic and Reusable**: The cache module can be easily integrated across various modules within your codebase or organization.
- **Fixed Capacity**: The cache is bounded by a fixed capacity specified during initialization.
- **Eviction Strategies**: Implements the Least Recently Used (LRU) eviction strategy to make room for newer objects when the cache reaches its capacity.
- **Concurrency Safe**: The cache module is thread-safe, ensuring safe access and modifications across multiple goroutines.
- **String Keys**: Simplified usage with string keys for storing and retrieving values.

## Inspiration

This implementation is inspired by the [Redis LRU Cache](https://redis.io/glossary/lru-cache/), which is a widely used caching strategy in Redis to manage memory by evicting the least recently used keys.

## Functional Requirements

1. **Generic, Reusable, and Easy to Integrate**: The cache module is designed to be easily integrated and reused across different parts of an application.
2. **Fixed Capacity**: The cache holds a fixed number of objects, specified at the time of initialization.
3. **Eviction Strategy**: Implements the LRU eviction strategy to remove the least recently used items when the cache reaches its capacity.
4. **Concurrency**: The cache is thread-safe, allowing concurrent access and modifications.
5. **String Keys**: Simplifies key management by using string keys.

## Non-functional Requirements

1. **Production-grade Implementation**: The cache module is designed with a judicious mix of code modularity, extensibility, and test coverage to ensure production readiness.
2. **No Third-party Libraries**: The implementation does not use any third-party libraries, relying solely on the Go standard library.
3. **Modularity and Extensibility**: The code is modular and extensible, making it easy to add new features or modify existing ones.

## Usage

### Initializing the Cache

To initialize the cache, use the `NewCache` function and specify the capacity:

```go
import "semrush/cache/cache"

cache, err := cache.NewCache(2)
if err != nil {
    log.Fatal(err)
}
```