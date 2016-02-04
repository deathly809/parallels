package parallels

import (
	"sync"

	"github.com/deathly809/gods/queue"
)

// thread data. idk, might need it?
type thread struct {
	myRuntime *runtime
	jobQ      chan Job
}

func (t *thread) Run() {
	for {
		select {
		case d := <-(t.jobQ):
			status := t.myRuntime.statusMap[d.ID()]
			if status == Enqueued {
				status = Running
				t.myRuntime.mutex.Lock()
				t.myRuntime.statusMap[d.ID()] = Running
				t.myRuntime.mutex.Unlock()
			}

			if status == Running {
				if d.Next() {
					t.jobQ <- d
				} else {
					t.myRuntime.mutex.Lock()
					t.myRuntime.statusMap[d.ID()] = Completed
					t.myRuntime.mutex.Unlock()
				}
			}
		default:
			break
		}
	}
}

type runtime struct {
	threads   int
	mutex     sync.Mutex
	jID       int
	q         queue.Queue
	statusMap map[int]JobStatus
	active    bool
	jobChan   chan Job
}

func (r *runtime) StartJob(theJob Job, async bool) int {
	result := -1
	if r.active {
		r.mutex.Lock()
		result = r.jID
		r.jID++
		r.q.Enqueue(theJob)
		r.statusMap[result] = Enqueued
		r.mutex.Unlock()

		if !async {
			for theJob.Next() {
				/* Empty */
			}
		}
	}
	return result
}

func (r *runtime) StopJob(jobID int) bool {
	result := InvalidJob
	if r.active {
		r.mutex.Lock()
		if t, ok := r.statusMap[jobID]; ok {
			if t == Enqueued || t == Running {
				result, r.statusMap[jobID] = Terminated, Terminated
			}
		}
		r.mutex.Unlock()
	}
	return result != InvalidJob
}

// JobStatus returns the status of the job
func (r *runtime) JobStatus(jobID int) JobStatus {
	result := InvalidJob
	if r.active {
		if t, ok := r.statusMap[jobID]; ok {
			result = t
		}
	}
	return result
}

func (r *runtime) Start() {
	if !r.active {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		if !r.active {
			r.q = queue.New()
			r.jobChan = make(chan Job, r.threads)
			r.statusMap = make(map[int]JobStatus)

			for i := 0; i < r.threads; i++ {
				t := &thread{
					myRuntime: r,
					jobQ:      r.jobChan,
				}
				go t.Run()
			}
			r.active = true
		}
	}
}

func (r *runtime) Stop() {
	if r.active {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		r.q = nil
		r.statusMap = nil
		close(r.jobChan)
		r.active = false
	}

}
