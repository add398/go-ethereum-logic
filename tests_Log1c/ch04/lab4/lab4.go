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
	dbase, err := leveldb.New(dir,8,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()



	triedb := trie.NewDatabase(dbase)
	tree := trie.NewEmpty(triedb)


	tree.Update(keys[0], value)
	root, nodes, _ := tree.Commit(false)
	for i := 1; i < size; i++ {
		v := tree.Get(keys[i])
		if v == nil {
			tree.Update(keys[i], value)
		} else {
			v = ValueUpdate(v)
			tree.Update(keys[i], v)
		}


		if i % 20000 == 19999 {
			root, nodes, _ = tree.Commit(false)
			triedb.Update(trie.NewWithNodeSet(nodes))
			triedb.Commit(root, true, nil)
			tree, _ = trie.New(trie.TrieID(root), triedb)
			fmt.Println(i)

		}


	}



	root, nodes, _ = tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)

	fmt.Println(root)
	return root
}

func Get_TrieDB(root common.Hash, dir  string,  size int,  keys [][]byte, value []byte ) {
	//root := common.BytesToHash(common.FromHex(str))
	//fmt.Println(root)
	dbase, err := leveldb.New(dir,8,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()



	triedb := trie.NewDatabase(dbase)
	tree, _ := trie.New(trie.TrieID(root), triedb)

	count := size / 10000

	for j := 0; j < size; j++ {
		if j % count == 0 {
			tree.TryGet(keys[j])
		}

	}
}




func main() {
	size := 10000
	dir :=  "store.logfile"

	keys, value := get_address(size)
	// ---------------------------  获取  key ，value

	root := Store_address( dir, size, keys, value)
	fmt.Println(root)
	Get_TrieDB(root, dir, size, keys, value)


	//count_address(50000000)
}
