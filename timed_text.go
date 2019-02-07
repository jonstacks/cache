// Package cache provides various caching implementations/strategies.
package cache

import (
	"bytes"
	"sync"
	"time"
)

// TimedText is caches strings for a given amount of time. It also supports
// the ability to easily append text to a given key.
type TimedText struct {
	m           sync.RWMutex
	cache       map[string]*bytes.Buffer
	expireAfter *time.Duration
}

// NewTimedText initializes a new TimedText with the given expiration
// time. Pass nil for no expiration of the cache.
func NewTimedText(expireAfter *time.Duration) TimedText {
	return TimedText{
		cache:       make(map[string]*bytes.Buffer),
		expireAfter: expireAfter,
	}
}

// CreateOrAppend creates the given key with the given value if the key does
// not exist in the cache. Otherwise, it attempts to append the supplied value
// to the current value. It returns both a bool and an error. The bool value
// will be true if the key was created, otherwise false. If there was trouble
// writing the value to the internal buffer, an error will be returned.
func (ttc *TimedText) CreateOrAppend(key, value string) (created bool, err error) {
	ttc.m.Lock()
	defer ttc.m.Unlock()

	_, exists := ttc.cache[key]
	created = !exists

	if !exists {
		var buff bytes.Buffer
		buff.WriteString(value)
		ttc.cache[key] = &buff
		go ttc.createKeyExpiration(key)
		return
	}

	_, err = ttc.cache[key].WriteString(value)
	return
}

// CreateOrReplace creates the given key with the given value if the key does
// not exist in the cache. Otherwise, it replaces the value. It returns true
// if the key was created, otherwise false
func (ttc *TimedText) CreateOrReplace(key, value string) bool {
	var buff bytes.Buffer
	buff.WriteString(value)

	ttc.m.Lock()
	defer ttc.m.Unlock()

	_, exists := ttc.cache[key]
	ttc.cache[key] = &buff

	if !exists {
		go ttc.createKeyExpiration(key)
	}
	return !exists
}

// Get retrieves the value of the given key, returning an empty string if does
// not exist and a bool, true if the key exists, false otherwise.
func (ttc *TimedText) Get(key string) (value string, exists bool) {
	var buff *bytes.Buffer
	buff, exists = ttc.cache[key]
	if !exists {
		return
	}

	value = buff.String()
	return
}

// deletes the given key from the cache after the cache expiration time
func (ttc *TimedText) createKeyExpiration(key string) {
	if ttc.expireAfter == nil {
		return
	}

	time.Sleep(*ttc.expireAfter)

	ttc.m.Lock()
	defer ttc.m.Unlock()

	delete(ttc.cache, key)
}
