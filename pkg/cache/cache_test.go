package cache_test

import (
	"reflect"
	"testing"

	"github.com/dSquadAdmin/cache/pkg/cache"
)

type Assertion struct {
	actual interface{}
	t      *testing.T
}

func (a Assertion) Equals(expected interface{}) {
	if reflect.TypeOf(a.actual) != reflect.TypeOf(expected) || a.actual != expected {
		a.t.Fatalf("Expected: %v, Actual: %v", expected, a.actual)
	}
}

//TODO: Implement other methods such as HasKey, Contains etc on Assertion struct

func Expect(t *testing.T, actual interface{}) Assertion {
	return Assertion{t: t, actual: actual}
}

func Test_Cache(t *testing.T) {
	cache := cache.NewCache(2)

	t.Log("Cache size should not grow beyond the capacity")
	{
		cache.Put("one", 1)
		cache.Put("two", 2)
		cache.Put("three", 3)
		cache.Put("four", 4)
		Expect(t, cache.Size()).Equals(2)
		Expect(t, cache.IsFull()).Equals(true)
	}

	t.Log("Cache should push old records toward tail")
	{
		cache.Purge()
		cache.Put("one", 1)
		cache.Put("two", 2)
		Expect(t, cache.Size()).Equals(2)
		Expect(t, cache.Cache()).Equals(`{"capacity": 2, "size": 2, "isFull": true, data:[{ "two": "2" },{ "one": "1" }]}`)
	}

	t.Log("Cache size should move accessed data toward head")
	{
		cache.Purge()
		cache.Put("one", 1)
		cache.Put("two", 2)
		cache.Get("one")
		Expect(t, cache.Size()).Equals(2)
		Expect(t, cache.Cache()).Equals(`{"capacity": 2, "size": 2, "isFull": true, data:[{ "one": "1" },{ "two": "2" }]}`)
	}

	t.Log("Cache size should remove oldest data incase of new data pushed after cache full")
	{
		cache.Purge()
		cache.Put("one", 1)
		cache.Put("two", 2)
		cache.Put("three", 3)
		Expect(t, cache.Size()).Equals(2)
		Expect(t, cache.Cache()).Equals(`{"capacity": 2, "size": 2, "isFull": true, data:[{ "three": "3" },{ "two": "2" }]}`)
	}

	t.Log("Cache.Delete should remove data and decrease cache size ")
	{
		cache.Purge()
		cache.Put("one", 1)
		cache.Put("two", 2)
		cache.Delete("one")
		Expect(t, cache.Size()).Equals(1)
		Expect(t, cache.Cache()).Equals(`{"capacity": 2, "size": 1, "isFull": false, data:[{ "two": "2" }]}`)
	}

	t.Log("Cache.Purge should reset size and remove all data")
	{
		cache.Put("one", 1)
		cache.Put("two", 2)
		cache.Purge()
		Expect(t, cache.Size()).Equals(0)
		Expect(t, cache.IsFull()).Equals(false)
	}

}
