/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/11/21 17:28
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("D:\\Code\\go-ethereum-logic\\tests_Log1c\\Test_dataset\\dataset\\from_to.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)


	writeFilePath := "D:\\Code\\go-ethereum-logic\\tests_Log1c\\Test_dataset\\dataset\\address.txt"
	writeFile, err := os.OpenFile(writeFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer writeFile.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(writeFile)


	count := 0

	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}

		if count == 0 {

		}else {
			from := str[:43]
			to := str[43:]
			//fmt.Println(str)
			//fmt.Println(from)
			//fmt.Println(to)

			write.WriteString(from)
			write.WriteString("\n")
			write.WriteString(to)

		}

		count++




	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()


	fmt.Println(count)
	fmt.Println("文件读取结束...")
}
