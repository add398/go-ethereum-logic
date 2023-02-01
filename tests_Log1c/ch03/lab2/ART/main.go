/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/2/1 13:45
 */

package main

import (
	"fmt"
	art "github.com/plar/go-adaptive-radix-tree"
	"time"
)





func main() {
	size := 5000000
	keys, value := makeAccounts(size)
	tree := art.New()

	timeSize := 5
	times := make([]int64, timeSize)
	timeCount := size / timeSize

	var Search_1w func( curSize int )
	Search_1w = func( curSize int) {
		search_size := 10000
		fra := curSize / search_size  //  1000w / 1w = 1000
		for i := 0; i < size; i++ {
			if i % fra == 0 {
				tree.Search(keys[i])
			}
		}
	}



	for i := 0; i < size; i++ {
		tree.Insert(keys[i], value)
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
