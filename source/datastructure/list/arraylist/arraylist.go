package arraylist

import (
	. "../../arrayutils"
	"fmt"
	"bytes"
)

// TODO: growby(n int) resize(cap int) shrink()

type ArrayList struct {
	maxSize   int
	numInList int
	listArray []interface{}
}

func NewArrayList(size int) *ArrayList {
	return &ArrayList{size, 0, make([]interface{}, size)}
}

func (a *ArrayList) Clear() {
	a.numInList = 0
}

func (a *ArrayList) Insert(index int, items ...interface{}) {
	// for each function of Array List, we must judge if list is full,
	// I wonder if there is a way to bind this action to every function.
	length := len(items)
	if a.numInList + length >= a.maxSize {
		panic(&FullListError{})
	}
	if !a.isInList(index) {
		panic(&BadCurrError{})
	}
	for i := a.numInList; i > index + length; i-- {
		a.listArray[i] = a.listArray[i-length]
	}
	for i, item := range items {
		a.listArray[index + i] = item
		a.numInList++
	}
}

func (a *ArrayList) Append(items ...interface{}) {
	length := len(items)
	if a.numInList + length> a.maxSize {
		panic(&FullListError{})
	}
	for _, item := range items {
		a.listArray[a.numInList] = item
		a.numInList++
	}
}

func (a *ArrayList) Remove(index int) interface{} {
	if a.IsEmpty() {
		panic(&EmptyListError{})
	}
	value := a.listArray[index]
	for i := index; i < a.numInList-1; i++ {
		a.listArray[i] = a.listArray[i+1]
	}
	a.numInList--
	return value
}

func (a *ArrayList) Length() int {
	return a.numInList
}

func (a *ArrayList) SetValue(index int, item interface{}) {
	if !a.isInList(index) {
		panic(&BadCurrError{})
	}
	a.listArray[index] = item
}

func (a *ArrayList) Get(index int) interface{} {
	if index < 0 || index > a.numInList {
		panic(&BadCurrError{})
	}
	return a.listArray[index]
}

func (a *ArrayList) Contains(item interface{}) bool {
	for _, value := range a.listArray {
		if value == item {
			return true
		}
	}
	return false
}

func (a *ArrayList) IsEmpty() bool {
	return a.numInList == 0
}

func (a *ArrayList) isInList(index int) bool {
	return index > 0 && index < a.numInList
}

func (a *ArrayList) String() string {
	var buffer bytes.Buffer
	if a.IsEmpty() {
		return "()"
	} else {
		buffer.WriteString("(")
		for i := 0; a.isInList(i); i++ {
			value := a.Get(i)
			buffer.WriteString(fmt.Sprintf("%v", value))
		}
		buffer.WriteString(")")
	}
	return buffer.String()
}
