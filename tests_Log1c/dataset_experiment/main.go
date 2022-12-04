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
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/trie"

	"io"
	"os"
	"time"
)


func get_address(name string, size int) (keys [][]byte,  value []byte)  {
	// 总行数  10374345
	// 1534136  独立地址
	file, err := os.Open("tests_Log1c/dataset_experiment/dataset/" + name + ".csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	keys = make([][]byte, size)
	for i := 0;   ; i++{
		if i == size {
			break
		}
		csvdata, err := reader.Read() // 按行读取数据,可控制读取部分
		if err == io.EOF {
			fmt.Println("总行数 ", i)
			break
		}
		address := csvdata[0]
		addbyte := stringTobyte(address)
		keys[i] = addbyte
	}

	value = make([]byte, 100)
	return
}




func stringTobyte(str  string ) []byte {
	// 地址字符串转换为  20 byte
	n := len(str)
	ans := make([]byte, 20)
	c := 0
	for i := 2; i < n; i = i + 2 {
		a, b := byteToint(str[i]), byteToint(str[i+1])
		ans[c] = byte(a * 15 + b)
		c++
	}
	return ans
}

func byteToint(a byte) int {
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

func valueUpdate(value []byte) []byte {
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



func main() {
	name := "1300wto1325w_BTXN"
	size := 10000000
	keys, value := get_address(name, size)

	fmt.Println("address")
	//tree := trie.NewEmpty(trie.NewDatabase(memorydb.New()))
	tree := trie.NewEmpty(trie.NewDatabase(rawdb.NewMemoryDatabase()))

	for i := 0; i < size; i++ {
		//if i % 1000000 == 0 {
		//	fmt.Println( i )
		//}
		tree.Update(keys[i], value)
	}

	fmt.Println("over")
	time.Sleep(time.Minute)
}




