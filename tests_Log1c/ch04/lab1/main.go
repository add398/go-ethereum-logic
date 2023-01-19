/**
 * @Author: Log1c
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/1/19 14:33
 */

package main

func main() {
	size := 10000
	cacheSize := 10000
	twoq := NewTwoQueue(size)
	
	keys := get_address(size)

	for i := 0; i < cacheSize; i++ {
		key := keys[i]
		twoq.Set(string(key))
	}

	for i := 0; i < size; i++ {
		
	}
	
}
