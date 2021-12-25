package waitTree

import "sync"

type EmptyWaitTree struct {
}

func (e EmptyWaitTree) Add(delta int) {
	//TODO implement me
	//panic("implement me")
}

func Background() *EmptyWaitTree {

	return &EmptyWaitTree{}
}

type WaitTree struct {
	parent    WaitTreeInterface
	waitGroup *sync.WaitGroup
}

func (w WaitTree) Add(delta int) {
	//TODO implement me
	//panic("implement me")
}

func NewWaitTree(parent WaitTreeInterface) *WaitTree {

	wait := sync.WaitGroup{}

	return &WaitTree{parent: parent, waitGroup: &wait}
}
