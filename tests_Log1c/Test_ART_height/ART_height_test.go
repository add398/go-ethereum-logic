/**
 * @Author: Log1c
 * @Description:
 * @File:  MPT_height_test
 * @Version: 1.0.0
 * @Date: 2022/11/23 16:12
 */

package Test_ART_height

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"math/rand"
	"testing"
)

func PrintHeight(nums []int)  {
	fmt.Print("数组nums：", nums, "    分支 和 扩展 : ",  nums[0],   nums[1] - 1 , "     ")
	fmt.Println("MPT height：  " ,   nums[0] + nums[1] - 1 ,  "   ART height：  " ,   nums[0] / 2 + nums[1] - 1)
}


func Test_tiny_1_art(t *testing.T) {

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
		PrintHeight(height)
	}



}


func Test_tiny_2_art(t *testing.T) {

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
		PrintHeight(height)
	}



}


func Test_tiny_3_art(t *testing.T) {

	trie := NewEmpty(NewDatabase(rawdb.NewMemoryDatabase()))
	ks := [][]byte{}

	key1 := []byte{0x12,0x02}
	key2 := []byte{0x21,0x01}
	ks = append(ks, key1)
	ks = append(ks, key2)


	size := len(ks)
	for i := 0; i < size; i++ {
		trie.Update(ks[i], []byte("value"))

	}

	for i := 0; i < size; i++ {
		fmt.Println("key", keybytesToHex(ks[i]))
		_, height := trie.Get(ks[i])
		PrintHeight(height)
	}
}


func Test_ART_height_10w(t *testing.T) {
	size := 50000000
	trie := NewEmpty(NewDatabase(rawdb.NewMemoryDatabase()))
	random := rand.New(rand.NewSource(0))

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}


	for i := 0; i < size; i++ {
		trie.Update(keys[i], []byte("0"))
		if i % 1000000 == 0 {
			fmt.Println(i)

		}
	}


	var h1, h2 float64


	for i := 0; i < size; i++ {
		count := size / 100
		if i % count == 0 {
			//fmt.Println("key", keybytesToHex(keys[i]))
			_, height := trie.Get(keys[i])
			//PrintHeight(height)
			h1 += float64(height[0] + height[1] - 1)
			h2 += float64(height[0]) / 2.0   + float64(height[1] - 1)
		}
	}

	fmt.Println("MPT: ", h1 / 100)
	fmt.Println("ART: ", h2 / 100)
	
}






func Test_ART_height_10w(t *testing.T) {
	size := 50000000
	trie := NewEmpty(NewDatabase(rawdb.NewMemoryDatabase()))
	random := rand.New(rand.NewSource(0))

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}


	for i := 0; i < size; i++ {
		trie.Update(keys[i], []byte("0"))
		if i % 1000000 == 0 {
			fmt.Println(i)

		}
	}


	var h1, h2 float64


	for i := 0; i < size; i++ {
		count := size / 100
		if i % count == 0 {
			//fmt.Println("key", keybytesToHex(keys[i]))
			_, height := trie.Get(keys[i])
			//PrintHeight(height)
			h1 += float64(height[0] + height[1] - 1)
			h2 += float64(height[0]) / 2.0   + float64(height[1] - 1)
		}
	}

	fmt.Println("MPT: ", h1 / 100)
	fmt.Println("ART: ", h2 / 100)














}
