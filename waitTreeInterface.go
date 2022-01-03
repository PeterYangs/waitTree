package waitTree

type WaitTreeInterface interface {
	Add(delta int)
	Wait()
	Done()
	Release()
}
