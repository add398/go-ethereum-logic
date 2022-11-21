/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/11/21 16:44
 */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.New(rand.NewSource(0))
	size := 100

	keys := make([][]byte, size)
	values := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	for i := 0; i < size; i++ {
		v := make([]byte, 30)
		random.Read(v)
		values[i] = v
	}

	fmt.Println(keys)
	fmt.Println()
	fmt.Println(values)

}
