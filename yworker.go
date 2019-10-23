package ypool

type yWorker struct {
	//worker priority
	priority TaskPriority
	//worker status,if <0, the worker need exit
	status int
	//current task
	currentTask *yTaskItem
	//pool manager
	manager *Ypool
}

func newWorker(priority TaskPriority, pool *Ypool) *yWorker {
	w := &yWorker{
		priority: priority,
		manager:  pool,
	}
	go w.Run()

	return w
}

func (this *yWorker) Run() {
	for {
		if this.status < 0 {
			return
		}
		select {}
	}
}
