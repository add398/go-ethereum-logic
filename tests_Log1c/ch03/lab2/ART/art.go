/**
 * @Author: Log1c
 * @Description:
 * @File:  art
 * @Version: 1.0.0
 * @Date: 2023/2/1 13:11
 */

package main


import (
	"fmt"
	"time"

	art "github.com/plar/go-adaptive-radix-tree"
)

func DumpTree() {
	tree := art.New()
	terms := []string{"A", "a", "aa"}
	for _, term := range terms {
		tree.Insert(art.Key(term), term)
	}
	fmt.Println(tree)
}

func SimpleTree() {
	tree := art.New()

	tree.Insert(art.Key("Hi, I'm Key"), "Nice to meet you, I'm Value")
	value, found := tree.Search(art.Key("Hi, I'm Key"))
	if found {
		fmt.Printf("Search value=%v\n", value)
	}

	tree.ForEach(func(node art.Node) bool {
		fmt.Printf("Callback value=%v\n", node.Value())
		return true
	})

	for it := tree.Iterator(); it.HasNext(); {
		value, _ := it.Next()
		fmt.Printf("Iterator value=%v\n", value.Value())
	}
}





func main11111() {
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

