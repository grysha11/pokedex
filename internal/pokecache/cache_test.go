package pokecache

import (
	"reflect"
	"testing"
	"time"
	"fmt"
	"sync"
)

func TestAddGetCache(t *testing.T) {
	const interval = 30 * time.Second
	cache := NewCache(interval)

	tests := map[string]struct {
		key		string
		val		[]byte
	}{
		"validTest1":		{key: "test1", val: []byte("test1")},
		"validTest2":	{key: "tesofgdakfp kdsakopfko sdakofpkosadkopf kosdpa fkopsad", val: []byte("tesofgdakfp kdsakopfko sdakofpkosadkopf kosdpa fkopsad")},
	}
	
	for test, tc := range tests {
		t.Run(test, func(t *testing.T) {
			cache.Add(tc.key, tc.val)

			got, ok := cache.Get(tc.key)
			if !ok {
				t.Fatalf("expected cache hit for key '%s', but got miss", tc.key)
			}

			if !reflect.DeepEqual(got, tc.val) {
				t.Fatalf("expected value '%s', but got '%s'", string(tc.val), string(got))
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const reapInterval = 50 * time.Millisecond
	const testWaitTime = reapInterval + (20 * time.Millisecond)
	
	cache := NewCache(reapInterval)

	key := "reaptest"
	cache.Add(key, []byte("data"))

	_, ok := cache.Get(key)
	if !ok {
		t.Fatalf("key '%s' was not added successfully", key)
	}

	time.Sleep(testWaitTime)

	_, ok = cache.Get(key)
	if ok {
		t.Fatalf("key '%s' was not reaped after %v", key, testWaitTime)
	}
}

func TestConcurrency(t *testing.T) {
	cache := NewCache(10 * time.Second)

	var wg sync.WaitGroup
	numRoutines := 100

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			key := fmt.Sprintf("key-%d", i)
			val := []byte(fmt.Sprintf("val-%d", i))
			cache.Add(key, val)

			cache.Get(key)
		}(i)
	}

	wg.Wait()
}