package ypool

import "go.uber.org/atomic"

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

//the pool manager
type Ypool struct {
	//max go routine count for the pool
	maxPoolSize atomic.Uint32

	//current goroutine size
	currentPoolSize atomic.Uint32

	//current running  goroutine size
	runningGoroutineSize atomic.Uint32

	//none level worker count
	noneWorkers atomic.Uint32

	//none task count
	noneTasks atomic.Uint32

	//normal worker count
	normalWorkers atomic.Uint32

	//normal task count
	normalTasks atomic.Uint32

	//high worker count
	highWorkers atomic.Uint32

	//high task count
	highTasks atomic.Uint32

	//emergent worker count
	emergentWorkers atomic.Uint32

	//emergent task count
	emergentTasks atomic.Uint32
}

type yTaskItem struct {
	priority  TaskPriority
	queueName interface{}
	task      func()
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
