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
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"

	"io"
	"os"
	"time"
)

func get_address(name string, size int) (keys [][]byte, value []byte) {
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

	value = make([]byte, 100)
	return
}

func main() {
	name := "1300wto1325w_BTXN"
	size := 5000000
	keys, value := get_address(name, size)
	fmt.Println(value)
	fmt.Println("address over")

	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	triedb := trie.NewDatabase(memorydb.New())
	tree := trie.NewEmpty(triedb)


	for i := 0; i < size; i++ {
		//if i % 1000000 == 0 {
		//	fmt.Println( i )
		//}

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

	fmt.Println("main over")
	time.Sleep(time.Minute)
}
