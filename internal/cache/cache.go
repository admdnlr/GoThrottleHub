package cache

import (
	"sync"
	"time"
)

// CacheItem struct'ı, önbellek değeri ve son erişim zamanını saklar.
type CacheItem struct {
	Value      interface{}
	ExpiryTime time.Time
}

// Cache struct'ı, önbellek girişlerini saklar ve bir mutex ile korunur.
type Cache struct {
	items sync.Map
}

// NewCache, yeni bir önbellek örneği oluşturur.
func NewCache() *Cache {
	return &Cache{}
}

// Set, önbelleğe bir değer ekler. ttl saniye cinsinden yaşam süresidir.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	expiryTime := time.Now().Add(ttl)
	item := CacheItem{
		Value:      value,
		ExpiryTime: expiryTime,
	}
	c.items.Store(key, item)
}

// Get, önbellekten bir değer alır. Değer mevcut ve süresi dolmamışsa geri döndürülür.
func (c *Cache) Get(key string) (interface{}, bool) {
	result, ok := c.items.Load(key)
	if !ok {
		return nil, false
	}

	// Önbellek girişini CacheItem olarak tip dönüşümü yap
	item := result.(CacheItem)
	if time.Now().After(item.ExpiryTime) {
		c.items.Delete(key)
		return nil, false
	}

	return item.Value, true
}

// Delete, önbellekten bir anahtarın değerini siler.
func (c *Cache) Delete(key string) {
	c.items.Delete(key)
}

// Cleanup, süresi dolmuş önbellek girişlerini temizler.
func (c *Cache) Cleanup() {
	c.items.Range(func(key, value interface{}) bool {
		item := value.(CacheItem)
		if time.Now().After(item.ExpiryTime) {
			c.items.Delete(key)
		}
		return true
	})
}
