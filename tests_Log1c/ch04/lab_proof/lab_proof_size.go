/**
 * @Author: Log1c
 * @Description:
 * @File:  lab2
 * @Version: 1.0.0
 * @Date: 2023/2/1 13:24
 */

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/trie"
	"math/rand"
)

func address_and_value(size int) ([][]byte,  []byte){
	random := rand.New(rand.NewSource(0))

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}
	value := make([]byte, 100)
	random.Read(value)
	return keys, value
}

func Proof_size() {
	mpt := NewTrie()
	size := 10
	keys, val := address_and_value(size)
	for i := 0; i < size; i++ {
		mpt.Put(keys[i], val)
	}
	proof, _ := mpt.Prove(keys[0])

	proDB := proof.(*ProofDB)
	proDB.PrintProofDB()

}



func EthProof_size() {
	mpt := new(trie.Trie)
	size := 10
	keys, val := address_and_value(size)
	for i := 0; i < size; i++ {
		mpt.Update(keys[i], val)
	}
	w := NewProofDB()
	_ = mpt.Prove(keys[0], 0, w)
	//fmt.Println(w.kv)
	fmt.Println(len(w.kv))
	sum := 0
	for _, v := range w.kv {
		//fmt.Println(k)
		fmt.Println(len(v))
		curLen := len(v) / 32
		fmt.Println(curLen - 1)
		sum += curLen - 1
	}
	fmt.Println(sum)
}

func main() {
	Proof_size()
}