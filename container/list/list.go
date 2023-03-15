package list

type List[T any] struct {
	head  *node[T]
	count int
}

type node[T any] struct {
	next *node[T]
	val  T
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func newNode[T any](v T, ptr *node[T]) *node[T] {
	return &node[T]{val: v, next: ptr}
}

func (l *List[T]) HeadAdd(items ...T) {
	defer func() {
		l.count++
	}()

	for _, v := range items {
		l.head = newNode(v, l.head)
	}
}

func (l *List[T]) Append(items ...T) {
	if len(items) == 0 {
		return
	}

	defer func() {
		l.count++
	}()

	if l.head == nil {
		l.head = newNode(items[0], nil)
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	for _, v := range items[1:] {
		cur.next = newNode(v, nil)
		cur = cur.next
	}
}

// DeleteIndex 删除指定下标节点 从0开始
func (l *List[T]) DeleteIndex(idx int) {
	defer func() {
		l.count--
	}()

	if idx == 0 {
		l.head = l.head.next
		return
	}

	i := 0
	cur := l.head
	for cur.next != nil {
		i++
		if i == idx {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}

}

func (l *List[T]) ForEach(fn func(v T)) {
	cur := l.head
	for cur != nil {
		fn(cur.val)
		cur = cur.next
	}
}

func (l *List[T]) List() []T {
	var list []T
	cur := l.head
	for cur != nil {
		list = append(list, cur.val)
		cur = cur.next
	}

	return list
}

func (l *List[T]) Len() int {
	var count int
	cur := l.head
	for cur != nil {
		count++
		cur = cur.next
	}
	return count
}
