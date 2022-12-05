/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/12/5 15:22
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // 获取当前时间
	time.Sleep(time.Second)
	elapsed1 := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed1)

}
