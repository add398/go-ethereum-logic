/**
 * @Author: Log1c
 * @Description:
 * @File:  LevelDB_100w_test
 * @Version: 1.0.0
 * @Date: 2022/10/24 11:55
 */

package _0_15

import "testing"

func Test_store_1000w(t *testing.T) {
	size := 10000000
	dir := "store1000w.logfile"

	test_store_Size(t, size, dir)
}




func Benchmark_read_FROM_leveldb_1000w(b *testing.B) {
	size := 10000000
	dir := "store1000w.logfile"
	benchmark_read_FROM_leveldb(b, size, dir)
}