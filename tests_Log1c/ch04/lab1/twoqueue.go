/**
 * @Author: Log1c
 * @Description:
 * @File:  twoqueue
 * @Version: 1.0.0
 * @Date: 2023/1/19 14:35
 */

package main


type Cache interface {
	Name() string
	Set(string)
	Get(string) bool
	Close()
}


type TwoQueue struct {
	v *TwoQueueCache
}

func NewTwoQueue(size int, ratio float64) Cache {
	cache, err := New2Q(size, ratio)
	if err != nil {
		panic(err)
	}
	return &TwoQueue{
		v: cache,
	}
}

func (c *TwoQueue) Name() string {
	return "two-queue"
}

func (c *TwoQueue) Set(key string) {
	c.v.Add(key, key)
}

func (c *TwoQueue) Get(key string) bool {
	_, ok := c.v.Get(key)
	return ok
}

func (c *TwoQueue) Close() {}

