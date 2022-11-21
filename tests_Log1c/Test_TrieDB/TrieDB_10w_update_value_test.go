/**
 * @Author: Log1c
 * @Description:
 * @File:  LevelDB_100w_test
 * @Version: 1.0.0
 * @Date: 2022/10/24 11:55
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

func Test_start_update_100w(t *testing.T) {
	size := 1000000
	dir := "update.logfile"

	value := "1qwertyuioplkjhgfdsazxcuioplkjhgfdsazxcvbnm"
	fmt.Println(value)
	//value += "1"
	//fmt.Println(value)


	test_store_TrieDB( size, dir,value)

}

func Test_update_100w(t *testing.T) {
	size := 1000000
	dir := "update.logfile"
	root := "0xca4625ed1bb8092d1bc3557cabdde4f28963ffc201f3551d127d730e69ffae5b"
	value := "ujf1de2s13214562103190poi1ytrgayidghu99"

	test_get_and_update_TrieDB( size, dir, root, value)

}

func Test_update2_100w_multi10(t *testing.T) {
	size := 1000000
	dir := "update.logfile"
	root := "0x92ca9767f3ea001ac9fd70ce2a1418b070940f32eaab4b98ce737e974037fedc"
	value := "ujf1de2jkls621a310poi1ytrgaydghu"

	num := 6

	for i := 0; i < num; i++ {
		value = value + "1"
		root = test_get_and_update_TrieDB( size, dir, root, value)
	}

}



func test_store_TrieDB( size int, dir string, value string) {

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
		tree.Update(keys[i], []byte(value))
	}
	root, nodes, _ := tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)

	fmt.Println(root)

}



func test_get_and_update_TrieDB(size int, dir string, str_root string, value string) string {
	root := common.BytesToHash(common.FromHex(str_root))
	fmt.Println(root)
	dbase, err := leveldb.New(dir, 8, 500, "cc", false)
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

	for i :=0; i < size; i++ {
		tree.Update(keys[i], []byte(value))
	}
	root, nodes, _ := tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)

	root2 := fmt.Sprint(root)
	fmt.Println("root", root)
	fmt.Println("root2", root2)

	return root2

}


func Benchmark_read_FROM_TrieDB_update_100w(b *testing.B) {
	size := 1000000
	dir := "update.logfile"
	str := "0xca4625ed1bb8092d1bc3557cabdde4f28963ffc201f3551d127d730e69ffae5b"
	benchmark_Get_TrieDB(b, size, dir, str)
}