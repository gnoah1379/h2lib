package iog

import (
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

type HandlePanic func(rcv interface{})

type IOG struct {
	tasksPending    *int32
	tasksChan       chan task
	workersPool     sync.Pool
	waitGroup       sync.WaitGroup
	firstTimeCallGo sync.Once
	cap             int
}

func New(cap int) *IOG {
	if cap == 0 {
		cap = math.MaxInt32
	}
	wp := &IOG{
		tasksPending: new(int32),
		cap:          cap,
	}

	wp.workersPool.New = func() interface{} {
		if wp.Idle() {
			return nil
		}
		return newWorker(wp)
	}
	return wp
}

// submit task
// just use panicHandler[0]
func (wp *IOG) Go(fn func(), panicHandler ...func(rcv interface{})) {
	wp.firstTimeCallGo.Do(wp.start)
	t := task{job: fn}
	if panicHandler != nil {
		t.panicHandler = panicHandler[0]
	}
	wp.tasksChan <- t
	wp.incrementPending()
	wp.waitGroup.Add(1)
}

func (wp *IOG) ForceClose() {
	close(wp.tasksChan)
}

func (wp *IOG) Wait() {
	wp.waitGroup.Wait()
}

func (wp *IOG) Close() {
	wp.ForceClose()
	wp.Wait()
}

// return tasksChan number waiting for execute
func (wp *IOG) TaskPending() int32 {
	return atomic.LoadInt32(wp.tasksPending)
}

// return idle state
// equal true when tasksPending task = 0
func (wp *IOG) Idle() bool {
	return wp.TaskPending() <= 0
}

func (wp *IOG) start() {
	var t task
	var w *worker
	wp.tasksChan = make(chan task, wp.cap)
	go func() {
		for t = range wp.tasksChan {
			w, _ = wp.workersPool.Get().(*worker)
			if w == nil {
				runtime.Gosched()
			} else {
				wp.decrementPending()
				w.taskChan <- t
			}
		}

	}()

}

func (wp *IOG) taskDone() {
	wp.waitGroup.Done()
}

func (wp *IOG) incrementPending() {
	atomic.AddInt32(wp.tasksPending, 1)
}

func (wp *IOG) decrementPending() {
	atomic.AddInt32(wp.tasksPending, -1)
}
