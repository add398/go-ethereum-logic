/**
 * @Author: Log1c
 * @Description:
 * @File:  leaf_head_tail
 * @Version: 1.0.0
 * @Date: 2022/12/1 16:52
 */

package ART_LRU

import "fmt"

// 只关心 tree 里面 head 和 tail 指针

// add last 添加
func (this *tree) AddLast(x *leaf) {
	x.pre = this.tail.pre
	x.next = this.tail
	this.tail.pre.next = x
	this.tail.pre = x
}

// 删除掉
func (this *tree) Remove(x *leaf)  {
	if this.head.next == this.tail {
		return
	}
	x.pre.next = x.next
	x.next.pre = x.pre
	x.pre = nil
	x.next = nil
}

// 更新， 先删除，再添加
func (t *tree) Update(x *leaf)  {
	t.Remove(x)
	t.AddLast(x)

}

// 超出限制，从头部删除
func (this *tree) RemoveFirst() *leaf {
	if this.head.next == this.tail {
		return nil
	}
	first := this.head.next
	this.Remove(first)
	return first
}

func (t *tree) ForEachByLeaf() {
	head := t.head
	tail := t.tail
	current := head.next
	for current != tail {
		key := current.key
		fmt.Println(string(key))
		current = current.next
	}
	fmt.Println("_____________________")
}

