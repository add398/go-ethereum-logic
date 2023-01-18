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
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
	"time"
)




//  MPT 存储在 disk 上，统计花费的世界， 采用真实数据集
func time_count_disk (num int) {
	size := num
	//  key 为 真实 交易数据
	keys := get_address(size)

	// value 为 5个 20byte 地址合起来
	value := make([]byte, 0)
	for i := 0; i < 5; i++ {
		value = append(value, keys[0]...)
	}
	fmt.Println(len(value))

	fmt.Println("---------------------------------------------------")
	fmt.Println("address and value get over")


	start := time.Now() // 获取当前时间
	times := make([]int64, num / 20000)

	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))

	dir := "store.logfile"
	dbase, _ := leveldb.New(dir,2048,500,"cc",false)
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

		if i % 20000 == 19999 {
			root, nodes, _ := tree.Commit(false)
			triedb.Update(trie.NewWithNodeSet(nodes))
			triedb.Commit(root, true, nil)
			tree, _ = trie.New(trie.TrieID(root), triedb)
			fmt.Println(i)

			elapsed1 := time.Since(start)
			//fmt.Println("该函数执行完成耗时：", elapsed1)
			times[i / 20000] = elapsed1.Milliseconds()
		}
	}
	fmt.Println("tree experiment over")
	elapsed1 := time.Since(start)
	fmt.Println("执行完成耗时：", elapsed1.Milliseconds())

	fmt.Println(times)
	time.Sleep(time.Minute)
}

func main() {
	num := 20000000
	time_count_disk(num)

	//count_address(50000000)
}
