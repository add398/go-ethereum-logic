/**
 * @Author: Log1c
 * @Description:
 * @File:  count_num_txn.go
 * @Version: 1.0.0
 * @Date: 2022/11/21 17:05
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_count_num (t *testing.T) {
	//打开文件
	file, err := os.Open("D:\\Code\\go-ethereum-logic\\tests_Log1c\\Test_dataset\\dataset\\address.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)


	count := 0
	keys := map[string]int{}

	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束

		keys[str] = 1
		count++
		if count == 40000000{
			break
		}

		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
	}

	fmt.Println(count)
	fmt.Println(len(keys))
	fmt.Println("文件读取结束...")
}
