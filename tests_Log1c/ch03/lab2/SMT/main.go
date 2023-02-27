/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/2/27 15:08
 */

package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/celestiaorg/smt"
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


func Search_In_MPT_1w(size int)  int64 {

	search_size := 10000
	keys, value := makeAccounts(size)
	fra := size / search_size  //  500w / 1w = 500

	count := 0

	nodeStore := smt.NewSimpleMap()
	valueStore := smt.NewSimpleMap()
	// Initialise the tree
	smtree := smt.NewSparseMerkleTree(nodeStore, valueStore, sha256.New())


	for j := 0; j < size; j++ {
		smtree.Update(keys[j],value)
	}

	start := time.Now() // 获取当前时间

	for i := 0; i < size; i++ {
		if i % fra == 0 {
			count++
			smtree.Get(keys[i])
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
	//times[0] = Search_In_MPT_1w(500000)
	times[1] = Search_In_MPT_1w(1000000)
	times[2] = Search_In_MPT_1w(5000000)
	//times[3] = Search_In_MPT_1w(2000000)
	//times[4] = Search_In_MPT_1w(2500000)

	fmt.Println(times)

}





func main1() {
	// Initialise two new key-value store to store the nodes and values of the tree
	nodeStore := smt.NewSimpleMap()
	valueStore := smt.NewSimpleMap()
	// Initialise the tree
	tree := smt.NewSparseMerkleTree(nodeStore, valueStore, sha256.New())

	// Update the key "foo" with the value "bar"
	_, _ = tree.Update([]byte("foo"), []byte("bar"))
	value,_ := tree.Get([]byte("foo"))

	fmt.Println(value)
}
