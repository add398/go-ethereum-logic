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
	"time"
)








func Count_MBT_Space(size int)  int64 {


	keys, value := makeAccounts(size)


	mbt := cache.New(5*time.Minute, 10*time.Minute)

	for i := 0; i < size; i++ {
		mbt.SetDefault(string(keys[i]), value)
	}


}




func main() {
	times := make([]int64, 5)
	times[0] = Search_In_MBT_1w(500000)
	times[1] = Search_In_MBT_1w(1000000)
	times[2] = Search_In_MBT_1w(1500000)
	times[3] = Search_In_MBT_1w(2000000)
	times[4] = Search_In_MBT_1w(2500000)

	fmt.Println(times)

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