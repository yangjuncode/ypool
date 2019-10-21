package ypool

type Ypool struct {
	//max go routine count for the pool
	maxGoroutine int
}

//adjust pool size min size is 4, return old pool size
func (this *Ypool) AdjustPoll(maxGoroutineCount int) int {
	old := this.maxGoroutine
	if maxGoroutineCount < 4 {
		this.maxGoroutine = 4
	} else {
		this.maxGoroutine = maxGoroutineCount
	}

	return old
}
