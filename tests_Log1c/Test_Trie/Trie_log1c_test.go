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
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"math/rand"
	"testing"
	"time"
)



func Benchmark_test_space(b *testing.B) {
	size := 100000
	acc, val := makeAccounts(size)
	for i := 0; i < b.N; i++ {
		benchmark_space(b, acc, val)
	}
	// benchmarkCommitAfterHashFixedSize 参考

	//b.ReportAllocs()
	//trie := NewEmpty(NewDatabase(rawdb.NewMemoryDatabase()))
	//for i := 0; i < len(addresses); i++ {
	//	trie.Update(crypto.Keccak256(addresses[i][:]), accounts[i])
	//}
	//// Insert the accounts into the trie and hash it
	//trie.Hash()
	//b.StartTimer()
	//trie.Commit(false)
	//b.StopTimer()
}


func benchmark_space(b *testing.B, addresses [][]byte, value []byte)  {
	b.ReportAllocs()
	b.StartTimer()

	tree := art.New()
	for i := 0; i < len(addresses); i++ {
		tree.Insert(addresses[i], value)
	}

	b.StopTimer()
}







func test_art_space(testSize int)  {
	size := 500000
	address, value := makeAccounts(size)

	tree := art.New()
	for i := 0; i < testSize; i++ {
		tree.Insert(address[i], value)
	}

	time.Sleep(1 * time.Minute)
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
	random := rand.New(rand.NewSource(0))

	size := 100000
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}
	v := make([]byte, 100)
	random.Read(v)

	fmt.Println(1)

	tree := trie.NewEmpty(triedb)
	for i := 0; i < size; i++ {

		tree.Update(keys[i], v)
	}

	fmt.Println(2)

	fmt.Println(3)



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
	b.ReportAllocs()
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

