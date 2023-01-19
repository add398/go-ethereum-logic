/**
 * @Author: Log1c
 * @Description:
 * @File:  ART_in_disk_test
 * @Version: 1.0.0
 * @Date: 2023/1/18 15:50
 */

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
)

func Store_address(dir string, size int, keys [][]byte, value []byte)  (common.Hash) {
	dbase, err := leveldb.New(dir,8000,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()



	triedb := trie.NewDatabase(dbase)
	tree := trie.NewEmpty(triedb)


	tree.Update(keys[0], value)

	fmt.Println("start  ")

	root, nodes, _ := tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)
	tree, _ = trie.New(trie.TrieID(root), triedb)

	for i := 1; i < size; i++ {

		v := tree.Get(keys[i])
		if v == nil {
			tree.Update(keys[i], value)
		} else {
			v = ValueUpdate(v)
			tree.Update(keys[i], v)
		}


		if i % 200000 == 999 {
			root, nodes, _ = tree.Commit(false)
			triedb.Update(trie.NewWithNodeSet(nodes))
			triedb.Commit(root, true, nil)
			tree, _ = trie.New(trie.TrieID(root), triedb)
			fmt.Println(i)
			fmt.Println(root)

		}


	}

	root, nodes, _ = tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)
	return root
}

func Get_TrieDB(root common.Hash, dir  string,  size int,  keys [][]byte, value []byte ) {
	dbase, err := leveldb.New(dir,8000,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	triedb := trie.NewDatabase(dbase)
	tree, _ := trie.New(trie.TrieID(root), triedb)

	count := size / 100

	for j := 0; j < size; j++ {
		if j % count == 0 {
			v, _ := tree.TryGet(keys[j])
			fmt.Println(v)
		}

	}
}

func Store8kw() {
	size := 80000000
	dir :=  "store.logfile"

	keys, value := get_address(size)
	// ---------------------------  获取  key ，value

	root := Store_address( dir, size, keys, value)
	fmt.Println(root)
	// 0x5530b04f623d58d4e9fa27599f771a6044574f37d730b4bb2b44740d87795fe8

}

func Get_100() {
	str := "0x5530b04f623d58d4e9fa27599f771a6044574f37d730b4bb2b44740d87795fe8"
	size := 800000
	dir :=  "store.logfile"

	keys, value := get_address(size)

	root := common.BytesToHash(common.FromHex(str))
	fmt.Println(root)

	Get_TrieDB(root, dir, size, keys, value)
}

//0x3cc6dc79df5c02abfb2b7f76e5237a0245e7abd9f949722b1bbdfac5d099c767
//79600999
//0x57637fd61fd5e3af0bc65aa46e74451b468c0094d075c3f864d126fada224261
//79800999
//0x07a13c4b745536f140582a65663d45f4558f7be6308df2febeb6b529d04c325c
//0x5530b04f623d58d4e9fa27599f771a6044574f37d730b4bb2b44740d87795fe8




func main() {
	mpt_8kw_2kw_2w()

}
