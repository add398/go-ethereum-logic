/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/1/18 18:57
 */

package main


import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)


func makeAccounts(size int) (addresses [][]byte, value []byte) {
	// Make the random benchmark deterministic
	random := rand.New(rand.NewSource(0))
	addresses = make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		addresses[i] = k
	}

	value = make([]byte, 100)
	random.Read(value)
	return
}





func main() {
	size := 100
	//keys, value := makeAccounts(size)
	url := "http://192.168.18.101:8080/ping"
	url2 := "http://192.168.18.101:8080/pong"


	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	start := time.Now() // 获取当前时间

	for i := 0; i < size; i++ {
		resp, _ = http.Get(url2)

	}

	elapsed1 := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed1.Seconds())

	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}



}
