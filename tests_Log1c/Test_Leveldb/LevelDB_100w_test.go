/**
 * @Author: Log1c
 * @Description:
 * @File:  LevelDB_100w_test
 * @Version: 1.0.0
 * @Date: 2022/10/24 11:55
 */

package _0_15

import (
	"testing"
)

func Test_store_100w(t *testing.T) {
	size := 2000000
	dir := "store100w.logfile"

	test_store_Size(t, size, dir)
}




func Benchmark_read_FROM_leveldb_100w(b *testing.B) {
	size := 2000000
	dir := "store100w.logfile"
	benchmark_read_FROM_leveldb(b, size, dir)
}
