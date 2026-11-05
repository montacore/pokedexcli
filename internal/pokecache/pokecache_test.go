package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(time.Hour)
	
	cache.Add("key1",  []byte("value1"))
	value, ok := cache.Get("key1")
	if !ok {
		t.Error("expected key1 to exist")
	}

	if string(value) != "value1"{
		t.Errorf("expected 'value1', got %s", value)
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(20 * time.Millisecond)

	deadline := time.Now().Add(3*time.Second)
	cache.Add("some-key", []byte("some-value"))
	for time.Now().Before(deadline) {
		_, found := cache.Get("some-key")
		if !found {
			return
		}

		time.Sleep(5 * time.Millisecond)
	}
	
	t.Fatal("expected item to be removed")
}
