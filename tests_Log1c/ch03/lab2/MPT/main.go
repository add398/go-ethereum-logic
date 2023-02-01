package main

import (
	"fmt"
	"math/rand"

	"time"
)


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






func main() {
	size := 5000000
	keys, value := makeAccounts(size)
	tree := NewTrie()

	timeSize := 5
	times := make([]int64, timeSize)
	timeCount := size / timeSize

	var Search_1w func( curSize int )
	Search_1w = func( curSize int) {
		search_size := 10000
		fra := curSize / search_size  //  1000w / 1w = 1000
		for i := 0; i < size; i++ {
			if i % fra == 0 {
				tree.Get(keys[i])
			}
		}
	}



	for i := 0; i < size; i++ {
		tree.Put(keys[i], value)
		if i == timeCount - 1 {
			fmt.Println(i)   // 9999999
			start := time.Now() // 获取当前时间
			Search_1w(i + 1)

			elapsed := time.Since(start)
			//fmt.Println("该函数执行完成耗时：", elapsed1)
			times[i / timeCount] = elapsed.Microseconds()
		}else if i == 2 * timeCount - 1 {
			fmt.Println(i)
			start := time.Now() // 获取当前时间
			Search_1w(i + 1)
			elapsed := time.Since(start)
			//fmt.Println("该函数执行完成耗时：", elapsed1)
			times[i / timeCount] = elapsed.Microseconds()
		}else if i == 3 * timeCount - 1 {
			fmt.Println(i)   // 9999999
			start := time.Now() // 获取当前时间
			Search_1w(i + 1)

			elapsed := time.Since(start)
			//fmt.Println("该函数执行完成耗时：", elapsed1)
			times[i / timeCount] = elapsed.Microseconds()
		}else if i == 4 * timeCount - 1 {
			fmt.Println(i)   // 9999999
			start := time.Now() // 获取当前时间
			Search_1w(i + 1)

			elapsed := time.Since(start)
			//fmt.Println("该函数执行完成耗时：", elapsed1)
			times[i / timeCount] = elapsed.Microseconds()
		}else if i == 5 * timeCount - 1 {
			fmt.Println(i)   // 9999999
			start := time.Now() // 获取当前时间
			Search_1w(i + 1)

			elapsed := time.Since(start)
			//fmt.Println("该函数执行完成耗时：", elapsed1)
			times[i / timeCount] = elapsed.Microseconds()
		}
	}

	fmt.Println(times)


}
