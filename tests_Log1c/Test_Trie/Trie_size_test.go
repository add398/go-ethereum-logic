/**
 * @Author: Log1c
 * @Description:
 * @File:  Trie_100w
 * @Version: 1.0.0
 * @Date: 2022/10/24 18:03
 */

package Test_Trie

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"math/rand"
	"testing"
)

func Benchmark_space(b *testing.B) {
	diskdb := memorydb.New()
	triedb := trie.NewDatabase(diskdb)
	random := rand.New(rand.NewSource(0))

	size := 10000

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	value := make([]byte, 100)
	random.Read(value)
	//fmt.Println(value)

	tree := trie.NewEmpty(triedb)

	for i := 0; i < size; i++ {
		tree.Update(keys[i], value)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < size; i++ {
		tree.Update(keys[i], value)
	}


	fmt.Println("over")
}











func Benchmark_read_from_trie_100w(b *testing.B) {
	size := 1000000
	benchmark_read_from_trie(b, size)
}

func Benchmark_read_from_trie_1000w(b *testing.B) {
	size := 20000000
	benchmark_read_from_trie(b, size)
}

func Benchmark_read_from_trie_10000w(b *testing.B) {
	size := 100000000
	benchmark_read_from_trie(b, size)
}



func benchmark_read_from_trie(b *testing.B, size int) {
	diskdb := memorydb.New()
	triedb := trie.NewDatabase(diskdb)
	random := rand.New(rand.NewSource(0))

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	value := make([]byte, 0)
	for i := 0; i < 5; i++ {
		value = append(value, keys[i]...)
	}
	//fmt.Println(value)

	tree := trie.NewEmpty(triedb)
	for i := 0; i < size; i++ {
		tree.Update(keys[i], []byte(""))
	}

	count := size / 10000

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < size; i++ {
			if i % count == 0 {
				tree.TryGet(keys[i])
			}

		}
	}
	b.StopTimer()
}

