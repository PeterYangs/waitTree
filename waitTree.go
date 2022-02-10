package waitTree

import (
	"fmt"
	"sync"
)

type emptyWaitTree struct {
}

func (e emptyWaitTree) Add(delta int) {
	//TODO implement me
	//panic("implement me")
}

func (e emptyWaitTree) Wait() {
	//TODO implement me
	//panic("implement me")
}

func (e emptyWaitTree) Done() {
	//TODO implement me
	//panic("implement me")
}

func (e emptyWaitTree) Release() {

}

func Background() *emptyWaitTree {

	return &emptyWaitTree{}
}

type WaitTree struct {
	son       sync.Map
	parent    *WaitTree
	waitGroup *sync.WaitGroup
	lock      sync.Mutex
}

func NewWaitTree(parent WaitTreeInterface) *WaitTree {

	wait := sync.WaitGroup{}

	ww := &WaitTree{son: sync.Map{}, waitGroup: &wait, lock: sync.Mutex{}}

	switch parent.(type) {

	case *emptyWaitTree:

	case *WaitTree:

		t := parent.(*WaitTree)

		t.setSon(ww)

		ww.parent = t

	}

	return ww
}

func (w *WaitTree) Add(delta int) {
	//TODO implement me

	w.waitGroup.Add(delta)

}

func (w *WaitTree) Wait() {
	//TODO implement me

	w.lock.Lock()

	defer w.lock.Unlock()

	w.waitGroup.Wait()

	w.son.Range(func(key, value interface{}) bool {

		tree := value.(*WaitTree)

		tree.Wait()

		return true

	})

}

func (w *WaitTree) Done() {
	//TODO implement me
	//panic("implement me")

	w.waitGroup.Done()

}

// Release 从父waitTree中释放
func (w *WaitTree) Release() {

	w.lock.Lock()

	defer w.lock.Unlock()

	w.parent.son.Delete(w)

}

// ReleaseSon 释放所有子waitTree
func (w *WaitTree) ReleaseSon() {

	w.son.Range(func(key, value interface{}) bool {

		w.son.Delete(key)

		return true
	})

}

func (w *WaitTree) setSon(tree *WaitTree) {

	w.son.Store(tree, tree)

}

func (w *WaitTree) GetSon() {

	w.son.Range(func(key, value interface{}) bool {

		fmt.Println(value)

		return true
	})

}
