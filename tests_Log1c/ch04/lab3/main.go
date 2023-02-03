/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/2/4 02:05
 */

package main

import (
	"fmt"
	art "github.com/plar/go-adaptive-radix-tree"
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


func ART_Size(size int, keys [][]byte, value []byte, r int) {
	// ART 空间  with  r
	tree := art.New()
	for i := 0; i < size; i++ {
		if i % 1000000 == 0 {
			fmt.Println(keys[i][:r])
			fmt.Println(value[:32])
		}
		tree.Insert(art.Key(keys[i][:r]), value[:32])
	}

}

func main() {
	size := 10000000
	keys, value := makeAccounts(size)
	ART_Size(size, keys, value, 5)

	fmt.Println("over ")
	//B_ART_Size(size)
	time.Sleep(1 * time.Hour)
}




