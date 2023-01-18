/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/12/4 05:23
 */

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/trie"
	"math/rand"
	"testing"
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

func Benchmark_Find(b *testing.B) {

	size := 1000000
	keys, value := makeAccounts(size)

	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	tree := trie.NewEmpty(trie.NewDatabase(rawdb.NewMemoryDatabase()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			tree.Update(keys[j], value)
		}
	}
	b.StopTimer()

}





func Test_Find(t *testing.T) {
	size := 5000000
	keys, value := makeAccounts(size)

	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	tree := trie.NewEmpty(trie.NewDatabase(rawdb.NewMemoryDatabase()))

	for i := 0; i < size; i++ {
		if i % 1000000 == 0 {
			fmt.Println( i )
		}
		tree.Update(keys[i], value)
	}

}




func main() {
	size := 5000000
	keys, value := makeAccounts(size)
	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	tree := trie.NewEmpty(trie.NewDatabase(rawdb.NewMemoryDatabase()))
	for i := 0; i < size; i++ {
		if i % 1000000 == 0 {
			fmt.Println( i )
		}
		tree.Update(keys[i], value)
	}

	fmt.Println("over")
	time.Sleep(time.Hour)
}
