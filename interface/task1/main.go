package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

var (
	ErrKeyNotFound  = errors.New("key not found")
	ErrTypeMismatch = errors.New("type mismatch")
)

type Cache struct {
	mu      sync.RWMutex
	entries map[string]Entry
}

type Entry struct {
	value     interface{}
	expiresAt time.Time
}

func NewCache() *Cache {
	return &Cache{
		entries: make(map[string]Entry),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.entries[key] = Entry{value, time.Now().Add(ttl)}
	c.mu.Unlock()
}
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()

	v, ok := c.entries[key]
	if !ok {
		c.mu.RUnlock()
		return nil, false
	}

	expAt := v.expiresAt
	if time.Now().Before(v.expiresAt) {
		c.mu.RUnlock()
		return v.value, true
	}
	c.mu.RUnlock()
	c.mu.Lock()
	if cur, ok := c.entries[key]; ok && cur.expiresAt.Equal(expAt) {
		delete(c.entries, key)
	}
	c.mu.Unlock()
	return nil, false
}

func (c *Cache) Exists(key string) bool {
	_, ok := c.Get(key)
	return ok
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.entries, key)
	c.mu.Unlock()
}

func (c *Cache) Clear() {
	c.mu.Lock()
	clear(c.entries)
	c.mu.Unlock()
}

func (c *Cache) ToJSON() ([]byte, error) {
	data := make(map[string]interface{})
	now := time.Now()
	c.mu.RLock()
	for k, v := range c.entries {
		if now.Before(v.expiresAt) {
			data[k] = v.value
		}
	}
	c.mu.RUnlock()
	return json.Marshal(data)
}

func (c *Cache) GetAs[T any](key string) (T, error) {
	var zero T
	v, ok := c.Get(key)
	if !ok {
		return zero, fmt.Errorf("%w: %s", ErrKeyNotFound, key)
	}
	retValue, ok := v.(T)
	if !ok {
		return zero, fmt.Errorf("%w: cannot convert value of type %T to type %s",
			ErrTypeMismatch, v, reflect.TypeFor[T]())
	}
	return retValue, nil
}

func main() {
	cache := NewCache()
	cache.Set("a", 1, time.Second*60)
	cache.Set("b", 2, time.Second*60)
	fmt.Println(cache.Get("a"))

	fmt.Println(cache.Exists("b"))
	jsonData, err := cache.ToJSON()
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(string(jsonData))
	fmt.Println(cache.GetAs[[]int]("a"))

	cache.Delete("b")
	fmt.Println(cache.Get("b"))
	fmt.Println(cache.Get("a"))
	cache.Set("c", 3, time.Millisecond*150)
	fmt.Println(cache.Get("c"))
	time.Sleep(time.Millisecond * 200)
	fmt.Println(cache.Get("c"))

	cache.Clear()
	fmt.Println(cache.entries)

}
