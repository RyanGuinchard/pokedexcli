package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {

	const interval = 5 * time.Second
	cache := NewCache(interval)

	key := "https://example.com"
	val := []byte("example value")

	cache.Add(key, val)

	retrievedVal, found := cache.Get(key)
	if !found {
		t.Errorf("Expected to find key %s", key)
	}
	if string(retrievedVal) != string(val) {
		t.Errorf("Expected value %s, got %s", val, retrievedVal)
	}
}

func TestReap(t *testing.T) {
	const interval = 1 * time.Second
	cache := NewCache(interval)

	key := "https://example.com"
	val := []byte("example value")

	cache.Add(key, val)

	// Wait for the reap interval to pass
	time.Sleep(2 * interval)

	_, found := cache.Get(key)
	if found {
		t.Errorf("Expected key %s to be reaped", key)
	}

}
