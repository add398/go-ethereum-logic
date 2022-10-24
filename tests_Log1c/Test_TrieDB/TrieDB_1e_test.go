/**
 * @Author: Log1c
 * @Description:
 * @File:  LevelDB_100w_test
 * @Version: 1.0.0
 * @Date: 2022/10/24 11:55
 */

package Test_TrieDB

import (
	"testing"
)

func Test_store_1e(t *testing.T) {
	size := 100000000
	dir := "store1e.logfile"

	test_Store_TrieDB(t, size, dir)
}




func Benchmark_read_FROM_leveldb_1e(b *testing.B) {
	size := 100000000
	dir := "store1e.logfile"
	str := ""
	benchmark_Get_TrieDB(b, size, dir, str)
}