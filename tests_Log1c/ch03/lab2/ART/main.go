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

func Search_In_ART()  {
	size := 25000000

	keys, value := makeAccounts(size)
	//keys1, _ := makeAccounts(10000)

	var Search_1w func(tree art.Tree,  curSize int )
	Search_1w = func( tree art.Tree, curSize int) {
		search_size := 10000
		fra := curSize / search_size  //  1000w / 1w = 1000
		count := 0
		for i := 0; i < curSize; i++ {
			if i % fra == 0 {
				count++
				tree.Search(keys[i])
			}
		}
		fmt.Println(count)
	}

	sizeNum := []int{5000000, 10000000, 15000000, 20000000, 25000000}
	timeSize := 5
	times := make([]int64, timeSize)
	timeCount := 0

	for i := 0; i < 5; i++ {
		tree := art.New()
		for j := 0; j < sizeNum[i]; j++ {
			tree.Insert(keys[i],value)
		}
		fmt.Println(sizeNum[i])   //
		start := time.Now() // 获取当前时间
		Search_1w(tree, sizeNum[i])

		elapsed := time.Since(start)
		fmt.Println("该函数执行完成耗时：", elapsed)
		times[timeCount] = elapsed.Microseconds()
		timeCount++
	}
	fmt.Println(times)
}



func Search_In_ART_1w(size int)  int64 {

	search_size := 10000
	keys, value := makeAccounts(size)
	fra := size / search_size  //  500w / 1w = 500

	count := 0

	tree := art.New()
	for j := 0; j < size; j++ {
		tree.Insert(keys[j],value)
	}

	start := time.Now() // 获取当前时间

	for i := 0; i < size; i++ {
		if i % fra == 0 {
			count++
			tree.Search(keys[i])
		}
	}

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
	timeNum := elapsed.Microseconds()   //   us
	fmt.Println("查询交易" ,timeNum, "us")
	fmt.Println("size = ", size)
	fmt.Println()
	return timeNum
}




func main() {
	times := make([]int64, 5)
	times[0] = Search_In_ART_1w(500000)
	times[1] = Search_In_ART_1w(1000000)
	times[2] = Search_In_ART_1w(1500000)
	times[3] = Search_In_ART_1w(2000000)
	times[4] = Search_In_ART_1w(2500000)

	fmt.Println(times)

}
