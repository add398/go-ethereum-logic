/**
 * @Author: Log1c
 * @Description:
 * @File:  LevelDB_100w_test
 * @Version: 1.0.0
 * @Date: 2022/10/24 11:55
 */

package _0_15

import (
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"log"
	"math/rand"
	"testing"
)

func Test_store_1e(t *testing.T) {
	size := 100000000
	dir := "store1e.logfile"

	test_store_Size(t, size, dir)
}




func Benchmark_read_FROM_leveldb_1e(b *testing.B) {
	size := 100000000
	dir := "store1e.logfile"
	benchmark_read_FROM_leveldb(b, size, dir)
}


func Benchmark_store_leveldb_1ww(b *testing.B) {
	size := 10000
	dir := "store1e.logfile"

	dbase, err := leveldb.New(dir,8,500,"cc",false)
	if err != nil {
		log.Println("database open fail")
	}
	defer dbase.Close()

	random := rand.New(rand.NewSource(1))
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}


	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < size; i++ {
			dbase.Put(keys[i], []byte("qwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnm"))
		}
	}
	b.StopTimer()

}