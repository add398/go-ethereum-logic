/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/1/19 14:33
 */

package main

import "fmt"

func main() {
	size := 10000000
	cacheSize := 10000
	twoq := NewTwoQueue(size)
	
	keys := get_address(size)

	for i := 0; i < cacheSize; i++ {
		key := keys[i]
		twoq.Set(string(key))
	}

	miss := 0
	sum := 0
	for i := cacheSize; i < size; i++ {
		if twoq.Get(string(keys[i])) == false {
			miss++
			twoq.Set(string(keys[i]))
		}
		sum++
	}

	fmt.Println("miss: ", miss,   "       sum: ", sum)
	a := 1 -  float64(miss) / float64(sum)
	fmt.Println("命中率为 ： " , a )
	
}
