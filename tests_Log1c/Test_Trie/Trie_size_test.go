/**
 * @Author: Log1c
 * @Description:
 * @File:  Trie_100w
 * @Version: 1.0.0
 * @Date: 2022/10/24 18:03
 */

package Test_Trie

import (
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"math/rand"
	"testing"
)

func Test_random_key_value_space(t *testing.T) {
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
}











func Benchmark_read_from_trie_100w(b *testing.B) {
	size := 1000000
	benchmark_read_from_trie(b, size)
}

func Benchmark_read_from_trie_1000w(b *testing.B) {
	size := 10000000
	benchmark_read_from_trie(b, size)
}

func Benchmark_read_from_trie_10000w(b *testing.B) {
	size := 100000000
	benchmark_read_from_trie(b, size)
}