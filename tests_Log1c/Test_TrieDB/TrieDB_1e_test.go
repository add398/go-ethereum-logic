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


// 存储 1e 个独立地址
// 0x997dc71cea4cce1fe9044f65a257babb6372481d85a49f5d348591dd3159d973
// 0x997dc71cea4cce1fe9044f65a257babb6372481d85a49f5d348591dd3159d973
//--- PASS: Test_store_1e_2 (21804.56s)
//PASS
func Test_store_1e_2(t *testing.T) {
	size := 100000000
	dir := "store1e.logfile"

	store_disk_triedb(size, dir)

	// 0x997dc71cea4cce1fe9044f65a257babb6372481d85a49f5d348591dd3159d973
	// 0x997dc71cea4cce1fe9044f65a257babb6372481d85a49f5d348591dd3159d973
	//--- PASS: Test_store_1e_2 (21804.56s)
	//PASS
}



// benchmark  读取地址
//10000
//Benchmark_read_triedb_1e_1w_2-12    	       1	27456606700 ns/op
//PASS

//Benchmark_read_triedb_1e_1w_2
//Benchmark_read_triedb_1e_1w_2-12    	       1	17040827800 ns/op
//PASS
//Benchmark_read_triedb_1e_1w_2
//Benchmark_read_triedb_1e_1w_2-12    	       1	15726043800 ns/op
//PASS

//Benchmark_read_triedb_1e_1w_2   32mb
//Benchmark_read_triedb_1e_1w_2-12    	       1	1276289584100 ns/op
func Benchmark_read_triedb_1e_1w_2(b *testing.B) {
	size := 100000000
	dir := "store1e.logfile"
	str := "0x997dc71cea4cce1fe9044f65a257babb6372481d85a49f5d348591dd3159d973"
	benchmark_read_triedb(b, size, dir, str)

}

//  time 函数 测试读取地址时间
// 执行完成耗时： 7534940500
//--- PASS: Test_read_triedb_1e_1w (13.24s)
//PASS
// tree experiment over
//执行完成耗时： 11676484000

func Test_read_triedb_1e_1w(t *testing.T) {
	size := 100000000
	dir := "store1e.logfile"
	str := "0x997dc71cea4cce1fe9044f65a257babb6372481d85a49f5d348591dd3159d973"
	read_triedb_time(size, dir, str)

}


func Test_store_1e(t *testing.T) {
	size := 1000000
	dir := "store1e.logfile"

	test_Store_TrieDB(t, size, dir)
}




func Benchmark_read_FROM_leveldb_1e(b *testing.B) {
	size := 100000000
	dir := "store1e.logfile"
	str := ""
	benchmark_Get_TrieDB(b, size, dir, str)
}