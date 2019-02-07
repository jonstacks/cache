package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimedTextCacheGet(t *testing.T) {
	cache := NewTimedTextCache(nil)

	value, exists := cache.Get("non-existent")

	assert.Equal(t, "", value, "TimedTextCache.Get should return an empty string if the key does not exist.")
	assert.False(t, exists, "TimedTextCache.Get should return false when the key does not exist.")

	cache.CreateOrAppend("key", "Some Value")
	value, exists = cache.Get("key")

	assert.Equal(t, "Some Value", value, "TimedTextCache.Get should return the value if it exists.")
	assert.True(t, exists, "TimedTextCache.Get should return true when the key exists")
}

func TestTimedTextCacheCreateOrAppend(t *testing.T) {
	duration := 400 * time.Millisecond
	cache := NewTimedTextCache(&duration)

	created, err := cache.CreateOrAppend("key", "Some Value\n")
	assert.True(t, created, "TimedTextCache.CreateOrAppend should return true if the key was created.")
	assert.Nil(t, err)

	created, err = cache.CreateOrAppend("key", "More Value\n")
	assert.False(t, created, "TimedTextCache.CreateOrAppend should return false if the key already exists.")
	assert.Nil(t, err)

	value, _ := cache.Get("key")
	assert.Equal(t, "Some Value\nMore Value\n", value, "TimedTextCache.CreateOrAppend should append text if the key already exists.")

	time.Sleep(2 * duration)

	_, exists := cache.Get("key")
	assert.False(t, exists, "TimedTextCache.CreateOrAppend should expire keys after the given duration.")
}

func TestTimedTextCacheCreateOrReplace(t *testing.T) {
	duration := 400 * time.Millisecond
	cache := NewTimedTextCache(&duration)

	created := cache.CreateOrReplace("key", "Some Value\n")
	assert.True(t, created, "TimedTextCache.CreateOrReplace should return true if the key was created.")

	created = cache.CreateOrReplace("key", "New Value\n")
	assert.False(t, created, "TimedTextCache.CreateOrReplace should return false if the key already exists.")

	value, _ := cache.Get("key")
	assert.Equal(t, "New Value\n", value, "TimedTextCache.CreateOrReplace should replace text if the key already exists.")

	time.Sleep(2 * duration)

	_, exists := cache.Get("key")
	assert.False(t, exists, "TimedTextCache.CreateOrReplace should expire keys after the given duration.")
}
