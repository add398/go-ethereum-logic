/**
 * @Author: Log1c
 * @Description:
 * @File:  Trie_3_test
 * @Version: 1.0.0
 * @Date: 2022/12/4 00:17
 */

package Test_Trie

import (
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"io"
	"os"
	"testing"
)



func Benchmark_read_from_memorydb_5kw(b *testing.B) {
	// 2500w 交易， 5kw个地址， 从该树中 执行 1w笔交易所需时间
	// 500w 独立地址
	b.StopTimer()

	size := 50000000
	keys := get_address(size)
	value := make([]byte, 0)
	for i := 0; i < 5; i++ {
		value = append(value, keys[0]...)
	}
	fmt.Println(keys[0])
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

	for i := 0; i < b.N; i++ {
		read_from_memorydb_5kw(b, tree, keys, value)
	}


	fmt.Println("tree experiment over")
}


func get_address( size int) (keys [][]byte) {
	// 一次获取所有的 address
	// tests_Log1c/Test_dataset/dataset/1100wto1200w_BlockTransaction_Address.csv
	file, err := os.Open("D:\\Code\\go-ethereum-logic\\tests_Log1c\\dataset\\1100wto1200w_BlockTransaction_Address.csv")
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

func StringTobyte(str  string ) []byte {
	// 地址字符串转换为  20 byte
	n := len(str)
	ans := make([]byte, 20)
	c := 0
	for i := 2; i < n; i = i + 2 {
		a, b := ByteToint(str[i]), ByteToint(str[i+1])
		ans[c] = byte(a * 15 + b)
		c++
	}
	return ans
}

func ByteToint(a byte) int {
	switch a {
	case '0':return 0
	case '1':return 1
	case '2':return 2
	case '3':return 3
	case '4':return 4
	case '5':return 5
	case '6':return 6
	case '7':return 7
	case '8':return 8
	case '9':return 9
	case 'a':return 10
	case 'b':return 11
	case 'c':return 12
	case 'd':return 13
	case 'e':return 14
	case 'f':return 15
	}
	return -1
}



func ValueUpdate(value []byte) []byte {
	// 改变 value
	n := len(value)
	for i := 0; i < n; i++ {
		if value[i] < byte(255) {
			value[i] = value[i] + byte(1)
			break
		}else{
			// == 255
			value[i] = byte(0)
		}
	}
	return value
}



func read_from_memorydb_5kw(b *testing.B, tree *trie.Trie,  keys [][]byte, value []byte) {
	size := len(keys)
	count := size / 10000
	b.StartTimer()
	for i := 0; i < size; i++ {
		if i % count == 1 {
			v1 := tree.Get(keys[i])
			v2 := tree.Get(keys[i+1])
			v1 = ValueUpdate(v1)
			v2 = ValueUpdate(v2)

			tree.Update(keys[i],  v1)
			tree.Update(keys[i+1], v2)
		}
	}
	b.StopTimer()
}



func Benchmark_3_10000w(b *testing.B) {
	size := 10000000
	keys, val := makeAccounts(size)
	for i := 0; i < b.N; i++ {
		benchmark_3_trie(b, keys, val)
	}

	fmt.Println("over")
}



func benchmark_3_trie(b *testing.B, keys [][]byte, value []byte) {
	b.ReportAllocs()

	size := len(keys)
	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	tree := trie.NewEmpty(trie.NewDatabase(rawdb.NewMemoryDatabase()))

	b.StartTimer()
	for i := 0; i < size; i++ {
		if i % 1000000 == 0 {
			fmt.Println( i )
		}
		tree.Update(keys[i], value)
	}
	b.StopTimer()
}


