/**
 * @Author: Log1c
 * @Description:
 * @File:  leaf_test
 * @Version: 1.0.0
 * @Date: 2022/12/1 17:16
 */

package ART_LRU

import (
	"math/rand"
	"testing"
)

func Test_Tree_ForEachByLeaf(t *testing.T) {
	tree := New()
	terms := []string{"1", "2", "3", "4", "5"}

	for _, term := range terms {
		tree.Insert(Key(term), term)
	}
	tree.ForEachByLeaf()
	//fmt.Println(tree)
	tree.Insert(Key("1"), "1")
	tree.ForEachByLeaf()

	tree.Delete(Key("1"))

	tree.ForEachByLeaf()
	//fmt.Println(tree)
}


func Benchmark_Tree_search(b *testing.B) {
	size := 100000

	benchmark_read_from_trie(b, size)

}


func benchmark_read_from_trie(b *testing.B, size int) {
	random := rand.New(rand.NewSource(0))
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	value := make([]byte, 100)
	random.Read(value)
	//fmt.Println(value)

	tree := New()
	for i := 0; i < size; i++ {
		tree.Insert(keys[i], value)
	}

	count := size / 10000

	b.ResetTimer()
	b.ReportAllocs()
	for j := 0; j < b.N; j++ {
		for i := 0; i < size; i++ {
			if i % count == 0 {
				tree.Search(keys[i])
			}

		}
	}
	b.StopTimer()
}


