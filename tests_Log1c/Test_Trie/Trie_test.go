/**
 * @Author: Log1c
 * @Description:
 * @File:  trie_test
 * @Version: 1.0.0
 * @Date: 2022/10/22 15:28
 */

package Test_Trie

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"math/rand"
	"testing"
)

func Test_Trie_Store(t *testing.T) {
	diskdb := memorydb.New()
	triedb := trie.NewDatabase(diskdb)
	random := rand.New(rand.NewSource(0))

	size := 10000000
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	fmt.Println(1)

	tree := trie.NewEmpty(triedb)
	for i := 0; i < size; i++ {

		tree.Update(keys[i], []byte("1"))
	}

	fmt.Println(2)

	fmt.Println(3)



}

func Benchmark_Test_Trie(b *testing.B) {
	diskdb := memorydb.New()
	triedb := trie.NewDatabase(diskdb)
	random := rand.New(rand.NewSource(0))

	size := 20000000
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	tree := trie.NewEmpty(triedb)
	for i := 0; i < size; i++ {

		tree.Update(keys[i], []byte("1"))
	}

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < size; i++ {
			if i % 2000 == 0 {
				tree.TryGet(keys[i])
			}
			//val, _ := tree.TryGet(keys[i])
			//fmt.Println(keys[i])
			//fmt.Println(val)
		}
	}
	b.StopTimer()
}

