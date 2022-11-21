/**
 * @Author: Log1c
 * @Description:
 * @File:  Update_dateset_test
 * @Version: 1.0.0
 * @Date: 2022/11/21 18:57
 */

package Test_update

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
	"math/rand"
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {
	// 46684549 条交易地址
	FilePath := "D:\\Code\\go-ethereum-logic\\tests_Log1c\\Test_dataset\\dataset\\address.txt"
	File, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer File.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(File)

	dir := "update.logfile"


	//  keys
	size := 10000000
	keys := make([]string, size)
	for i := 0; i < size; i++ {
		str, _ := reader.ReadString('\n') //读到一个换行就结束
		keys[i] = str
	}
	//fmt.Println(keys)
	// values
	random := rand.New(rand.NewSource(0))
	values := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		values[i] = k
	}
	//fmt.Println(values)


	dbase, err := leveldb.New(dir,8,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	triedb := trie.NewDatabase(dbase)
	tree := trie.NewEmpty(triedb)

	for i :=0; i < size; i++ {
		tree.Update([]byte(keys[i]), values[i])
	}
	root, nodes, _ := tree.Commit(false)
	triedb.Update(trie.NewWithNodeSet(nodes))
	triedb.Commit(root, true, nil)

	fmt.Println(root)


}


// 0xf364913f32e1e1ed31c5d36ff17828a4a31caad8fa052f7c4e1fc365657c788e
/*
test performance from a update trieDB
*/
func Benchmark_read_FROM_TrieDB_update_100w(b *testing.B) {
	size := 10000
	dir := "update.logfile"
	str := "0xc51bd935560a00c60ca3c582aebd5803af7d72396fcb69919bd1b08935675641"
	benchmark_Get_TrieDB(b, size, dir, str)
}


func benchmark_Get_TrieDB(b *testing.B, size int, dir string, r string) {
	root := common.BytesToHash(common.FromHex(r))
	//fmt.Println(root)
	dbase, err := leveldb.New(dir,8,500,"cc",false)
	if err != nil {
		fmt.Println("database create wrong!")
	}
	defer dbase.Close()

	triedb := trie.NewDatabase(dbase)
	tree, _ := trie.New(trie.TrieID(root), triedb)


	FilePath := "D:\\Code\\go-ethereum-logic\\tests_Log1c\\Test_dataset\\dataset\\address.txt"
	File, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer File.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(File)
	keys := make([]string, size)
	for i := 0; i < size; i++ {
		str, _ := reader.ReadString('\n') //读到一个换行就结束
		keys[i] = str
	}

	fmt.Println(keys)



	count := size / 10000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			if j % count == 0 {
				ans, _ := tree.TryGet([]byte(keys[j]))
				fmt.Println(keys[j])
				fmt.Println(ans)
			}

		}
	}
	b.StopTimer()
}

