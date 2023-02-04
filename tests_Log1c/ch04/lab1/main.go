/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/1/19 14:33
 */

package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

func help1(cacheSize, size int, ratio float64, keys [][]byte)  float64 {
	twoq, _ := lru.New2QParams(cacheSize, ratio, 1- ratio)



	for i := 0; i < cacheSize; i++ {
		key := keys[i]
		twoq.Add(string(key), key)
	}

	miss := 0
	sum := 0
	for i := cacheSize; i < size; i++ {
		key := keys[i]
		_, ok := twoq.Get(string(key))
		if  ok == false {
			miss++
			twoq.Add(string(key), key)
		}
		sum++
	}

	fmt.Println("miss: ", miss,   "       sum: ", sum)
	a := 1 -  float64(miss) / float64(sum)
	fmt.Println("当 ratio 为: ",  ratio ,  "  命中率为 ： " , a )
	return a
}

//  选择合适的 ratio
func Choose_cacheSize()  {
	size := 10000000
	cacheSize := []int{10000, 100000}
	keys := get_address(size)
	ans := []float64{}
	for i := 0; i < len(cacheSize); i++ {
		a := help1(cacheSize[i],size, 0.25 , keys)
		ans = append(ans, a)
	}

	fmt.Println(ans)
}


//  lab1
func ChooseHigh() {
	//   通过改变 ratio ，选择最高命中率
	size := 20000000
	cacheSize := 50000

	keys := get_address(size)

	nums := []float64{0.05, 0.1, 0.15, 0.2, 0.25, 0.3, 0.35, 0.4, 0.45, 0.5,
		0.55, 0.6, 0.65, 0.7, 0.75, 0.8, 0.85, 0.9, 0.95}

	ans := []float64{}
	for i := 0; i < len(nums); i++ {
		a := help1(cacheSize, size, nums[i], keys)
		ans = append(ans, a)
	}

	fmt.Println(ans)
}




func main() {
	ChooseHigh()
	
}
