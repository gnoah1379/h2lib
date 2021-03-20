package iog

type task struct {
	job          func()
	panicHandler func(rcv interface{})
}

type worker struct {
	taskChan chan task
	pool     *IOG
}

func newWorker(pool *IOG) *worker {
	w := &worker{
		taskChan: make(chan task, 1),
		pool:     pool,
	}
	go w.run()
	return w
}

func (w *worker) run() {
	for task := range w.taskChan {
		w.executeTask(task)
		w.pool.taskDone()
		if w.pool.Idle() {
			close(w.taskChan)
		} else {
			w.pool.workersPool.Put(w)
		}
	}
}

func (w *worker) executeTask(t task) {
	defer func() {
		rcv := recover()
		if rcv != nil {
			if t.panicHandler != nil {
				t.panicHandler(rcv)
			} else {
				panic(rcv)
			}
		}
	}()
	if t.job != nil {
		t.job()
	}
}
