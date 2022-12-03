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



func Benchmark_test_space(b *testing.B) {
	size := 100000
	acc, val := makeAccounts(size)
	for i := 0; i < b.N; i++ {
		benchmark_space(b, acc, val)
	}
	// 参考  benchmarkCommitAfterHashFixedSize

	fmt.Println("over")
}


func benchmark_space(b *testing.B, addresses [][]byte, value []byte)  {
	b.ReportAllocs()
	b.StartTimer()

	trie := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	for i := 0; i < len(addresses); i++ {
		if i % 100 == 0 {
			fmt.Println(i )
		}
		trie.Update(addresses[i], value)
	}
	//v := trie.Get(addresses[0])
	//fmt.Println(v)
	//trie.Commit(false)

	// commit 以后不可以 get数据
	//v = trie.Get(addresses[0])
	//fmt.Println(v)

	b.StopTimer()
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





func Test_Trie_Store(t *testing.T) {
	diskdb := memorydb.New()
	triedb := trie.NewDatabase(diskdb)
	size := 100000
	keys, v := makeAccounts(size)

	fmt.Println(1)

	tree := trie.NewEmpty(triedb)
	for i := 0; i < size; i++ {
		tree.Update(keys[i], v)
	}

	fmt.Println(3)
}
