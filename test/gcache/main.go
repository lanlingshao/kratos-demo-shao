package main

import (
	"fmt"
	"github.com/lanlingshao/kratos-demo-shao/internal/storage/cache"
)

func main() {
	cacheCli := cache.NewLocalCacheClient(&cache.LocalCacheOption{Size: 3}, nil)
	cacheCli.Set("key1", 1)
	cacheCli.Set("key2", 2)
	cacheCli.Set("key3", 3)
	cacheCli.Set("key4", 4)
	v, err := cacheCli.Get("key1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

	cacheCli2 := cache.NewLocalCacheClient(&cache.LocalCacheOption{Size: 4}, nil)
	cacheCli2.Set("key1", 1)
	cacheCli2.Set("key2", 2)
	cacheCli2.Set("key3", 3)
	cacheCli2.Set("key4", 4)
	v, err = cacheCli2.Get("key1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
