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
	file, err := os.Open("dataset/from_to.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)


	writeFilePath := "dataset/address.txt"
	writeFile, err := os.OpenFile(writeFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer writeFile.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(writeFile)

	//for i := 0; i < 5; i++ {
	//	write.WriteString("http://c.biancheng.net/golang/ \n")
	//}


	count := 0

	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束

		if count == 0 {

		}else if count > 0 && count < 5000001{
			from := str[:43]
			to := str[43:85]
			fmt.Println(str)
			fmt.Println(from)
			fmt.Println(to)

			write.WriteString(from)
			write.WriteString("\n")
			write.WriteString(to)
			write.WriteString("\n")

		}else if count == 5000001 {
			break
		}

		count++

		write.Flush()


		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
	}
	//Flush将缓存的文件真正写入到文件中

	fmt.Println(count)
	fmt.Println("文件读取结束...")
}
