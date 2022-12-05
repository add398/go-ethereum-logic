/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/12/4 05:37
 */

package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"

	"io"
	"os"
	"time"
)

func get_address( size int) (keys [][]byte) {
	// 一次获取所有的 address
	// tests_Log1c/Test_dataset/dataset/1100wto1200w_BlockTransaction_Address.csv
	file, err := os.Open("tests_Log1c/dataset/1100wto1200w_BlockTransaction_Address.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	keys = make([][]byte, size)
	for i := 0; ; i++ {
		if i == size {
			break
		}
		csvdata, err := reader.Read() // 按行读取数据,可控制读取部分
		if err == io.EOF {
			fmt.Println("总行数 ", i)
			break
		}
		address := csvdata[0]
		addbyte := StringTobyte(address)

		keys[i] = addbyte
	}

	return
}

func Space_count(num int) {
	// 计算 空间的实验  有 重复地址
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

	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	triedb := trie.NewDatabase(memorydb.New())
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
		}
	}

	fmt.Println("tree experiment over")
	time.Sleep(time.Minute)
}


func time_count_memory (num int) {
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
	triedb := trie.NewDatabase(memorydb.New())
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
			go func() {
				fmt.Println(i)
			}()
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
	dbase, _ := leveldb.New(dir,10240,500,"cc",false)
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



func count_address( size int) (keys [][]byte) {
	// 一次获取所有的 address
	// tests_Log1c/Test_dataset/dataset/1100wto1200w_BlockTransaction_Address.csv
	file, err := os.Open("tests_Log1c/dataset/1100wto1200w_BlockTransaction_Address.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	hashmap := map[string]struct{}{}

	keys = make([][]byte, size)
	for i := 0; ; i++ {
		if i == size {
			break
		}
		csvdata, err := reader.Read() // 按行读取数据,可控制读取部分
		if err == io.EOF {
			fmt.Println("总行数 ", i)
			break
		}
		address := csvdata[0]
		addbyte := StringTobyte(address)

		hashmap[address] = struct{}{}

		keys[i] = addbyte
	}

	fmt.Println(len(hashmap))
	return
}

func main() {
	//num := 50000000
	//time_count_disk(num)

	count_address(50000000)
}
