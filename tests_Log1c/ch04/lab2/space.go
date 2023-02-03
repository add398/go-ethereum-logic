/**
 * @Author: Log1c
 * @Description:
 * @File:  space
 * @Version: 1.0.0
 * @Date: 2023/2/3 22:51
 */

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/tests_Log1c/ch04/lab2/ART_LRU"
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

func ART_Size(size int, keys [][]byte, value []byte) {
	// ART 空间
	tree := art.New()
	for i := 0; i < size; i++ {
		tree.Insert(art.Key(keys[i]), value)
	}

}


func B_ART_Size(size int, keys [][]byte, value []byte) {
	//  B ART 空间

	tree := ART_LRU.New()
	for i := 0; i < size; i++ {
		tree.Insert(ART_LRU.Key(keys[i]), value)
	}


}





func main() {
	size := 10000000
	size1 := 10000000
	keys, value := makeAccounts(size)
	//ART_Size(size1, keys, value)
	B_ART_Size(size1, keys, value)
	fmt.Println("over ")
	time.Sleep(1 * time.Hour)
}










//  确认  LRU 可以使用
func Test_ART_LRU() {
	tree := ART_LRU.New()
	terms := []string{"1", "2", "3", "4", "5"}

	for _, term := range terms {
		tree.Insert(ART_LRU.Key(term), term)
	}
	tree.ForEachByLeaf()
	//fmt.Println(tree)
	tree.Insert(ART_LRU.Key("1"), "1")
	tree.ForEachByLeaf()

	tree.Delete(ART_LRU.Key("1"))

	tree.ForEachByLeaf()
	//fmt.Println(tree)
}


