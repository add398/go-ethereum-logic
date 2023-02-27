/**
 * @Author: Log1c
 * @Description:
 * @File:  lab
 * @Version: 1.0.0
 * @Date: 2023/2/27 20:49
 */

package main

import (
	"fmt"
	art "github.com/plar/go-adaptive-radix-tree"
	"math/rand"
	"time"
)



func build_ART(size int, keys [][]byte, value []byte)  int64 {
	//keys, value := makeAccounts(size)

	start := time.Now() // 获取当前时间

	tree := art.New()
	for j := 0; j < size; j++ {
		tree.Insert(keys[j],value)
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
	build_ART(size, keys, value)
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
