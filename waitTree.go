package waitTree

import "sync"

type EmptyWaitTree struct {
}

func (e EmptyWaitTree) Add(delta int) {
	//TODO implement me
	//panic("implement me")
}

func (e EmptyWaitTree) Wait() {
	//TODO implement me
	//panic("implement me")
}

func (e EmptyWaitTree) Done() {
	//TODO implement me
	//panic("implement me")
}

func Background() *EmptyWaitTree {

	return &EmptyWaitTree{}
}

type WaitTree struct {
	son       WaitTreeInterface
	waitGroup *sync.WaitGroup
}

func NewWaitTree(son WaitTreeInterface) *WaitTree {

	wait := sync.WaitGroup{}

	return &WaitTree{son: son, waitGroup: &wait}
}

func (w WaitTree) Add(delta int) {
	//TODO implement me
	//panic("implement me")

	w.waitGroup.Add(delta)

	//w.waitGroup.Wait()
	//w.waitGroup.Done()

}

func (w WaitTree) Wait() {
	//TODO implement me
	//panic("implement me")
	w.waitGroup.Wait()

	w.son.Wait()

}

func (w WaitTree) Done() {
	//TODO implement me
	//panic("implement me")

	w.waitGroup.Done()

}
