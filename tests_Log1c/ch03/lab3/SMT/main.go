/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/2/27 21:16
 */

package SMT

import (
	"crypto/sha256"
	"fmt"
	"github.com/celestiaorg/smt"
	"math/rand"
	"time"
)

func build_SMT(size int, keys [][]byte, value []byte)  int64 {

	start := time.Now() // 获取当前时间

	nodeStore := smt.NewSimpleMap()
	valueStore := smt.NewSimpleMap()
	// Initialise the tree
	smtree := smt.NewSparseMerkleTree(nodeStore, valueStore, sha256.New())


	for j := 0; j < size; j++ {
		smtree.Update(keys[j],value)
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
	build_SMT(size, keys, value)

	time.Sleep(10 * time.Minute)
}



func main() {
	Space(0)
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





