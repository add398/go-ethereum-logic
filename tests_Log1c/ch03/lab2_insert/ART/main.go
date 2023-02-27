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



func build_ART(size int)  int64 {
	keys, value := makeAccounts(size)

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

func Insert_time() {
	times := make([]int64, 8)
	times[0] = build_ART(1000000)
	times[1] = build_ART(2000000)
	times[2] = build_ART(3000000)
	times[3] = build_ART(4000000)

	times[4] = build_ART( 5000000)
	times[5] = build_ART(10000000)
	times[6] = build_ART(15000000)
	times[7] = build_ART(20000000)

	fmt.Println(times)
	time.Sleep(1 * time.Hour)
}



func main() {
	Insert_time()

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

