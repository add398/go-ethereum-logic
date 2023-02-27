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






func Search_In_MAP()  {
	size := 25000000

	keys, value := makeAccounts(size)

	var Search_1w func(mbt *cache.Cache,  curSize int )
	Search_1w = func(mbt *cache.Cache, curSize int) {
		search_size := 10000
		fra := curSize / search_size  //  1000w / 1w = 1000
		count := 0
		for i := 0; i < curSize; i++ {
			if i % fra == 0 {
				count++
				mbt.Get(string(keys[i]))
			}
		}
		fmt.Println(count)
	}

	sizeNum := []int{5000000, 10000000, 15000000, 20000000, 25000000}
	timeSize := 5
	times := make([]int64, timeSize)
	timeCount := 0

	for i := 0; i < 5; i++ {
		mbt := cache.New(5*time.Minute, 10*time.Minute)
		for j := 0; j < sizeNum[i]; j++ {
			mbt.SetDefault(string(keys[i]), value)
		}
		fmt.Println(sizeNum[i])   //
		start := time.Now() // 获取当前时间
		Search_1w(mbt, sizeNum[i])

		elapsed := time.Since(start)
		fmt.Println("该函数执行完成耗时：", elapsed)
		times[timeCount] = elapsed.Microseconds()
		timeCount++
	}
	fmt.Println(times)
}




func Search_In_MBT_1w(size int)  int64 {

	search_size := 10000
	keys, value := makeAccounts(size)
	fra := size / search_size  //  500w / 1w = 500

	count := 0

	mbt := cache.New(5*time.Minute, 10*time.Minute)

	for i := 0; i < size; i++ {
		mbt.SetDefault(string(keys[i]), value)
	}

	start := time.Now() // 获取当前时间

	for i := 0; i < size; i++ {
		if i % fra == 0 {
			count++
			mbt.Get(string(keys[i]))
		}
	}

	elapsed := time.Since(start)
	fmt.Println("size = ", size)
	fmt.Println("执行完成耗时：", elapsed)
	timeNum := elapsed.Microseconds()   //   us
	fmt.Println("查询交易" ,timeNum, "us")
	fmt.Println()
	return timeNum
}


func main() {
	times := make([]int64, 8)
	times[0] = Search_In_MBT_1w(1000000)
	times[1] = Search_In_MBT_1w(2000000)
	times[2] = Search_In_MBT_1w(3000000)
	times[3] = Search_In_MBT_1w(4000000)
	times[4] = Search_In_MBT_1w(5000000)
	times[5] = Search_In_MBT_1w(10000000)
	times[6] = Search_In_MBT_1w(15000000)
	times[7] = Search_In_MBT_1w(20000000)

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