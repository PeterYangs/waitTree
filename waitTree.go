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

func Background() *emptyWaitTree {

	return &emptyWaitTree{}
}

type WaitTree struct {
	son       []WaitTreeInterface
	waitGroup *sync.WaitGroup
	lock      sync.Mutex
}

func NewWaitTree(parent WaitTreeInterface) *WaitTree {

	wait := sync.WaitGroup{}

	ww := &WaitTree{son: []WaitTreeInterface{}, waitGroup: &wait, lock: sync.Mutex{}}

	switch parent.(type) {

	case *emptyWaitTree:

	case *WaitTree:

		fmt.Println("nice啊")

		t := parent.(*WaitTree)

		t.setSon(ww)

	}

	return ww
}

func (w *WaitTree) Add(delta int) {
	//TODO implement me
	//panic("implement me")

	w.waitGroup.Add(delta)

	//w.waitGroup.Wait()
	//w.waitGroup.Done()

}

func (w *WaitTree) Wait() {
	//TODO implement me
	//panic("implement me")
	w.waitGroup.Wait()

	for _, treeInterface := range w.son {

		treeInterface.Wait()
	}

}

func (w *WaitTree) Done() {
	//TODO implement me
	//panic("implement me")

	w.waitGroup.Done()

}

func (w *WaitTree) setSon(tree *WaitTree) {

	tree.lock.Lock()

	defer tree.lock.Unlock()

	//tree.son = append(tree.son, w)

	w.son = append(w.son, tree)

}

func (w *WaitTree) GetSon() []WaitTreeInterface {

	return w.son
}
