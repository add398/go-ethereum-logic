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

	s := "0x9665878971fd07f1e2af3387566057dc88c6842debadc841957213c88021a5a4"
	root := common.BytesToHash(common.FromHex(s))
	fmt.Println(root)
	dbase, err := leveldb.New("store.logfile",8,500,"cc",false)

	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	random := rand.New(rand.NewSource(0))
	size := 1000
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



	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {

			if j % 100 == 0 {
				tree.TryGet(keys[j])
			}

		}
	}
}