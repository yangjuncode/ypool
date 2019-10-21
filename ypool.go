package ypool

import "go.uber.org/atomic"

type Ypool struct {
	//max go routine count for the pool
	maxPoolSize atomic.Uint32

	//current goroutine size
	currentPoolSize atomic.Uint32

	//current running  goroutine size
	runningGoroutineSize atomic.Uint32
}

//adjust pool max size min size is 4, return old pool  max size
func (this *Ypool) SetPoolMaxSize(newPoolMaxSize uint32) uint32 {
	old := this.maxPoolSize.Load()
	if newPoolMaxSize < 4 {
		this.maxPoolSize.Store(4)
	} else {
		this.maxPoolSize.Store(newPoolMaxSize)
	}

	return old
}

// get pool max size
func (this *Ypool) GetPoolMaxSize() uint32 {
	return this.maxPoolSize.Load()
}

func (this *Ypool) Submit(task func()) error {
	//todo impl submit
	return nil
}
func (this *Ypool) SubmitPriority(priority int, task func()) error {
	//todo impl SubmitQueue
	return nil
}

func (this *Ypool) SubmitQueue(queueName interface{}, task func()) error {
	//todo impl SubmitQueue
	return nil
}
func (this *Ypool) SubmitQueuePriority(priority int, queueName interface{}, task func()) error {
	//todo impl SubmitQueue
	return nil
}
