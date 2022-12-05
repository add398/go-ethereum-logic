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
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"testing"
)



func Benchmark_read_from_memorydb_5kw(b *testing.B) {
	// 2500w 交易， 5kw个地址， 从该树中 执行 1w笔交易所需时间
	size := 50000000
	keys, val := makeAccounts(size)


	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	triedb := trie.NewDatabase(memorydb.New())
	tree := trie.NewEmpty(triedb)

	for i := 0; i < size; i++ {
		tree.Update(keys[i], val)
		if i % 5000 == 4999 {
			root, nodes, _ := tree.Commit(false)
			triedb.Update(trie.NewWithNodeSet(nodes))
			triedb.Commit(root, true, nil)
			tree, _ = trie.New(trie.TrieID(root), triedb)
			fmt.Println(i)
		}
	}

	for i := 0; i < b.N; i++ {
		read_from_memorydb_5kw(b, tree, keys, val)
	}

	fmt.Println("over")
}

func ValueUpdate(value []byte) []byte {
	// 改变 value
	n := len(value)
	for i := 0; i < n; i++ {
		if value[i] < byte(255) {
			value[i] = value[i] + byte(1)
			break
		}else{
			// == 255
			value[i] = byte(0)
		}
	}
	return value
}



func read_from_memorydb_5kw(b *testing.B, tree *trie.Trie,  keys [][]byte, value []byte) {
	size := len(keys)
	count := size / 10000
	b.StartTimer()
	for i := 0; i < size; i++ {
		if i % count == 1 {
			v1 := tree.Get(keys[i])
			v2 := tree.Get(keys[i+1])
			v1 = ValueUpdate(v1)
			v2 = ValueUpdate(v2)

			tree.Update(keys[i],  v1)
			tree.Update(keys[i+1], v2)
		}
	}
	b.StopTimer()
}



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


