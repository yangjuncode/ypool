package ypool

import "go.uber.org/atomic"

//the pool manager
type Ypool struct {
	//max go routine count for the pool
	maxPoolSize atomic.Uint32

	//current goroutine size
	currentPoolSize atomic.Uint32

	//current running  goroutine size
	runningGoroutineSize atomic.Uint32
}

//task priority type with level: None=0,Normal=1,High=2,Emergent=3
type TaskPriority int

const (
	//the lowest priority level,use for logging and no time requirement task
	None TaskPriority = iota
	//the normal priority for common use
	Normal
	//for high level use
	High
	//the highest priority
	Emergent
)

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

//submit a task in pool, equal to SubmitPriority(Normal,task)
func (this *Ypool) Submit(task func()) error {
	//todo impl submit
	return nil
}

//submit a task with priority
func (this *Ypool) SubmitPriority(priority TaskPriority, task func()) error {
	//todo impl SubmitQueue
	return nil
}

//submit a task to pool in queue with name:queueName, equal to SubmitQueuePriority(Normal,queueName,task)
func (this *Ypool) SubmitQueue(queueName interface{}, task func()) error {
	//todo impl SubmitQueue
	return nil
}

//submit a task to pool in queue with priority and queueName
func (this *Ypool) SubmitQueuePriority(priority TaskPriority, queueName interface{}, task func()) error {
	//todo impl SubmitQueue
	return nil
}
