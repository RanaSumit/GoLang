package main
import (
	"fmt"
	"testing"
	"testify/assert" )

func TestCache(t *testing.T) {
	assert := assert.New(t)
	fmt.Println("Lab 1 - Part II \n---- LRU Cache ----")
	// Populates entries into the Cache
	l := New(CACHE_SIZE)
	l.Set(1, 10)
	l.Set(2, 20)
	l.Set(3, 30)

	assert.Equal(10, l.Get(1), "Get(1) should return 10")
	assert.Equal(20, l.Get(2), "Get(2) should return 20")
	assert.Equal(30, l.Get(3), "Get(3) should return 30")

	l.Set(4, 40)
	assert.Equal(40, l.Get(4), "Get(4) should return 40")
	// Checks Cache Invalidation
	assert.Equal(-1, l.Get(1), "Get(1) should return -1 as it was the least recently used entry")
}

