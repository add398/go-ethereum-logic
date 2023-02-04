package main

import (
	"math/rand"
)


func makeAccounts1(size int) (addresses [][]byte, value []byte) {
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

func MPT_Size(size int, keys [][]byte, value []byte) {
	// MPT 空间


	tree := NewTrie()
	for i := 0; i < size; i++ {
		tree.Put(keys[i], value)
	}

}




//func main() {
//	size := 10000000
//	size1 := 10000000
//	keys, value := makeAccounts1(size)
//	MPT_Size(size1, keys, value)
//
//	fmt.Println("over ")
//	time.Sleep(1 * time.Hour)
//
//
//
//
//}
