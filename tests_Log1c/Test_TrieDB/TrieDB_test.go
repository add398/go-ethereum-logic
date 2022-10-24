/**
 * @Author: Log1c
 * @Description:
 * @File:  trie_test
 * @Version: 1.0.0
 * @Date: 2022/10/20 17:53
 */

package Test_TrieDB

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
	"math/rand"
	"testing"
)





func TestStore_TrieDB(t *testing.T) {
	dbase, err := leveldb.New("1.logfile",8,500,"cc",false)

	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	random := rand.New(rand.NewSource(0))
	size := 10
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
		fmt.Println(k)
	}


	triedb := trie.NewDatabase(dbase)
	tree := trie.NewEmpty(triedb)


	for i :=0; i < size; i++ {
		tree.Update(keys[i], []byte("1qwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnm"))
	}

	root, nodes, _ := tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)

	fmt.Println(root)

}



func Test_RootGet(t *testing.T) {
	s := "0x047db56a3a524930b55b1375a05d3bd39fc952c6c5d2599a9c3fcc88ac90e196"
	root := common.BytesToHash(common.FromHex(s))
	fmt.Println(root)
	dbase, err := leveldb.New("1.logfile",8,500,"cc",false)

	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	random := rand.New(rand.NewSource(0))
	size := 10
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}


	triedb := trie.NewDatabase(dbase)
	tree, _ := trie.New(trie.TrieID(root), triedb)
	for i := 0; i < size; i++ {
		val, _ := tree.TryGet(keys[i])
		fmt.Println(val)

	}

}

func Benchmark_Get_Value_From_TrieDB(b *testing.B) {

	s := "0xe75427e9da29b7eeabc44c2e1714e00fc02f9040235cda92f609447eefd420a3"
	root := common.BytesToHash(common.FromHex(s))
	//fmt.Println(root)
	dbase, err := leveldb.New("store.logfile",8,500,"cc",false)

	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	random := rand.New(rand.NewSource(0))
	size := 1000000
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}


	triedb := trie.NewDatabase(dbase)
	tree, _ := trie.New(trie.TrieID(root), triedb)


	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {

			if j % 100 == 0 {
				tree.TryGet(keys[j])
			}

		}
	}
	b.StopTimer()
}


func test_Store_TrieDB(t *testing.T, size int, dir string) {

	dbase, err := leveldb.New(dir,8,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	random := rand.New(rand.NewSource(0))

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	triedb := trie.NewDatabase(dbase)
	tree := trie.NewEmpty(triedb)

	for i :=0; i < size; i++ {
		tree.Update(keys[i], []byte("1qwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnm"))
	}
	root, nodes, _ := tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)

	fmt.Println(root)

}




func benchmark_Get_TrieDB(b *testing.B, size int, dir string, str string) {
	root := common.BytesToHash(common.FromHex(str))
	//fmt.Println(root)
	dbase, err := leveldb.New(dir,8,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	random := rand.New(rand.NewSource(0))

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}


	triedb := trie.NewDatabase(dbase)
	tree, _ := trie.New(trie.TrieID(root), triedb)

	count := size / 10000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			if j % count == 0 {
				tree.TryGet(keys[j])
			}

		}
	}
	b.StopTimer()
}


