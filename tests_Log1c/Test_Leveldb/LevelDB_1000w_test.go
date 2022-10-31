/**
 * @Author: Log1c
 * @Description:
 * @File:  LevelDB_100w_test
 * @Version: 1.0.0
 * @Date: 2022/10/24 11:55
 */

package _0_15

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"log"
	"math/rand"
	"testing"
)

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


func Benchmark_store_leveldb_1000w(b *testing.B) {
	size := 10000
	dir := "store1000w.logfile"

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
	fmt.Println(1)

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < size; i++ {
			dbase.Put(keys[i], []byte("qwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnm"))
		}
	}
	b.StopTimer()

}
