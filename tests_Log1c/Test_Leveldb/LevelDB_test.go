/**
 * @Author: Log1c
 * @Description:
 * @File:  a_test
 * @Version: 1.0.0
 * @Date: 2022/10/15 22:17
 */

package _0_15

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"log"
	"math/rand"
	"testing"
)

func Test_store_one(t *testing.T) {
	dbase, err := leveldb.New("store.logfile",8,500,"cc",false)
	if err != nil {
		log.Println("database open fail")
	}
	dbase.Put([]byte("123"), []byte("abcde"))

	//defer dbase.Close()
}

func Test_read_one(t *testing.T) {
	dbase, err := leveldb.New("store.logfile",8,500,"cc",false)

	if err != nil {
		log.Println("database open fail")
	}
	ans, _ := dbase.Get([]byte("123"))
	fmt.Println(string(ans))

	defer dbase.Close()
}

func Test_store_100000000(t *testing.T) {
	dbase, err := leveldb.New("store.logfile",8,500,"cc",false)
	if err != nil {
		log.Println("database open fail")
	}
	random := rand.New(rand.NewSource(0))
	size := 100000000
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	for i := 0; i < size; i++ {
		dbase.Put(keys[i], []byte("qwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnmqwertyuioplkjhgfdsazxcvbnm"))
	}


	defer dbase.Close()
}



func Test_read_SOME(t *testing.T) {
	dbase, err := leveldb.New("store.logfile",8,500,"cc",false)
	if err != nil {
		log.Println("database open fail")
	}
	random := rand.New(rand.NewSource(0))
	size := 100
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}

	for i := 0; i < size; i++ {
		v, _ := dbase.Get(keys[i])
		fmt.Println(v)

	}


	defer dbase.Close()
}




func Benchmark_read_FROM_leveldb(b *testing.B) {
	dbase, err := leveldb.New("store.logfile",8,500,"cc",false)
	if err != nil {
		log.Println("database open fail")
	}
	defer dbase.Close()

	random := rand.New(rand.NewSource(0))

	size := 100000000
	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}


	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		for j := 0; j < size; j++ {
			if j % 10000 == 0 {
				_, _ = dbase.Get(keys[j])
			}
			//fmt.Println(v)
		}
	}

}

