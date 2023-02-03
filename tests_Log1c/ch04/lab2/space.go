/**
 * @Author: Log1c
 * @Description:
 * @File:  space
 * @Version: 1.0.0
 * @Date: 2023/2/3 22:51
 */

package main

import (
	"github.com/ethereum/go-ethereum/tests_Log1c/ch04/lab2/ART_LRU"
	"math/rand"
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

func ART_LRU_Size(size int) {
	keys, value := makeAccounts(size)

}

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



func main() {

}



