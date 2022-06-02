package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	FirstItem *ListItem
	LastItem  *ListItem
	Length    int
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.FirstItem
}

func (l *list) Back() *ListItem {
	return l.LastItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: nil, Next: l.FirstItem}
	if newItem.Next != nil {
		l.FirstItem.Prev = newItem
	}
	l.FirstItem = newItem
	if l.LastItem == nil {
		l.LastItem = newItem
	}
	l.Length++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: l.Back(), Next: nil}
	if newItem.Prev != nil {
		l.LastItem.Next = newItem
	} else {
		l.FirstItem = newItem
	}
	l.LastItem = newItem
	l.Length++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i != nil {
		if i.Prev != nil {
			i.Prev.Next = i.Next
		}
		if i.Next != nil {
			i.Next.Prev = i.Prev
		}
		if i == l.FirstItem {
			l.FirstItem = l.FirstItem.Next
		}
		if i == l.LastItem {
			l.LastItem = l.LastItem.Prev
		}
		l.Length--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i != nil && i != l.FirstItem {
		l.PushFront(i.Value)
		l.Remove(i)
	}
}

func NewList() List {
	return new(list)
}
