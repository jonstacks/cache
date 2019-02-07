package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimedTextGet(t *testing.T) {
	cache := NewTimedText(nil)

	value, exists := cache.Get("non-existent")

	assert.Equal(t, "", value, "TimedText.Get should return an empty string if the key does not exist.")
	assert.False(t, exists, "TimedText.Get should return false when the key does not exist.")

	cache.CreateOrAppend("key", "Some Value")
	value, exists = cache.Get("key")

	assert.Equal(t, "Some Value", value, "TimedText.Get should return the value if it exists.")
	assert.True(t, exists, "TimedText.Get should return true when the key exists")
}

func TestTimedTextCreateOrAppend(t *testing.T) {
	duration := 400 * time.Millisecond
	cache := NewTimedText(&duration)

	created, err := cache.CreateOrAppend("key", "Some Value\n")
	assert.True(t, created, "TimedText.CreateOrAppend should return true if the key was created.")
	assert.Nil(t, err)

	created, err = cache.CreateOrAppend("key", "More Value\n")
	assert.False(t, created, "TimedText.CreateOrAppend should return false if the key already exists.")
	assert.Nil(t, err)

	value, _ := cache.Get("key")
	assert.Equal(t, "Some Value\nMore Value\n", value, "TimedText.CreateOrAppend should append text if the key already exists.")

	time.Sleep(2 * duration)

	_, exists := cache.Get("key")
	assert.False(t, exists, "TimedText.CreateOrAppend should expire keys after the given duration.")
}

func TestTimedTextCreateOrReplace(t *testing.T) {
	duration := 400 * time.Millisecond
	cache := NewTimedText(&duration)

	created := cache.CreateOrReplace("key", "Some Value\n")
	assert.True(t, created, "TimedText.CreateOrReplace should return true if the key was created.")

	created = cache.CreateOrReplace("key", "New Value\n")
	assert.False(t, created, "TimedText.CreateOrReplace should return false if the key already exists.")

	value, _ := cache.Get("key")
	assert.Equal(t, "New Value\n", value, "TimedText.CreateOrReplace should replace text if the key already exists.")

	time.Sleep(2 * duration)

	_, exists := cache.Get("key")
	assert.False(t, exists, "TimedText.CreateOrReplace should expire keys after the given duration.")
}
