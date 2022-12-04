/**
 * @Author: Log1c
 * @Description:
 * @File:  Trie_3_test
 * @Version: 1.0.0
 * @Date: 2022/12/4 00:17
 */

package Test_Trie

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/trie"
	"testing"
)

func Benchmark_3_10000w(b *testing.B) {
	size := 40000000
	keys, val := makeAccounts(size)
	for i := 0; i < b.N; i++ {
		benchmark_3_trie(b, keys, val)
	}

	fmt.Println("over")
}



func benchmark_3_trie(b *testing.B, keys [][]byte, value []byte) {
	b.ReportAllocs()

	size := len(keys)
	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	tree := trie.NewEmpty(trie.NewDatabase(rawdb.NewMemoryDatabase()))

	b.StartTimer()
	for i := 0; i < size; i++ {
		if i % 1000000 == 0 {
			fmt.Println( i )
		}
		tree.Update(keys[i], value)
	}
	b.StopTimer()
}


