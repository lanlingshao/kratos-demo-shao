package main

import (
	"fmt"
	"github.com/bluele/gcache"
)

func main() {
	cache := gcache.New(4).Build()
	cache.Set("key1", 1)
	cache.Set("key2", 2)
	cache.Set("key3", 3)
	cache.Set("key4", 4)
	v, err := cache.Get("key1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

	cache2 := gcache.New(3).Build()
	cache2.Set("key1", 1)
	cache2.Set("key2", 2)
	cache2.Set("key3", 3)
	cache2.Set("key4", 4)
	v, err = cache2.Get("key1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
