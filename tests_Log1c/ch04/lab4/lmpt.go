/**
 * @Author: Log1c
 * @Description:
 * @File:  lmpt
 * @Version: 1.0.0
 * @Date: 2023/1/19 14:54
 */

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
	"time"
)

// 已经存储了  8kw 地址，   在读取 2000w， commit 为 2000， 统计时间
func amt_8kw_2kw_2w()  {
	str := "0x5530b04f623d58d4e9fa27599f771a6044574f37d730b4bb2b44740d87795fe8"
	size := 20000000
	dir :=  "store.logfile"

	keys, value := get_address(size)

	root := common.BytesToHash(common.FromHex(str))
	fmt.Println("root: ", root)

	lmpt_Get_TrieDB(root, dir, size, keys, value, 2048, 2000)
}

func lmpt_Get_TrieDB(root common.Hash, dir  string,  size int,  keys [][]byte, value []byte , cache int, commitNum int) {
	dbase, err := leveldb.New(dir,cache,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	triedb := trie.NewDatabase(dbase)
	tree, _ := trie.New(trie.TrieID(root), triedb)


	TimeSize := 1000
	start1 := time.Now() // 获取当前时间
	start2 := time.Now() // 获取当前时间
	times := make([]int64, TimeSize)

	tree.Update(keys[0], value)
	root, nodes, _ := tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)
	tree, _ = trie.New(trie.TrieID(root), triedb)

	timeCount := size / TimeSize

	for i := 1; i < size; i++ {
		v := tree.Get(keys[i])
		if ValueCompare(v, value) == false {
			tree.Update(keys[i], value)
		} else {
			v = ValueUpdate(v)
			tree.Update(keys[i], v)
		}

		if i % commitNum == commitNum - 1 {
			root, nodes, _ = tree.Commit(false)
			triedb.Update(trie.NewWithNodeSet(nodes))
			triedb.Commit(root, true, nil)
			tree, _ = trie.New(trie.TrieID(root), triedb)

			if i % 10000 == commitNum - 1{
				fmt.Println(i)
				fmt.Println(root)
			}
		}
		if i % timeCount == timeCount - 1 {
			elapsed := time.Since(start2)
			//fmt.Println("该函数执行完成耗时：", elapsed1)
			times[i / timeCount] = elapsed.Milliseconds()
			start2 = time.Now() // 获取当前时间
		}
	}




	fmt.Println("Experiment over")
	elapsed := time.Since(start1)
	fmt.Println("执行完成耗时：", elapsed.Milliseconds())
	fmt.Println(times)
}






// 已经存储了  8kw 地址，   在读取 2000w， commit 为 20000， 统计时间
func mpt_8kw_2kw_2w()  {
	str := "0x5530b04f623d58d4e9fa27599f771a6044574f37d730b4bb2b44740d87795fe8"
	size := 20000000
	dir :=  "store.logfile"

	fmt.Println(str, size)

	keys, value := get_address(size)

	root := common.BytesToHash(common.FromHex(str))
	fmt.Println("root: ", root)

	lmpt_Get_TrieDB(root, dir, size, keys, value, 512, 2000)
}




//  lmpt 增大 一次 commit 数量  1000000  ，统计时间， 采用真实数据集
func time_count_lmpt_disk (num int) {
	size := num
	//  key 为 真实 交易数据
	keys, value := get_address(size)
	fmt.Println(len(value))

	fmt.Println("---------------------------------------------------")
	fmt.Println("address and value get over")


	start := time.Now() // 获取当前时间
	times := make([]int64, num / 1000000)

	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))

	dir := "store.logfile"
	dbase, _ := leveldb.New(dir,8192,500,"cc",false)
	triedb := trie.NewDatabase(dbase)
	tree := trie.NewEmpty(triedb)

	for i := 0; i < size; i++ {
		v := tree.Get(keys[i])
		if v == nil {
			tree.Update(keys[i], value)
		} else {
			v = ValueUpdate(v)
			tree.Update(keys[i], v)
		}

		if i % 100000 == 9999 {
			fmt.Println(i)

			if i % 1000000 == 9999 {
				root, nodes, _ := tree.Commit(false)
				triedb.Update(trie.NewWithNodeSet(nodes))
				triedb.Commit(root, true, nil)
				tree, _ = trie.New(trie.TrieID(root), triedb)

				elapsed1 := time.Since(start)
				//fmt.Println("该函数执行完成耗时：", elapsed1)
				times[i / 1000000] = elapsed1.Milliseconds()

			}

		}
	}
	fmt.Println("tree experiment over")
	elapsed1 := time.Since(start)
	fmt.Println("执行完成耗时：", elapsed1.Milliseconds())

	fmt.Println(times)

}
