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


// 存储 100w 个独立地址
func Test_store_100w_2(t *testing.T) {
	size := 1000000
	dir := "store100w.logfile"

	store_disk_triedb(size, dir)
}

// benchmark  读取地址
func Benchmark_read_triedb_100w_1w_2(b *testing.B) {
	size := 1000000
	dir := "store100w.logfile"
	str := "0x3b4d46cb7d4b91496ba54b88216f67d74673c955df28ad50d1bbc28d697ef742"
	benchmark_read_triedb(b, size, dir, str)

}

//  time 函数 测试读取地址时间
func Test_read_triedb_100w_1w(t *testing.T) {
	size := 1000000
	dir := "store100w.logfile"
	str := "0x3b4d46cb7d4b91496ba54b88216f67d74673c955df28ad50d1bbc28d697ef742"
	read_triedb_time(size, dir, str)
	// Benchmark_read_triedb_1e_1w_2-12    	      57	  18613919 ns/op
	// 执行完成耗时： 556271600
}







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