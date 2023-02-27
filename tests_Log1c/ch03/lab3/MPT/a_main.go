package main

import (
	"fmt"
	"math/rand"
	"time"
)



func build_MPT(size int, keys [][]byte, value []byte)  int64 {

	start := time.Now() // 获取当前时间

	tree := NewTrie()
	for i := 0; i < size; i++ {
		tree.Put(keys[i], value)
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
	build_MPT(size, keys, value)
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







func MPT_Size(size int, keys [][]byte, value []byte) {
	// MPT 空间


	tree := NewTrie()
	for i := 0; i < size; i++ {
		tree.Put(keys[i], value)
	}

}




func main1() {
	size := 10000000
	size1 := 10000000
	keys, value := makeAccounts(size)
	MPT_Size(size1, keys, value)

	fmt.Println("over ")
	time.Sleep(1 * time.Hour)




}
