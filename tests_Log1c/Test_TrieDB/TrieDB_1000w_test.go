/**
 * @Author: Log1c
 * @Description:
 * @File:  LevelDB_100w_test
 * @Version: 1.0.0
 * @Date: 2022/10/24 11:55
 */

package Test_TrieDB

import "testing"

func Test_store_1000w(t *testing.T) {
	size := 10000000
	dir := "store1000w.logfile"

	test_Store_TrieDB(t, size, dir)
}




func Benchmark_read_FROM_leveldb_1000w(b *testing.B) {
	size := 10000000
	dir := "store1000w.logfile"
	str := "0xa0219c0aa81ebffbb7c5ee4b9981e35827c6d70aff298236cbcf692df34af8a0"
	benchmark_Get_TrieDB(b, size, dir, str)
}