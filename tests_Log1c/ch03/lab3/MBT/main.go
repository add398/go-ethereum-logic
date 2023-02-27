/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/2/27 14:04
 */

package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"time"
)






func build_MBT(size int, keys [][]byte, value []byte)  int64 {

	start := time.Now() // 获取当前时间


	mbt := cache.New(5*time.Minute, 10*time.Minute)

	for i := 0; i < size; i++ {
		mbt.SetDefault(string(keys[i]), value)
	}


	elapsed := time.Since(start)
	timeNum := elapsed.Microseconds()   //   us
	fmt.Println("size = ", size)
	fmt.Println("Build 耗时：", elapsed)
	fmt.Println("Insert 耗时：" ,timeNum, "us")
	fmt.Println()
	return timeNum
}



// 可以测算 Insert 时间 和 空间大小


func Space(size int)  {
	keys, value := makeAccounts(10000000)
	build_MBT(size, keys, value)
	fmt.Println("build over")
	time.Sleep(10 * time.Minute)
}





func main() {
	Space(10000000)

}




func makeAccounts(size int) (addresses [][]byte, value []byte) {
	// Make the random benchmark deterministic
	random := rand.New(rand.NewSource(0))
	addresses = make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		addresses[i] = k
	}

	value = make([]byte, 100)
	random.Read(value)
	return
}






func main0() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
		fmt.Println(foo.(string))

	}

	foo, found = c.Get("baz")
	if found {
		fmt.Println(foo)
		fmt.Println(foo.(int))

	}




}