/**
 * @Author: Log1c
 * @Description:
 * @File:  lab2
 * @Version: 1.0.0
 * @Date: 2023/2/1 13:23
 */

package main

import (
	"math/rand"
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


