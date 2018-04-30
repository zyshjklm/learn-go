package common

import (
	"log"
)

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan Worker
	JobChannel chan Job
	quit       chan bool
}

// NewWorker new a worker
func NewWorker() Worker {
	return Worker{
		WorkerPool: nil, // WorkerPool is control by dispatcher
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) run() {
	for {
		// 1) when a Worker run, add itself to the WorkPool of Dispatcher
		// 2) worker wating for Job by select-case of w.JobChannel
		// 3) read from channel and process it
		// next loop: register the current worker into the WorkerPool again.
		w.WorkerPool <- w

		select {
		case job := <-w.JobChannel:
			// we have received a work request.
			if err := job.DataInfo.ProcessData(); err != nil {
				log.Printf("Process error: %s", err.Error())
			}
		case <-w.quit:
			// we have received a signal to stop
			return
		}
	}
}

// Start method starts the run loop for the worker
func (w Worker) Start() {
	go w.run()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
