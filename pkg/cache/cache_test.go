package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheMain(t *testing.T) {
	// 示例1：启用定时清理，定期清理时间间隔为5秒
	cacheWithClean := NewCache(5 * time.Second)

	cacheWithClean.Set("key1", "value1", 15) // 设置过期时间为15秒
	cacheWithClean.Set("key2", "value2", 30) // 设置过期时间为30秒

	// 示例2：禁用定时清理
	cacheWithoutClean := NewCache(0)

	cacheWithoutClean.Set("key3", "value3", 60) // 设置过期时间为60秒

	// 等待一段时间，使cacheWithClean中的key1和key2的缓存过期
	time.Sleep(20 * time.Second)

	value, found := cacheWithClean.Get("key1")
	if found {
		fmt.Println("Value found in cacheWithClean:", value)
	} else {
		fmt.Println("Value not found in cacheWithClean")
	}

	value, found = cacheWithoutClean.Get("key3")
	if found {
		fmt.Println("Value found in cacheWithoutClean:", value)
	} else {
		fmt.Println("Value not found in cacheWithoutClean")
	}
}
