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

func Test_store_100w(t *testing.T) {
	size := 1000000
	dir := "store100w.logfile"

	test_Store_TrieDB(t, size, dir)
}




func Benchmark_read_FROM_leveldb_100w(b *testing.B) {
	size := 1000000
	dir := "store100w.logfile"
	str := "0xe75427e9da29b7eeabc44c2e1714e00fc02f9040235cda92f609447eefd420a3"
	benchmark_Get_TrieDB(b, size, dir, str)
}