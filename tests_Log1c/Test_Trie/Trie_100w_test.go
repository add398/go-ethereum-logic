/**
 * @Author: Log1c
 * @Description:
 * @File:  Trie_100w
 * @Version: 1.0.0
 * @Date: 2022/10/24 18:03
 */

package Test_Trie

import "testing"

func Benchmark_read_from_trie_100w(b *testing.B) {
	size := 1000000
	benchmark_read_from_trie(b, size)
}

func Benchmark_read_from_trie_1000w(b *testing.B) {
	size := 10000000
	benchmark_read_from_trie(b, size)
}

func Benchmark_read_from_trie_10000w(b *testing.B) {
	size := 100000000
	benchmark_read_from_trie(b, size)
}