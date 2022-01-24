package main

// 循环链表
// golang 中有  container/ring

// 循环链表
type Ring struct {
	prv, next *Ring       // 前驱，后驱节点
	Value     interface{} // 存储数据
}

// 初始化: 前驱和后驱都指向自己
func (r *Ring) init() *Ring {
	r.prv = r
	r.next = r
	return r
}

func main() {
	r := new(Ring)
	r.init()
}

// 创建一个指定大小 N 的循环链表
// 连续绑定前驱和后驱节点，时间复杂度 O(n)
func New(n int) *Ring {
	if n < 0 {
		return nil
	}

	r := new(Ring)
	p := r
	for i := 0; i < n; i++ {
		p.next = &Ring{prv: p}
		p = p.next
	}
	p.next = r
	r.prv = p
	return r
}

// 获取前驱，后驱节点  O(1)
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring) Prv() *Ring {
	if r.prv == nil {
		return r.init()
	}
	return r.prv
}

// 选择节点
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prv
		}
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prv()
		r.next = s
		s.prv = r
		n.prv = p
		p.next = n
	}
	return n
}

// 删除节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// 获取长度 Len
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.Next() {
			n++
		}
	}
	return n
}
