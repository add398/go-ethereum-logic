/**
 * @Author: Log1c
 * @Description:
 * @File:  log1c_look_test
 * @Version: 1.0.0
 * @Date: 2023/1/18 15:22
 */

package main

import (
	"math/rand"
	"testing"
	"time"
)

func makeAccounts(size int) (addresses [][]byte, value []byte) {
	// Make the random benchmark deterministic
	random := rand.New(rand.NewSource(0))
	addresses = make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 32)
		random.Read(k)
		addresses[i] = k
	}

	value = make([]byte, 100)
	random.Read(value)
	return
}

func Test_Space(t *testing.T) {
	size := 500
	keys, value := makeAccounts(size)
	tree := NewTrie()
	for j := 0; j < size; j++ {
		tree.Put(keys[j], value)
	}
	time.Sleep(1 * time.Hour)



}


func BenchmarkLook(b *testing.B) {
	size := 1000000
	keys, value := makeAccounts(size)
	tree := NewTrie()
	for j := 0; j < size; j++ {
		tree.Put(keys[j], value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := size / 10000
		for j := 0; j < size; j++ {
			if j % count == 99 {
				tree.Get(keys[j])
			}
		}
	}
	b.StopTimer()

}



func BenchmarkUpdate(b *testing.B) {
	size := 1000000
	keys, value := makeAccounts(size)
	tree := NewTrie()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			tree.Put(keys[j], value)
		}
	}
	b.StopTimer()

}




