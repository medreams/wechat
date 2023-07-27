package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data          map[string]cacheEntry
	mutex         sync.RWMutex
	cleanupTicker *time.Ticker
}

type cacheEntry struct {
	value      string
	expiration time.Time
}

var cacheInstance *Cache
var once sync.Once

// NewCache 创建一个新的缓存实例（单例模式）
func NewCache(cleanupInterval time.Duration) *Cache {
	once.Do(func() {
		cacheInstance = &Cache{
			data:          make(map[string]cacheEntry),
			cleanupTicker: nil,
		}
		if cleanupInterval > 0 {
			// 启动后台清理过期缓存的goroutine
			cacheInstance.startCleanup(cleanupInterval)
		}
	})
	return cacheInstance
}

// Set 将键值对设置到缓存中，并指定多少秒后过期
func (c *Cache) Set(key, value string, expirationSeconds int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expiration := time.Now().Add(time.Duration(expirationSeconds) * time.Second)
	c.data[key] = cacheEntry{
		value:      value,
		expiration: expiration,
	}
}

// Get 从缓存中获取给定键的值，如果不存在或已过期，返回空字符串和false
func (c *Cache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, found := c.data[key]
	if !found {
		return "", false
	}

	// 检查是否过期
	if c.cleanupTicker != nil && time.Now().After(entry.expiration) {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		delete(c.data, key)
		return "", false
	}

	return entry.value, true
}

// startCleanup 启动定期清理过期缓存的goroutine
func (c *Cache) startCleanup(interval time.Duration) {
	if c.cleanupTicker == nil {
		c.cleanupTicker = time.NewTicker(interval)
		go c.cleanupExpired()
	}
}

// cleanupExpired 定期清理过期的缓存项
func (c *Cache) cleanupExpired() {
	for range c.cleanupTicker.C {
		c.mutex.Lock()
		now := time.Now()
		for key, entry := range c.data {
			if now.After(entry.expiration) {
				delete(c.data, key)
			}
		}
		c.mutex.Unlock()
	}
}

// func main() {
// 	// 示例1：启用定时清理，定期清理时间间隔为5秒
// 	cacheWithClean := NewCache(5 * time.Second)

// 	cacheWithClean.Set("key1", "value1", 15) // 设置过期时间为15秒
// 	cacheWithClean.Set("key2", "value2", 30) // 设置过期时间为30秒

// 	// 示例2：禁用定时清理
// 	cacheWithoutClean := NewCache(0)

// 	cacheWithoutClean.Set("key3", "value3", 60) // 设置过期时间为60秒

// 	// 等待一段时间，使cacheWithClean中的key1和key2的缓存过期
// 	time.Sleep(20 * time.Second)

// 	value, found := cacheWithClean.Get("key1")
// 	if found {
// 		fmt.Println("Value found in cacheWithClean:", value)
// 	} else {
// 		fmt.Println("Value not found in cacheWithClean")
// 	}

// 	value, found = cacheWithoutClean.Get("key3")
// 	if found {
// 		fmt.Println("Value found in cacheWithoutClean:", value)
// 	} else {
// 		fmt.Println("Value not found in cacheWithoutClean")
// 	}
// }
