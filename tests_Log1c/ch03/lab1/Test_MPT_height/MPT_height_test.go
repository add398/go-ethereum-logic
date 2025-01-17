/**
 * @Author: Log1c
 * @Description:
 * @File:  MPT_height_test
 * @Version: 1.0.0
 * @Date: 2022/11/23 16:12
 */

package Test_MPT_height

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"math/rand"
	"testing"
)

func Test_MPT_height(t *testing.T) {

	trie := NewEmpty(NewDatabase(rawdb.NewMemoryDatabase()))
	ks := [][]byte{}

	key1 := []byte("abcde")
	key2 := []byte("abcdf")
	key3 := []byte("abcef")
	key4 := []byte("abddd")
	key5 := []byte("aaeee")
	key6 := []byte("baeee")



	ks = append(ks,key1)
	ks = append(ks,key2)
	ks = append(ks,key3)
	ks = append(ks,key4)
	ks = append(ks,key5)
	ks = append(ks,key6)

	size := len(ks)
	for i := 0; i < size; i++ {
		trie.Update(ks[i], []byte("value"))
	}

	for i := 0; i < size; i++ {
		fmt.Println("key", keybytesToHex(ks[i]))

		_, height := trie.Get(ks[i])
		fmt.Println( "height",height)
	}



}



func Test_tiny_1_mpt(t *testing.T) {

	trie := NewEmpty(NewDatabase(rawdb.NewMemoryDatabase()))
	ks := [][]byte{}

	key1 := []byte{0x01}
	ks = append(ks, key1)

	size := len(ks)
	for i := 0; i < size; i++ {
		trie.Update(ks[i], []byte("value"))
	}

	for i := 0; i < size; i++ {
		fmt.Println("key", keybytesToHex(ks[i]))
		_, height := trie.Get(ks[i])
		fmt.Println( "height",height)
	}



}


func Test_tiny_2_mpt(t *testing.T) {

	trie := NewEmpty(NewDatabase(rawdb.NewMemoryDatabase()))
	ks := [][]byte{}

	key1 := []byte{0x01,0x02}
	key2 := []byte{0x01,0x01}
	ks = append(ks, key1)
	ks = append(ks, key2)


	size := len(ks)
	for i := 0; i < size; i++ {
		trie.Update(ks[i], []byte("value"))

	}

	for i := 0; i < size; i++ {
		fmt.Println("key", keybytesToHex(ks[i]))
		_, height := trie.Get(ks[i])
		fmt.Println( "height",height - 1)
	}



}



func Test_MPT_height_10w(t *testing.T) {
	size := 10000000
	trie := NewEmpty(NewDatabase(rawdb.NewMemoryDatabase()))
	random := rand.New(rand.NewSource(0))

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 3)
		random.Read(k)
		keys[i] = k
	}


	for i := 0; i < size; i++ {
		trie.Update(keys[i], []byte("value"))
	}

	for i := 0; i <  100; i++ {
		fmt.Println("key", keybytesToHex(keys[i]))
		_, height := trie.Get(keys[i])
		fmt.Println( "height",height)
	}



}
