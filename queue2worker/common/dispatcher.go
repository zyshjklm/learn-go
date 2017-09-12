package common

import (
	"log"
)

// Job represents the job to be run
type Job struct {
	DataInfo DataInfo
}

// Dispatcher dispatch job from JQ to worker in pool
type Dispatcher struct {
	JobQueue   chan Job    // Job Queue
	WorkerPool chan Worker // Work pool that are registered with the dispatcher
	MaxWorker  int
	MaxQueue   int
}

// NewDispatcher new a dispatcher
func NewDispatcher(maxQue int, maxWok int) *Dispatcher {
	jq := make(chan Job, maxWok)
	wk := make(chan Worker, maxQue)

	return &Dispatcher{
		JobQueue:   jq,
		WorkerPool: wk,
		MaxWorker:  maxWok,
		MaxQueue:   maxQue,
	}
}

// Run start all worker and dispatch job to it
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorker; i++ {
		worker := NewWorker()
		worker.WorkerPool = d.WorkerPool
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		log.Printf("[dispatch] queue:%d; pool:%d\n", len(d.JobQueue), len(d.WorkerPool))
		select {
		case job := <-d.JobQueue:
			// a job request has been received
			go func() {
				// try to obtain a worker. this will block until a worker is idle
				worker := <-d.WorkerPool

				// dispatch the job to the worker's job channel
				worker.JobChannel <- job
			}()
		}
	}
}
