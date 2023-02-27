/**
 * @Author: Log1c
 * @Description:
 * @File:  util
 * @Version: 1.0.0
 * @Date: 2023/2/27 16:17
 */

package trie

import "bytes"

/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */



var (
	// Trie default value : [byte(0)]
	DefaultLeaf = []byte{0}
)

const (
	HashLength   = 32
	maxPastTries = 300
)

type Hash [HashLength]byte

func bitIsSet(bits []byte, i int) bool {
	return bits[i/8]&(1<<uint(7-i%8)) != 0
}
func bitSet(bits []byte, i int) {
	bits[i/8] |= 1 << uint(7-i%8)
}

// for sorting test data
type DataArray [][]byte

func (d DataArray) Len() int {
	return len(d)
}
func (d DataArray) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d DataArray) Less(i, j int) bool {
	return bytes.Compare(d[i], d[j]) == -1
}
