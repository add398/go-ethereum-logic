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

func Search_In_MPT()  {

	size := 20000000

	keys, value := makeAccounts(size)
	var Search_1w func(tree *Trie,  curSize int )
	Search_1w = func( tree *Trie, curSize int) {
		search_size := 10000
		fra := curSize / search_size  //  1000w / 1w = 1000
		count := 0
		for i := 0; i < curSize; i++ {
			if i % fra == 0 {
				count++
				tree.Get(keys[i])
			}
		}
		fmt.Println(count)
	}

	sizeNum := []int{5000000, 10000000, 15000000, 20000000}
	timeSize := 4
	times := make([]int64, timeSize)
	timeCount := 0

	for i := 0; i < 4; i++ {
		tree := NewTrie()
		for j := 0; j < sizeNum[i]; j++ {
			tree.Put(keys[i],value)
		}
		fmt.Println(size)   //
		start := time.Now() // 获取当前时间
		Search_1w(tree, sizeNum[i])

		elapsed := time.Since(start)
		fmt.Println("该函数执行完成耗时：", elapsed)
		times[timeCount] = elapsed.Microseconds()
		timeCount++
	}
	fmt.Println(times)
}




func main() {





}
