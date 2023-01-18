/**
 * @Author: Log1c
 * @Description:
 * @File:  ART_in_disk_test
 * @Version: 1.0.0
 * @Date: 2023/1/18 15:50
 */

package dataset

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
	"time"
)

func time_count_disk (num int) {
	// 计算 时间 的实验
	//name := "1300wto1325w_BTXN"
	size := num
	keys := get_address(size)
	value := make([]byte, 0)
	for i := 0; i < 5; i++ {
		value = append(value, keys[0]...)
	}
	fmt.Println(len(value))
	fmt.Println("---------------------------------------------------")
	fmt.Println("address and value get over")


	start := time.Now() // 获取当前时间
	times := make([]int64, num / 5000)

	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))

	dir := "store.logfile"
	dbase, _ := leveldb.New(dir,5120,500,"cc",false)
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

		if i % 5000 == 4999 {
			root, nodes, _ := tree.Commit(false)
			triedb.Update(trie.NewWithNodeSet(nodes))
			triedb.Commit(root, true, nil)
			tree, _ = trie.New(trie.TrieID(root), triedb)
			fmt.Println(i)

			elapsed1 := time.Since(start)
			//fmt.Println("该函数执行完成耗时：", elapsed1)
			times[i / 5000] = elapsed1.Milliseconds()
		}
	}
	fmt.Println("tree experiment over")
	elapsed1 := time.Since(start)
	fmt.Println("执行完成耗时：", elapsed1.Milliseconds())

	fmt.Println(times)
	time.Sleep(time.Minute)
}

